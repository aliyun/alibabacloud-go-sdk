// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SkipVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *SkipVideoFileRequest
	GetCalledNumber() *string
	SetOutId(v string) *SkipVideoFileRequest
	GetOutId() *string
	SetOwnerId(v int64) *SkipVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SkipVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SkipVideoFileRequest
	GetResourceOwnerId() *int64
	SetSkipTimes(v int64) *SkipVideoFileRequest
	GetSkipTimes() *int64
}

type SkipVideoFileRequest struct {
	// example:
	//
	// 116012854210^10281427****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// PR0210428****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 04:02:01
	SkipTimes *int64 `json:"SkipTimes,omitempty" xml:"SkipTimes,omitempty"`
}

func (s SkipVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SkipVideoFileRequest) GoString() string {
	return s.String()
}

func (s *SkipVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *SkipVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *SkipVideoFileRequest) GetOutId() *string {
	return s.OutId
}

func (s *SkipVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SkipVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SkipVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SkipVideoFileRequest) GetSkipTimes() *int64 {
	return s.SkipTimes
}

func (s *SkipVideoFileRequest) SetCallId(v string) *SkipVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *SkipVideoFileRequest) SetCalledNumber(v string) *SkipVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *SkipVideoFileRequest) SetOutId(v string) *SkipVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *SkipVideoFileRequest) SetOwnerId(v int64) *SkipVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SkipVideoFileRequest) SetResourceOwnerAccount(v string) *SkipVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SkipVideoFileRequest) SetResourceOwnerId(v int64) *SkipVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SkipVideoFileRequest) SetSkipTimes(v int64) *SkipVideoFileRequest {
	s.SkipTimes = &v
	return s
}

func (s *SkipVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
