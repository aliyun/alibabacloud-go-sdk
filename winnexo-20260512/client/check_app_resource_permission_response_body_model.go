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
	// Indicates whether the target user has the permission on the resource.
	//
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This field is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The target permission.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The dot-separated resource key. The first segment is the APP application key.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type. Set the value to APP_RESOURCE.
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
