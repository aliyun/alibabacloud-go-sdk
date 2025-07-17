// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTemplateAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *RevokeTemplateAuthorityRequest
	GetTemplateId() *int64
	SetTid(v int64) *RevokeTemplateAuthorityRequest
	GetTid() *int64
	SetUserIds(v string) *RevokeTemplateAuthorityRequest
	GetUserIds() *string
}

type RevokeTemplateAuthorityRequest struct {
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
	// The IDs of users from whom you want to revoke permissions by using a permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// [12***,34***,56***]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s RevokeTemplateAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTemplateAuthorityRequest) GoString() string {
	return s.String()
}

func (s *RevokeTemplateAuthorityRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *RevokeTemplateAuthorityRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RevokeTemplateAuthorityRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *RevokeTemplateAuthorityRequest) SetTemplateId(v int64) *RevokeTemplateAuthorityRequest {
	s.TemplateId = &v
	return s
}

func (s *RevokeTemplateAuthorityRequest) SetTid(v int64) *RevokeTemplateAuthorityRequest {
	s.Tid = &v
	return s
}

func (s *RevokeTemplateAuthorityRequest) SetUserIds(v string) *RevokeTemplateAuthorityRequest {
	s.UserIds = &v
	return s
}

func (s *RevokeTemplateAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
