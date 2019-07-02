package booking

import (
	"fmt"
	"time"
)

type MockErrorService struct{}

func (*MockErrorService) Book(booking Booking) error {
	return fmt.Errorf("mock error")
}

func (*MockErrorService) UnBook(booking Booking) error {
	return fmt.Errorf("mock error")
}

func (*MockErrorService) MyBookings() ([]Booking, error) {
	return nil, fmt.Errorf("mock error")
}

func (*MockErrorService) Available(start time.Time, end time.Time) ([]string, error) {
	return nil, fmt.Errorf("mock error")
}

type MockStaticService struct {
	bookings []Booking
	rooms    []string
}

func NewMockStaticService(bookings []Booking, rooms []string) *MockStaticService {
	return &MockStaticService{bookings: bookings, rooms: rooms}
}

func (*MockStaticService) Book(booking Booking) error {
	return nil
}

func (*MockStaticService) UnBook(booking Booking) error {
	return nil
}

func (ms *MockStaticService) MyBookings() ([]Booking, error) {
	return ms.bookings, nil
}

func (ms *MockStaticService) Available(start time.Time, end time.Time) ([]string, error) {
	return ms.rooms, nil
}

type MockService struct {
	Bookings map[string]*Booking
	Rooms    []string
}

func NewMockService(rooms []string) *MockService {
	return &MockService{Bookings: make(map[string]*Booking), Rooms: rooms}
}

func (bs *MockService) Book(b Booking) error {
	if bs.Bookings[b.Room] != nil {
		return fmt.Errorf("room already booked")
	}

	bs.Bookings[b.Room] = &b
	return nil
}

func (bs *MockService) UnBook(b Booking) error {
	if bs.Bookings[b.Room] == nil {
		return fmt.Errorf("room not booked")
	}

	delete(bs.Bookings, b.Room)
	return nil
}

func (bs *MockService) MyBookings() ([]Booking, error) {
	var bookings []Booking
	for _, v := range bs.Bookings {
		bookings = append(bookings, *v)
	}
	return bookings, nil
}

func (bs *MockService) Available(start time.Time, end time.Time) ([]string, error) {
	var available []string
	for _, room := range bs.Rooms {
		if bs.Bookings[room] == nil {
			available = append(available, room)
		}
	}
	return available, nil
}
