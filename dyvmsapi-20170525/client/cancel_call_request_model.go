// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *CancelCallRequest
	GetCallId() *string
	SetOwnerId(v int64) *CancelCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelCallRequest
	GetResourceOwnerId() *int64
}

type CancelCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 116012854210^10281427****
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelCallRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCallRequest) GoString() string {
	return s.String()
}

func (s *CancelCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *CancelCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelCallRequest) SetCallId(v string) *CancelCallRequest {
	s.CallId = &v
	return s
}

func (s *CancelCallRequest) SetOwnerId(v int64) *CancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCallRequest) SetResourceOwnerAccount(v string) *CancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelCallRequest) SetResourceOwnerId(v int64) *CancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelCallRequest) Validate() error {
	return dara.Validate(s)
}
