// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *GetAuthorityTemplateRequest
	GetTemplateId() *int64
	SetTid(v int64) *GetAuthorityTemplateRequest
	GetTid() *int64
}

type GetAuthorityTemplateRequest struct {
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

func (s GetAuthorityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetAuthorityTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetAuthorityTemplateRequest) SetTemplateId(v int64) *GetAuthorityTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAuthorityTemplateRequest) SetTid(v int64) *GetAuthorityTemplateRequest {
	s.Tid = &v
	return s
}

func (s *GetAuthorityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
