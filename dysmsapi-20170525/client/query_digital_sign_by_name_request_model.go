// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDigitalSignByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryDigitalSignByNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryDigitalSignByNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryDigitalSignByNameRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QueryDigitalSignByNameRequest
	GetSignName() *string
}

type QueryDigitalSignByNameRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QueryDigitalSignByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDigitalSignByNameRequest) GoString() string {
	return s.String()
}

func (s *QueryDigitalSignByNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryDigitalSignByNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryDigitalSignByNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryDigitalSignByNameRequest) GetSignName() *string {
	return s.SignName
}

func (s *QueryDigitalSignByNameRequest) SetOwnerId(v int64) *QueryDigitalSignByNameRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDigitalSignByNameRequest) SetResourceOwnerAccount(v string) *QueryDigitalSignByNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDigitalSignByNameRequest) SetResourceOwnerId(v int64) *QueryDigitalSignByNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryDigitalSignByNameRequest) SetSignName(v string) *QueryDigitalSignByNameRequest {
	s.SignName = &v
	return s
}

func (s *QueryDigitalSignByNameRequest) Validate() error {
	return dara.Validate(s)
}
