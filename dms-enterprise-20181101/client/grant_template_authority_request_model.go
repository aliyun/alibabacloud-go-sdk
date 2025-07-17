// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTemplateAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *GrantTemplateAuthorityRequest
	GetComment() *string
	SetExpireDate(v string) *GrantTemplateAuthorityRequest
	GetExpireDate() *string
	SetTemplateId(v int64) *GrantTemplateAuthorityRequest
	GetTemplateId() *int64
	SetTid(v int64) *GrantTemplateAuthorityRequest
	GetTid() *int64
	SetUserIds(v string) *GrantTemplateAuthorityRequest
	GetUserIds() *string
}

type GrantTemplateAuthorityRequest struct {
	// The reason why you want to grant permissions on resources to the users by using the permission template.
	//
	// example:
	//
	// Business test.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the permission expires. Specify the time in the yyyy-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-08 17:58:53
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
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
	// The IDs of users to which you want to grant permissions on resources by using the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// [12***,34***,56***]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s GrantTemplateAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantTemplateAuthorityRequest) GoString() string {
	return s.String()
}

func (s *GrantTemplateAuthorityRequest) GetComment() *string {
	return s.Comment
}

func (s *GrantTemplateAuthorityRequest) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GrantTemplateAuthorityRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GrantTemplateAuthorityRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GrantTemplateAuthorityRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *GrantTemplateAuthorityRequest) SetComment(v string) *GrantTemplateAuthorityRequest {
	s.Comment = &v
	return s
}

func (s *GrantTemplateAuthorityRequest) SetExpireDate(v string) *GrantTemplateAuthorityRequest {
	s.ExpireDate = &v
	return s
}

func (s *GrantTemplateAuthorityRequest) SetTemplateId(v int64) *GrantTemplateAuthorityRequest {
	s.TemplateId = &v
	return s
}

func (s *GrantTemplateAuthorityRequest) SetTid(v int64) *GrantTemplateAuthorityRequest {
	s.Tid = &v
	return s
}

func (s *GrantTemplateAuthorityRequest) SetUserIds(v string) *GrantTemplateAuthorityRequest {
	s.UserIds = &v
	return s
}

func (s *GrantTemplateAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
