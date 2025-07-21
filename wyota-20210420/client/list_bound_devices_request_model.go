// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBoundDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *ListBoundDevicesRequest
	GetAdDomain() *string
	SetAlias(v string) *ListBoundDevicesRequest
	GetAlias() *string
	SetClientType(v int32) *ListBoundDevicesRequest
	GetClientType() *int32
	SetDirectoryId(v string) *ListBoundDevicesRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *ListBoundDevicesRequest
	GetEndUserId() *string
	SetInManage(v bool) *ListBoundDevicesRequest
	GetInManage() *bool
	SetLastLoginUser(v string) *ListBoundDevicesRequest
	GetLastLoginUser() *string
	SetPageNumber(v int32) *ListBoundDevicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBoundDevicesRequest
	GetPageSize() *int32
	SetSerialNo(v string) *ListBoundDevicesRequest
	GetSerialNo() *string
	SetUserType(v string) *ListBoundDevicesRequest
	GetUserType() *string
	SetUuid(v string) *ListBoundDevicesRequest
	GetUuid() *string
}

type ListBoundDevicesRequest struct {
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

func (s ListBoundDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBoundDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListBoundDevicesRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *ListBoundDevicesRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListBoundDevicesRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListBoundDevicesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListBoundDevicesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListBoundDevicesRequest) GetInManage() *bool {
	return s.InManage
}

func (s *ListBoundDevicesRequest) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListBoundDevicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBoundDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBoundDevicesRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListBoundDevicesRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListBoundDevicesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListBoundDevicesRequest) SetAdDomain(v string) *ListBoundDevicesRequest {
	s.AdDomain = &v
	return s
}

func (s *ListBoundDevicesRequest) SetAlias(v string) *ListBoundDevicesRequest {
	s.Alias = &v
	return s
}

func (s *ListBoundDevicesRequest) SetClientType(v int32) *ListBoundDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *ListBoundDevicesRequest) SetDirectoryId(v string) *ListBoundDevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListBoundDevicesRequest) SetEndUserId(v string) *ListBoundDevicesRequest {
	s.EndUserId = &v
	return s
}

func (s *ListBoundDevicesRequest) SetInManage(v bool) *ListBoundDevicesRequest {
	s.InManage = &v
	return s
}

func (s *ListBoundDevicesRequest) SetLastLoginUser(v string) *ListBoundDevicesRequest {
	s.LastLoginUser = &v
	return s
}

func (s *ListBoundDevicesRequest) SetPageNumber(v int32) *ListBoundDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBoundDevicesRequest) SetPageSize(v int32) *ListBoundDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBoundDevicesRequest) SetSerialNo(v string) *ListBoundDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListBoundDevicesRequest) SetUserType(v string) *ListBoundDevicesRequest {
	s.UserType = &v
	return s
}

func (s *ListBoundDevicesRequest) SetUuid(v string) *ListBoundDevicesRequest {
	s.Uuid = &v
	return s
}

func (s *ListBoundDevicesRequest) Validate() error {
	return dara.Validate(s)
}
