// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPermission(v *GetPermissionResponseBodyPermission) *GetPermissionResponseBody
	GetPermission() *GetPermissionResponseBodyPermission
	SetRequestId(v string) *GetPermissionResponseBody
	GetRequestId() *string
}

type GetPermissionResponseBody struct {
	// The information about the permission.
	Permission *GetPermissionResponseBodyPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2F23CFB6-A721-4E90-AC1E-0E30FA8B45DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) GetPermission() *GetPermissionResponseBodyPermission {
	return s.Permission
}

func (s *GetPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPermissionResponseBody) SetPermission(v *GetPermissionResponseBodyPermission) *GetPermissionResponseBody {
	s.Permission = v
	return s
}

func (s *GetPermissionResponseBody) SetRequestId(v string) *GetPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPermissionResponseBodyPermission struct {
	// The creation time.
	//
	// example:
	//
	// 2020-12-07T07:39:01.818Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the permission is the default permission. Valid values:
	//
	// 	- false: The permission is not the default permission.
	//
	// 	- true: The permission is the default permission.
	//
	// example:
	//
	// true
	DefaultPermission *bool `json:"DefaultPermission,omitempty" xml:"DefaultPermission,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// 	- false: The version is not the default version.
	//
	// 	- true: The version is the default version.
	//
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The document of the policy related to the permission.
	//
	// example:
	//
	// {"Effect":"Allow","Action":["vpc:DescribeVSwitches","vpc:DescribeVSwitchAttributes"]}
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// AliyunRSDefaultPermissionVSwitch
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	// The version of the permission.
	//
	// example:
	//
	// v1
	PermissionVersion *string `json:"PermissionVersion,omitempty" xml:"PermissionVersion,omitempty"`
	// The type of the shared resources.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-12-07T07:39:01.818Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPermissionResponseBodyPermission) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionResponseBodyPermission) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyPermission) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPermissionResponseBodyPermission) GetDefaultPermission() *bool {
	return s.DefaultPermission
}

func (s *GetPermissionResponseBodyPermission) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *GetPermissionResponseBodyPermission) GetPermission() *string {
	return s.Permission
}

func (s *GetPermissionResponseBodyPermission) GetPermissionName() *string {
	return s.PermissionName
}

func (s *GetPermissionResponseBodyPermission) GetPermissionVersion() *string {
	return s.PermissionVersion
}

func (s *GetPermissionResponseBodyPermission) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetPermissionResponseBodyPermission) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPermissionResponseBodyPermission) SetCreateTime(v string) *GetPermissionResponseBodyPermission {
	s.CreateTime = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetDefaultPermission(v bool) *GetPermissionResponseBodyPermission {
	s.DefaultPermission = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetDefaultVersion(v bool) *GetPermissionResponseBodyPermission {
	s.DefaultVersion = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermission(v string) *GetPermissionResponseBodyPermission {
	s.Permission = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermissionName(v string) *GetPermissionResponseBodyPermission {
	s.PermissionName = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetPermissionVersion(v string) *GetPermissionResponseBodyPermission {
	s.PermissionVersion = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetResourceType(v string) *GetPermissionResponseBodyPermission {
	s.ResourceType = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) SetUpdateTime(v string) *GetPermissionResponseBodyPermission {
	s.UpdateTime = &v
	return s
}

func (s *GetPermissionResponseBodyPermission) Validate() error {
	return dara.Validate(s)
}
