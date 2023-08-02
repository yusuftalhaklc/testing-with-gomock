package models

import (
	"github.com/golang/mock/gomock"
	"github.com/yusuftalhaklc/testing-with-gomock/mocks"
	"testing"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := mocks.NewMockHandler(mockCtrl)
	testUser := &User{Handler: mockController}

	mockController.EXPECT().CreateUser("User created").Return(nil).Times(1)

	testUser.Create()
}
