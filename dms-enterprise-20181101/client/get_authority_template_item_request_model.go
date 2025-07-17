// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *GetAuthorityTemplateItemRequest
	GetTemplateId() *int64
	SetTid(v int64) *GetAuthorityTemplateItemRequest
	GetTid() *int64
}

type GetAuthorityTemplateItemRequest struct {
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

func (s GetAuthorityTemplateItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateItemRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateItemRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetAuthorityTemplateItemRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetAuthorityTemplateItemRequest) SetTemplateId(v int64) *GetAuthorityTemplateItemRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAuthorityTemplateItemRequest) SetTid(v int64) *GetAuthorityTemplateItemRequest {
	s.Tid = &v
	return s
}

func (s *GetAuthorityTemplateItemRequest) Validate() error {
	return dara.Validate(s)
}
