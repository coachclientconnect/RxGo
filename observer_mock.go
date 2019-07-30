package rxgo

import (
	"github.com/stretchr/testify/mock"
)

// NewObserverMock creates a mock observer
func NewObserverMock() *ObserverMock {
	obMock := new(ObserverMock)
	obMock.On("OnDone").Return()
	obMock.On("OnNext", mock.Anything).Return()
	obMock.On("OnError", mock.Anything).Return()
	return obMock
}

// ObserverMock is a mock observer
// TODO Visibility
type ObserverMock struct {
	mock.Mock
}

// OnDone provides a mock function with given fields:
func (m *ObserverMock) OnDone() {
	m.Called()
}

// OnError provides a mock function with given fields: err
func (m *ObserverMock) OnError(err error) {
	m.Called(err)
}

// OnNext provides a mock function with given fields: item
func (m *ObserverMock) OnNext(item interface{}) {
	m.Called(item)
}

// Capture captures observer inputs
func (m *ObserverMock) Capture() Observer {
	ob := NewObserver(
		NextFunc(func(el interface{}) {
			m.OnNext(el)
		}),
		ErrFunc(func(err error) {
			m.OnError(err)
		}),
		DoneFunc(func() {
			m.OnDone()
		}),
	)
	return ob
}
