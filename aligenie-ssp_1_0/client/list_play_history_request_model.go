// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlayHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListPlayHistoryRequestDeviceInfo) *ListPlayHistoryRequest
	GetDeviceInfo() *ListPlayHistoryRequestDeviceInfo
	SetRequest(v *ListPlayHistoryRequestRequest) *ListPlayHistoryRequest
	GetRequest() *ListPlayHistoryRequestRequest
	SetUserInfo(v *ListPlayHistoryRequestUserInfo) *ListPlayHistoryRequest
	GetUserInfo() *ListPlayHistoryRequestUserInfo
}

type ListPlayHistoryRequest struct {
	// This parameter is required.
	DeviceInfo *ListPlayHistoryRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Request *ListPlayHistoryRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListPlayHistoryRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListPlayHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequest) GetDeviceInfo() *ListPlayHistoryRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListPlayHistoryRequest) GetRequest() *ListPlayHistoryRequestRequest {
	return s.Request
}

func (s *ListPlayHistoryRequest) GetUserInfo() *ListPlayHistoryRequestUserInfo {
	return s.UserInfo
}

func (s *ListPlayHistoryRequest) SetDeviceInfo(v *ListPlayHistoryRequestDeviceInfo) *ListPlayHistoryRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListPlayHistoryRequest) SetRequest(v *ListPlayHistoryRequestRequest) *ListPlayHistoryRequest {
	s.Request = v
	return s
}

func (s *ListPlayHistoryRequest) SetUserInfo(v *ListPlayHistoryRequestUserInfo) *ListPlayHistoryRequest {
	s.UserInfo = v
	return s
}

func (s *ListPlayHistoryRequest) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListPlayHistoryRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListPlayHistoryRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListPlayHistoryRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListPlayHistoryRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListPlayHistoryRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListPlayHistoryRequestDeviceInfo) SetEncodeKey(v string) *ListPlayHistoryRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetEncodeType(v string) *ListPlayHistoryRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetId(v string) *ListPlayHistoryRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetIdType(v string) *ListPlayHistoryRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) SetOrganizationId(v string) *ListPlayHistoryRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListPlayHistoryRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryRequestRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// music
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPlayHistoryRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryRequestRequest) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPlayHistoryRequestRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPlayHistoryRequestRequest) GetType() *string {
	return s.Type
}

func (s *ListPlayHistoryRequestRequest) SetPageNum(v int32) *ListPlayHistoryRequestRequest {
	s.PageNum = &v
	return s
}

func (s *ListPlayHistoryRequestRequest) SetPageSize(v int32) *ListPlayHistoryRequestRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlayHistoryRequestRequest) SetType(v string) *ListPlayHistoryRequestRequest {
	s.Type = &v
	return s
}

func (s *ListPlayHistoryRequestRequest) Validate() error {
	return dara.Validate(s)
}

type ListPlayHistoryRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListPlayHistoryRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListPlayHistoryRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListPlayHistoryRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListPlayHistoryRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListPlayHistoryRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListPlayHistoryRequestUserInfo) SetEncodeKey(v string) *ListPlayHistoryRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetEncodeType(v string) *ListPlayHistoryRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetId(v string) *ListPlayHistoryRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetIdType(v string) *ListPlayHistoryRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) SetOrganizationId(v string) *ListPlayHistoryRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListPlayHistoryRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
