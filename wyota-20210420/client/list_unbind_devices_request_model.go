// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnbindDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *ListUnbindDevicesRequest
	GetAdDomain() *string
	SetAlias(v string) *ListUnbindDevicesRequest
	GetAlias() *string
	SetClientType(v int32) *ListUnbindDevicesRequest
	GetClientType() *int32
	SetDirectoryId(v string) *ListUnbindDevicesRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *ListUnbindDevicesRequest
	GetEndUserId() *string
	SetInManage(v bool) *ListUnbindDevicesRequest
	GetInManage() *bool
	SetLastLoginUser(v string) *ListUnbindDevicesRequest
	GetLastLoginUser() *string
	SetPageNumber(v int32) *ListUnbindDevicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUnbindDevicesRequest
	GetPageSize() *int32
	SetSerialNo(v string) *ListUnbindDevicesRequest
	GetSerialNo() *string
	SetUserType(v string) *ListUnbindDevicesRequest
	GetUserType() *string
	SetUuid(v string) *ListUnbindDevicesRequest
	GetUuid() *string
}

type ListUnbindDevicesRequest struct {
	AdDomain      *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	Alias         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	InManage      *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	LastLoginUser *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo      *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnbindDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnbindDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListUnbindDevicesRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *ListUnbindDevicesRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListUnbindDevicesRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListUnbindDevicesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUnbindDevicesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListUnbindDevicesRequest) GetInManage() *bool {
	return s.InManage
}

func (s *ListUnbindDevicesRequest) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListUnbindDevicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUnbindDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnbindDevicesRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListUnbindDevicesRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListUnbindDevicesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnbindDevicesRequest) SetAdDomain(v string) *ListUnbindDevicesRequest {
	s.AdDomain = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetAlias(v string) *ListUnbindDevicesRequest {
	s.Alias = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetClientType(v int32) *ListUnbindDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetDirectoryId(v string) *ListUnbindDevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetEndUserId(v string) *ListUnbindDevicesRequest {
	s.EndUserId = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetInManage(v bool) *ListUnbindDevicesRequest {
	s.InManage = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetLastLoginUser(v string) *ListUnbindDevicesRequest {
	s.LastLoginUser = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetPageNumber(v int32) *ListUnbindDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetPageSize(v int32) *ListUnbindDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetSerialNo(v string) *ListUnbindDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetUserType(v string) *ListUnbindDevicesRequest {
	s.UserType = &v
	return s
}

func (s *ListUnbindDevicesRequest) SetUuid(v string) *ListUnbindDevicesRequest {
	s.Uuid = &v
	return s
}

func (s *ListUnbindDevicesRequest) Validate() error {
	return dara.Validate(s)
}
