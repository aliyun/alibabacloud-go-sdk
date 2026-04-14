// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTagRequest
	GetResourceOwnerId() *int64
	SetTagDescription(v string) *CreateTagRequest
	GetTagDescription() *string
	SetTagName(v string) *CreateTagRequest
	GetTagName() *string
}

type CreateTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Tag description
	//
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// Tag name. Limitations: 1-50 characters, allowing English letters, numbers, and underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTagRequest) GetTagDescription() *string {
	return s.TagDescription
}

func (s *CreateTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *CreateTagRequest) SetOwnerId(v int64) *CreateTagRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTagRequest) SetResourceOwnerAccount(v string) *CreateTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTagRequest) SetResourceOwnerId(v int64) *CreateTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTagRequest) SetTagDescription(v string) *CreateTagRequest {
	s.TagDescription = &v
	return s
}

func (s *CreateTagRequest) SetTagName(v string) *CreateTagRequest {
	s.TagName = &v
	return s
}

func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
