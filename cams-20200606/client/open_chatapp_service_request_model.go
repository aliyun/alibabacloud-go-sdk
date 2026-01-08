// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenChatappServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenChatappServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *OpenChatappServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenChatappServiceRequest
	GetResourceOwnerId() *int64
}

type OpenChatappServiceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenChatappServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenChatappServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenChatappServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenChatappServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenChatappServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenChatappServiceRequest) SetOwnerId(v int64) *OpenChatappServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenChatappServiceRequest) SetResourceOwnerAccount(v string) *OpenChatappServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenChatappServiceRequest) SetResourceOwnerId(v int64) *OpenChatappServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenChatappServiceRequest) Validate() error {
	return dara.Validate(s)
}
