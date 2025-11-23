// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *DeleteAuthorityTemplateRequest
	GetTemplateId() *int64
	SetTid(v int64) *DeleteAuthorityTemplateRequest
	GetTid() *int64
}

type DeleteAuthorityTemplateRequest struct {
	// The ID of the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2592
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteAuthorityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorityTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorityTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteAuthorityTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteAuthorityTemplateRequest) SetTemplateId(v int64) *DeleteAuthorityTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteAuthorityTemplateRequest) SetTid(v int64) *DeleteAuthorityTemplateRequest {
	s.Tid = &v
	return s
}

func (s *DeleteAuthorityTemplateRequest) Validate() error {
	return dara.Validate(s)
}
