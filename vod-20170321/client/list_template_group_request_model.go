// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsContainsTemplates(v bool) *ListTemplateGroupRequest
	GetIsContainsTemplates() *bool
	SetOwnerId(v int64) *ListTemplateGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTemplateGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTemplateGroupRequest
	GetResourceOwnerId() *int64
}

type ListTemplateGroupRequest struct {
	IsContainsTemplates  *bool   `json:"IsContainsTemplates,omitempty" xml:"IsContainsTemplates,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupRequest) GetIsContainsTemplates() *bool {
	return s.IsContainsTemplates
}

func (s *ListTemplateGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTemplateGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTemplateGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTemplateGroupRequest) SetIsContainsTemplates(v bool) *ListTemplateGroupRequest {
	s.IsContainsTemplates = &v
	return s
}

func (s *ListTemplateGroupRequest) SetOwnerId(v int64) *ListTemplateGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTemplateGroupRequest) SetResourceOwnerAccount(v string) *ListTemplateGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTemplateGroupRequest) SetResourceOwnerId(v int64) *ListTemplateGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
