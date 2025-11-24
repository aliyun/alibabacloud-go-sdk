// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v string) *GrantUserPermissionsRequest
	GetPermissions() *string
	SetSubAccountUserId(v string) *GrantUserPermissionsRequest
	GetSubAccountUserId() *string
	SetSubAccountUserIds(v []*string) *GrantUserPermissionsRequest
	GetSubAccountUserIds() []*string
}

type GrantUserPermissionsRequest struct {
	// The permissions that are granted to an entity. The content is a string that consists of JSON arrays. You must specify all permissions that you want to grant to an entity. You can add or remove permissions by modifying the content. Field definition of the sample code:
	//
	// 	- `IsCustom`: a parameter that is required by the system. Set the value to `true`.
	//
	// 	- `Cluster`: the ID of the Service Mesh (ASM) instance.
	//
	// 	- `RoleName`: the name of the permissions. Valid values: `istio-admin`, `istio-ops`, and `istio-readonly`. A value of istio-admin indicates the permissions of ASM administrators. A value of istio-ops indicates the permissions of ASM restricted users. A value of istio-readonly indicates the read-only permissions.
	//
	// 	- `IsRamRole`: the entity to which you want to grant the permissions. To grant the permissions to a RAM role, set the value to `true`. Otherwise, set the value to `false`.
	//
	// example:
	//
	// [{"IsCustom":true,"RoleName":"istio-ops","Cluster":"c57b848115458460583a4260cb713****","RoleType":"custom","IsRamRole":false},{"IsCustom":true,"RoleName":"istio-ops","Cluster":"c4ec191f4e12145c09286ea3e2b8f****","RoleType":"custom","IsRamRole":false}]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the RAM user or RAM role.
	//
	// example:
	//
	// 27852573609480****
	SubAccountUserId *string `json:"SubAccountUserId,omitempty" xml:"SubAccountUserId,omitempty"`
	// The IDs of RAM users or RAM roles. You can grant permissions to multiple users at a time.
	SubAccountUserIds []*string `json:"SubAccountUserIds,omitempty" xml:"SubAccountUserIds,omitempty" type:"Repeated"`
}

func (s GrantUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsRequest) GetPermissions() *string {
	return s.Permissions
}

func (s *GrantUserPermissionsRequest) GetSubAccountUserId() *string {
	return s.SubAccountUserId
}

func (s *GrantUserPermissionsRequest) GetSubAccountUserIds() []*string {
	return s.SubAccountUserIds
}

func (s *GrantUserPermissionsRequest) SetPermissions(v string) *GrantUserPermissionsRequest {
	s.Permissions = &v
	return s
}

func (s *GrantUserPermissionsRequest) SetSubAccountUserId(v string) *GrantUserPermissionsRequest {
	s.SubAccountUserId = &v
	return s
}

func (s *GrantUserPermissionsRequest) SetSubAccountUserIds(v []*string) *GrantUserPermissionsRequest {
	s.SubAccountUserIds = v
	return s
}

func (s *GrantUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
