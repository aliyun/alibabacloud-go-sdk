// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowed(v bool) *CheckAppResourcePermissionResponseBody
	GetAllowed() *bool
	SetCode(v string) *CheckAppResourcePermissionResponseBody
	GetCode() *string
	SetMessage(v string) *CheckAppResourcePermissionResponseBody
	GetMessage() *string
	SetPermission(v string) *CheckAppResourcePermissionResponseBody
	GetPermission() *string
	SetRequestId(v string) *CheckAppResourcePermissionResponseBody
	GetRequestId() *string
	SetResourceKey(v string) *CheckAppResourcePermissionResponseBody
	GetResourceKey() *string
	SetResourceType(v string) *CheckAppResourcePermissionResponseBody
	GetResourceType() *string
}

type CheckAppResourcePermissionResponseBody struct {
	// 目标用户是否拥有该资源权限
	//
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// 业务状态码；成功为 200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述；成功时为空
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 目标权限点
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 首段为 APP 应用键的点分资源键
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// 固定为 APP_RESOURCE
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CheckAppResourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionResponseBody) GetAllowed() *bool {
	return s.Allowed
}

func (s *CheckAppResourcePermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAppResourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAppResourcePermissionResponseBody) GetPermission() *string {
	return s.Permission
}

func (s *CheckAppResourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAppResourcePermissionResponseBody) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CheckAppResourcePermissionResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckAppResourcePermissionResponseBody) SetAllowed(v bool) *CheckAppResourcePermissionResponseBody {
	s.Allowed = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetCode(v string) *CheckAppResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetMessage(v string) *CheckAppResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetPermission(v string) *CheckAppResourcePermissionResponseBody {
	s.Permission = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetRequestId(v string) *CheckAppResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetResourceKey(v string) *CheckAppResourcePermissionResponseBody {
	s.ResourceKey = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) SetResourceType(v string) *CheckAppResourcePermissionResponseBody {
	s.ResourceType = &v
	return s
}

func (s *CheckAppResourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
