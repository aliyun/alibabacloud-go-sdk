// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverCallInConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *RecoverCallInConfigRequest
	GetNumber() *string
	SetOwnerId(v int64) *RecoverCallInConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RecoverCallInConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RecoverCallInConfigRequest
	GetResourceOwnerId() *int64
}

type RecoverCallInConfigRequest struct {
	// The China 400 number that is used to transfer the inbound call.
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

func (s RecoverCallInConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverCallInConfigRequest) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigRequest) GetNumber() *string {
	return s.Number
}

func (s *RecoverCallInConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RecoverCallInConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RecoverCallInConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RecoverCallInConfigRequest) SetNumber(v string) *RecoverCallInConfigRequest {
	s.Number = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetOwnerId(v int64) *RecoverCallInConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetResourceOwnerAccount(v string) *RecoverCallInConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetResourceOwnerId(v int64) *RecoverCallInConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RecoverCallInConfigRequest) Validate() error {
	return dara.Validate(s)
}
