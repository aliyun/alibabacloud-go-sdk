// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallSoundRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetCallSoundRecordRequest
	GetCallId() *string
	SetCreateTime(v string) *GetCallSoundRecordRequest
	GetCreateTime() *string
	SetOwnerId(v int64) *GetCallSoundRecordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCallSoundRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCallSoundRecordRequest
	GetResourceOwnerId() *int64
}

type GetCallSoundRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 125165515022^11195613****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-10-15 08:56:23
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallSoundRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallSoundRecordRequest) GoString() string {
	return s.String()
}

func (s *GetCallSoundRecordRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetCallSoundRecordRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCallSoundRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCallSoundRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCallSoundRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCallSoundRecordRequest) SetCallId(v string) *GetCallSoundRecordRequest {
	s.CallId = &v
	return s
}

func (s *GetCallSoundRecordRequest) SetCreateTime(v string) *GetCallSoundRecordRequest {
	s.CreateTime = &v
	return s
}

func (s *GetCallSoundRecordRequest) SetOwnerId(v int64) *GetCallSoundRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallSoundRecordRequest) SetResourceOwnerAccount(v string) *GetCallSoundRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallSoundRecordRequest) SetResourceOwnerId(v int64) *GetCallSoundRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCallSoundRecordRequest) Validate() error {
	return dara.Validate(s)
}
