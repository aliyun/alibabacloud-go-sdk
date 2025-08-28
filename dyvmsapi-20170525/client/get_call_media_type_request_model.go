// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallMediaTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *GetCallMediaTypeRequest
	GetCallId() *string
	SetCalledNumber(v string) *GetCallMediaTypeRequest
	GetCalledNumber() *string
	SetOwnerId(v int64) *GetCallMediaTypeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCallMediaTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCallMediaTypeRequest
	GetResourceOwnerId() *int64
}

type GetCallMediaTypeRequest struct {
	// example:
	//
	// 116012854210^10281427****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1580000****
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallMediaTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallMediaTypeRequest) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetCallMediaTypeRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *GetCallMediaTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCallMediaTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCallMediaTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCallMediaTypeRequest) SetCallId(v string) *GetCallMediaTypeRequest {
	s.CallId = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetCalledNumber(v string) *GetCallMediaTypeRequest {
	s.CalledNumber = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetOwnerId(v int64) *GetCallMediaTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetResourceOwnerAccount(v string) *GetCallMediaTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetResourceOwnerId(v int64) *GetCallMediaTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCallMediaTypeRequest) Validate() error {
	return dara.Validate(s)
}
