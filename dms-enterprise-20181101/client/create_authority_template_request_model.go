// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAuthorityTemplateRequest
	GetDescription() *string
	SetName(v string) *CreateAuthorityTemplateRequest
	GetName() *string
	SetTid(v int64) *CreateAuthorityTemplateRequest
	GetTid() *int64
}

type CreateAuthorityTemplateRequest struct {
	// The description of the permission template.
	//
	// example:
	//
	// This template is used for business testing.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateAuthorityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorityTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorityTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAuthorityTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateAuthorityTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateAuthorityTemplateRequest) SetDescription(v string) *CreateAuthorityTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateAuthorityTemplateRequest) SetName(v string) *CreateAuthorityTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAuthorityTemplateRequest) SetTid(v int64) *CreateAuthorityTemplateRequest {
	s.Tid = &v
	return s
}

func (s *CreateAuthorityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
