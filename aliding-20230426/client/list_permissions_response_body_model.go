// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int64) *ListPermissionsResponseBody
	GetDuration() *int64
	SetNextToken(v string) *ListPermissionsResponseBody
	GetNextToken() *string
	SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody
	GetPermissions() []*ListPermissionsResponseBodyPermissions
	SetRequestId(v string) *ListPermissionsResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ListPermissionsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListPermissionsResponseBody
	GetVendorType() *string
}

type ListPermissionsResponseBody struct {
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1
	NextToken   *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Permissions []*ListPermissionsResponseBodyPermissions `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *ListPermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionsResponseBody) GetPermissions() []*ListPermissionsResponseBodyPermissions {
	return s.Permissions
}

func (s *ListPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListPermissionsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListPermissionsResponseBody) SetDuration(v int64) *ListPermissionsResponseBody {
	s.Duration = &v
	return s
}

func (s *ListPermissionsResponseBody) SetNextToken(v string) *ListPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) SetVendorRequestId(v string) *ListPermissionsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) SetVendorType(v string) *ListPermissionsResponseBody {
	s.VendorType = &v
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
	// example:
	//
	// 123456
	DentryUuid *string                                       `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	Member     *ListPermissionsResponseBodyPermissionsMember `json:"Member,omitempty" xml:"Member,omitempty" type:"Struct"`
	Role       *ListPermissionsResponseBodyPermissionsRole   `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *ListPermissionsResponseBodyPermissions) GetMember() *ListPermissionsResponseBodyPermissionsMember {
	return s.Member
}

func (s *ListPermissionsResponseBodyPermissions) GetRole() *ListPermissionsResponseBodyPermissionsRole {
	return s.Role
}

func (s *ListPermissionsResponseBodyPermissions) SetDentryUuid(v string) *ListPermissionsResponseBodyPermissions {
	s.DentryUuid = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetMember(v *ListPermissionsResponseBodyPermissionsMember) *ListPermissionsResponseBodyPermissions {
	s.Member = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetRole(v *ListPermissionsResponseBodyPermissionsRole) *ListPermissionsResponseBodyPermissions {
	s.Role = v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionsResponseBodyPermissionsMember struct {
	// example:
	//
	// 123456
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ORG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsMember) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsMember) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsMember) GetCorpId() *string {
	return s.CorpId
}

func (s *ListPermissionsResponseBodyPermissionsMember) GetId() *string {
	return s.Id
}

func (s *ListPermissionsResponseBodyPermissionsMember) GetType() *string {
	return s.Type
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetCorpId(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.CorpId = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetId(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.Id = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsMember) SetType(v string) *ListPermissionsResponseBodyPermissionsMember {
	s.Type = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsMember) Validate() error {
	return dara.Validate(s)
}

type ListPermissionsResponseBodyPermissionsRole struct {
	// example:
	//
	// OWNER
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 拥有者
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsRole) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsRole) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsRole) GetId() *string {
	return s.Id
}

func (s *ListPermissionsResponseBodyPermissionsRole) GetName() *string {
	return s.Name
}

func (s *ListPermissionsResponseBodyPermissionsRole) SetId(v string) *ListPermissionsResponseBodyPermissionsRole {
	s.Id = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsRole) SetName(v string) *ListPermissionsResponseBodyPermissionsRole {
	s.Name = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsRole) Validate() error {
	return dara.Validate(s)
}
