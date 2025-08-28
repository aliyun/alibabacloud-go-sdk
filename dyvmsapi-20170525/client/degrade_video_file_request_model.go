// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDegradeVideoFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *DegradeVideoFileRequest
	GetCallId() *string
	SetCalledNumber(v string) *DegradeVideoFileRequest
	GetCalledNumber() *string
	SetMediaType(v string) *DegradeVideoFileRequest
	GetMediaType() *string
	SetOutId(v string) *DegradeVideoFileRequest
	GetOutId() *string
	SetOwnerId(v int64) *DegradeVideoFileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DegradeVideoFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DegradeVideoFileRequest
	GetResourceOwnerId() *int64
}

type DegradeVideoFileRequest struct {
	// example:
	//
	// 116012354148^1028137841****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1590****000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 225869*****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DegradeVideoFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DegradeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileRequest) GetCallId() *string {
	return s.CallId
}

func (s *DegradeVideoFileRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DegradeVideoFileRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *DegradeVideoFileRequest) GetOutId() *string {
	return s.OutId
}

func (s *DegradeVideoFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DegradeVideoFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DegradeVideoFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DegradeVideoFileRequest) SetCallId(v string) *DegradeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetCalledNumber(v string) *DegradeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *DegradeVideoFileRequest) SetMediaType(v string) *DegradeVideoFileRequest {
	s.MediaType = &v
	return s
}

func (s *DegradeVideoFileRequest) SetOutId(v string) *DegradeVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetOwnerId(v int64) *DegradeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetResourceOwnerAccount(v string) *DegradeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DegradeVideoFileRequest) SetResourceOwnerId(v int64) *DegradeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DegradeVideoFileRequest) Validate() error {
	return dara.Validate(s)
}
