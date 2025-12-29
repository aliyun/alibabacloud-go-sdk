// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryAvailableAuthCodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryAvailableAuthCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAvailableAuthCodeRequest
	GetResourceOwnerId() *int64
	SetTagId(v int64) *QueryAvailableAuthCodeRequest
	GetTagId() *int64
}

type QueryAvailableAuthCodeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 22
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s QueryAvailableAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAvailableAuthCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAvailableAuthCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAvailableAuthCodeRequest) GetTagId() *int64 {
	return s.TagId
}

func (s *QueryAvailableAuthCodeRequest) SetOwnerId(v int64) *QueryAvailableAuthCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetResourceOwnerAccount(v string) *QueryAvailableAuthCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetResourceOwnerId(v int64) *QueryAvailableAuthCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) SetTagId(v int64) *QueryAvailableAuthCodeRequest {
	s.TagId = &v
	return s
}

func (s *QueryAvailableAuthCodeRequest) Validate() error {
	return dara.Validate(s)
}
