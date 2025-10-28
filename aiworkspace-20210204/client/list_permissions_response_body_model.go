// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody
	GetPermissions() []*ListPermissionsResponseBodyPermissions
	SetRequestId(v string) *ListPermissionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPermissionsResponseBody
	GetTotalCount() *int64
}

type ListPermissionsResponseBody struct {
	// The permissions.
	Permissions []*ListPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2AE63638-5420-56DC-B******8174039A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of permissions that meet the filter conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) GetPermissions() []*ListPermissionsResponseBodyPermissions {
	return s.Permissions
}

func (s *ListPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) SetTotalCount(v int64) *ListPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPermissionsResponseBody) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionsResponseBodyPermissions struct {
	// The permission name, which is unique in a region. For more information about permissions, see [Appendix: Roles and permissions](https://help.aliyun.com/document_detail/2840449.html). The example value PaiDLC:GetTensorboard indicates the permission to view details about a TensorBoard job on the Deep Learning Containers (DLC) page.
	//
	// example:
	//
	// PaiDLC:GetTensorboard
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	// The permission rules.
	PermissionRules []*ListPermissionsResponseBodyPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) GetPermissionCode() *string {
	return s.PermissionCode
}

func (s *ListPermissionsResponseBodyPermissions) GetPermissionRules() []*ListPermissionsResponseBodyPermissionsPermissionRules {
	return s.PermissionRules
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionCode(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionCode = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionRules(v []*ListPermissionsResponseBodyPermissionsPermissionRules) *ListPermissionsResponseBodyPermissions {
	s.PermissionRules = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) Validate() error {
	if s.PermissionRules != nil {
		for _, item := range s.PermissionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionsResponseBodyPermissionsPermissionRules struct {
	// The accessibility of the permission rule. Valid values:
	//
	// 	- PUBLIC: All members in the workspace can access the permission rule.
	//
	// 	- PRIVATE: Only the creator can access the permission rule.
	//
	// 	- ANY: All users can access the permission rule.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The type of access. If you set Accessibility to PUBLIC, all users can access the workspace. This parameter is invalid. If you set Accessibility to PRIVATE, the permissions are determined based on the value of EntityAccessType. The value of EntityAccessType can be:
	//
	// 	- CREATOR: Only the creator can access the workspace.
	//
	// 	- ANY: All users can access the workspace.
	//
	// example:
	//
	// CREATOR
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsPermissionRules) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) GetEntityAccessType() *string {
	return s.EntityAccessType
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) SetAccessibility(v string) *ListPermissionsResponseBodyPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) SetEntityAccessType(v string) *ListPermissionsResponseBodyPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) Validate() error {
	return dara.Validate(s)
}
