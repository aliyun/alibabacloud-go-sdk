// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifyTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTagRequest
	GetResourceOwnerId() *int64
	SetTagDescription(v string) *ModifyTagRequest
	GetTagDescription() *string
	SetTagId(v int32) *ModifyTagRequest
	GetTagId() *int32
	SetTagName(v string) *ModifyTagRequest
	GetTagName() *string
}

type ModifyTagRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Tag Description
	//
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// Tag ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 100674
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// Tag Name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ModifyTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTagRequest) GoString() string {
	return s.String()
}

func (s *ModifyTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTagRequest) GetTagDescription() *string {
	return s.TagDescription
}

func (s *ModifyTagRequest) GetTagId() *int32 {
	return s.TagId
}

func (s *ModifyTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *ModifyTagRequest) SetOwnerId(v int64) *ModifyTagRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTagRequest) SetResourceOwnerAccount(v string) *ModifyTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTagRequest) SetResourceOwnerId(v int64) *ModifyTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTagRequest) SetTagDescription(v string) *ModifyTagRequest {
	s.TagDescription = &v
	return s
}

func (s *ModifyTagRequest) SetTagId(v int32) *ModifyTagRequest {
	s.TagId = &v
	return s
}

func (s *ModifyTagRequest) SetTagName(v string) *ModifyTagRequest {
	s.TagName = &v
	return s
}

func (s *ModifyTagRequest) Validate() error {
	return dara.Validate(s)
}
