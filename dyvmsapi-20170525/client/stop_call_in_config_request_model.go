// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCallInConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *StopCallInConfigRequest
	GetNumber() *string
	SetOwnerId(v int64) *StopCallInConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopCallInConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopCallInConfigRequest
	GetResourceOwnerId() *int64
}

type StopCallInConfigRequest struct {
	// The China 400 number from which the inbound call to be stopped is transferred.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400***
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopCallInConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCallInConfigRequest) GoString() string {
	return s.String()
}

func (s *StopCallInConfigRequest) GetNumber() *string {
	return s.Number
}

func (s *StopCallInConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopCallInConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopCallInConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopCallInConfigRequest) SetNumber(v string) *StopCallInConfigRequest {
	s.Number = &v
	return s
}

func (s *StopCallInConfigRequest) SetOwnerId(v int64) *StopCallInConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCallInConfigRequest) SetResourceOwnerAccount(v string) *StopCallInConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopCallInConfigRequest) SetResourceOwnerId(v int64) *StopCallInConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopCallInConfigRequest) Validate() error {
	return dara.Validate(s)
}
