// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v int64) *ListSmsTemplateRequest
	GetCustomerId() *int64
	SetOwnerId(v int64) *ListSmsTemplateRequest
	GetOwnerId() *int64
	SetProdCode(v string) *ListSmsTemplateRequest
	GetProdCode() *string
	SetQueryTemplate(v string) *ListSmsTemplateRequest
	GetQueryTemplate() *string
	SetResourceOwnerAccount(v string) *ListSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListSmsTemplateRequest
	GetResourceOwnerId() *int64
}

type ListSmsTemplateRequest struct {
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	QueryTemplate        *string `json:"QueryTemplate,omitempty" xml:"QueryTemplate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListSmsTemplateRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *ListSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSmsTemplateRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListSmsTemplateRequest) GetQueryTemplate() *string {
	return s.QueryTemplate
}

func (s *ListSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListSmsTemplateRequest) SetCustomerId(v int64) *ListSmsTemplateRequest {
	s.CustomerId = &v
	return s
}

func (s *ListSmsTemplateRequest) SetOwnerId(v int64) *ListSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSmsTemplateRequest) SetProdCode(v string) *ListSmsTemplateRequest {
	s.ProdCode = &v
	return s
}

func (s *ListSmsTemplateRequest) SetQueryTemplate(v string) *ListSmsTemplateRequest {
	s.QueryTemplate = &v
	return s
}

func (s *ListSmsTemplateRequest) SetResourceOwnerAccount(v string) *ListSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSmsTemplateRequest) SetResourceOwnerId(v int64) *ListSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
