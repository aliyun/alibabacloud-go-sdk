// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsContainsTemplates(v bool) *ListTemplateGroupConsoleRequest
	GetIsContainsTemplates() *bool
	SetOwnerId(v int64) *ListTemplateGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTemplateGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTemplateGroupConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *ListTemplateGroupConsoleRequest
	GetResourceRealOwnerId() *int64
}

type ListTemplateGroupConsoleRequest struct {
	IsContainsTemplates  *bool   `json:"IsContainsTemplates,omitempty" xml:"IsContainsTemplates,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s ListTemplateGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateGroupConsoleRequest) GetIsContainsTemplates() *bool {
	return s.IsContainsTemplates
}

func (s *ListTemplateGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTemplateGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTemplateGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTemplateGroupConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ListTemplateGroupConsoleRequest) SetIsContainsTemplates(v bool) *ListTemplateGroupConsoleRequest {
	s.IsContainsTemplates = &v
	return s
}

func (s *ListTemplateGroupConsoleRequest) SetOwnerId(v int64) *ListTemplateGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTemplateGroupConsoleRequest) SetResourceOwnerAccount(v string) *ListTemplateGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTemplateGroupConsoleRequest) SetResourceOwnerId(v int64) *ListTemplateGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTemplateGroupConsoleRequest) SetResourceRealOwnerId(v int64) *ListTemplateGroupConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ListTemplateGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
