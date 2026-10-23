// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterAppResourcePermissionsResponseBody
	GetCode() *string
	SetMessage(v string) *RegisterAppResourcePermissionsResponseBody
	GetMessage() *string
	SetRecords(v []*RegisterAppResourcePermissionsResponseBodyRecords) *RegisterAppResourcePermissionsResponseBody
	GetRecords() []*RegisterAppResourcePermissionsResponseBodyRecords
	SetRequestId(v string) *RegisterAppResourcePermissionsResponseBody
	GetRequestId() *string
}

type RegisterAppResourcePermissionsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message. This parameter is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The newly created or idempotently matched permission records.
	Records []*RegisterAppResourcePermissionsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RegisterAppResourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterAppResourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterAppResourcePermissionsResponseBody) GetRecords() []*RegisterAppResourcePermissionsResponseBodyRecords {
	return s.Records
}

func (s *RegisterAppResourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterAppResourcePermissionsResponseBody) SetCode(v string) *RegisterAppResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetMessage(v string) *RegisterAppResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetRecords(v []*RegisterAppResourcePermissionsResponseBodyRecords) *RegisterAppResourcePermissionsResponseBody {
	s.Records = v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetRequestId(v string) *RegisterAppResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RegisterAppResourcePermissionsResponseBodyRecords struct {
	// The creation time.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-12-31T23:59:59+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The platform user ID of the grantor.
	//
	// example:
	//
	// 10086
	GrantedBy *string `json:"grantedBy,omitempty" xml:"grantedBy,omitempty"`
	// The principal ID.
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// The principal name.
	//
	// example:
	//
	// Zhang San
	GranteeName *string `json:"granteeName,omitempty" xml:"granteeName,omitempty"`
	// The tenant to which the cross-tenant USER belongs.
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// The permission.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The permission record ID.
	//
	// example:
	//
	// 1001
	PermissionId *int64 `json:"permissionId,omitempty" xml:"permissionId,omitempty"`
	// The resource key.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s RegisterAppResourcePermissionsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGrantedBy() *string {
	return s.GrantedBy
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeId() *string {
	return s.GranteeId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeName() *string {
	return s.GranteeName
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeType() *string {
	return s.GranteeType
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetPermission() *string {
	return s.Permission
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetPermissionId() *int64 {
	return s.PermissionId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetResourceType() *string {
	return s.ResourceType
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetCreatedAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.CreatedAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetExpireAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ExpireAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGrantedBy(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GrantedBy = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeId(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeName(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeName = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeTenantId(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeType(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeType = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetPermission(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.Permission = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetPermissionId(v int64) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.PermissionId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetResourceKey(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ResourceKey = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetResourceType(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ResourceType = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetUpdatedAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
