// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateAuthorityTemplateRequest
	GetDescription() *string
	SetName(v string) *UpdateAuthorityTemplateRequest
	GetName() *string
	SetTemplateId(v int64) *UpdateAuthorityTemplateRequest
	GetTemplateId() *int64
	SetTid(v int64) *UpdateAuthorityTemplateRequest
	GetTid() *int64
}

type UpdateAuthorityTemplateRequest struct {
	// The description of the permission template.
	//
	// >  You must specify the Name or Description parameter. Otherwise, the API call fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// This template is used for business testing.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission template.
	//
	// >  You must specify the Name or Description parameter. Otherwise, the API call fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateAuthorityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorityTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorityTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAuthorityTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAuthorityTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpdateAuthorityTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateAuthorityTemplateRequest) SetDescription(v string) *UpdateAuthorityTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateAuthorityTemplateRequest) SetName(v string) *UpdateAuthorityTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateAuthorityTemplateRequest) SetTemplateId(v int64) *UpdateAuthorityTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateAuthorityTemplateRequest) SetTid(v int64) *UpdateAuthorityTemplateRequest {
	s.Tid = &v
	return s
}

func (s *UpdateAuthorityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
