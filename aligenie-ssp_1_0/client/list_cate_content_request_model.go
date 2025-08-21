// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListCateContentRequestDeviceInfo) *ListCateContentRequest
	GetDeviceInfo() *ListCateContentRequestDeviceInfo
	SetRequest(v *ListCateContentRequestRequest) *ListCateContentRequest
	GetRequest() *ListCateContentRequestRequest
	SetUserInfo(v *ListCateContentRequestUserInfo) *ListCateContentRequest
	GetUserInfo() *ListCateContentRequestUserInfo
}

type ListCateContentRequest struct {
	// This parameter is required.
	DeviceInfo *ListCateContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Request *ListCateContentRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListCateContentRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListCateContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentRequest) GetDeviceInfo() *ListCateContentRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListCateContentRequest) GetRequest() *ListCateContentRequestRequest {
	return s.Request
}

func (s *ListCateContentRequest) GetUserInfo() *ListCateContentRequestUserInfo {
	return s.UserInfo
}

func (s *ListCateContentRequest) SetDeviceInfo(v *ListCateContentRequestDeviceInfo) *ListCateContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListCateContentRequest) SetRequest(v *ListCateContentRequestRequest) *ListCateContentRequest {
	s.Request = v
	return s
}

func (s *ListCateContentRequest) SetUserInfo(v *ListCateContentRequestUserInfo) *ListCateContentRequest {
	s.UserInfo = v
	return s
}

func (s *ListCateContentRequest) Validate() error {
	return dara.Validate(s)
}

type ListCateContentRequestDeviceInfo struct {
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
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListCateContentRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListCateContentRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListCateContentRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListCateContentRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListCateContentRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListCateContentRequestDeviceInfo) SetEncodeKey(v string) *ListCateContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetEncodeType(v string) *ListCateContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetId(v string) *ListCateContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetIdType(v string) *ListCateContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) SetOrganizationId(v string) *ListCateContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListCateContentRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListCateContentRequestRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 80010
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsAlbum *bool `json:"IsAlbum,omitempty" xml:"IsAlbum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// hot_score
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListCateContentRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentRequestRequest) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *ListCateContentRequestRequest) GetIsAlbum() *bool {
	return s.IsAlbum
}

func (s *ListCateContentRequestRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListCateContentRequestRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCateContentRequestRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCateContentRequestRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListCateContentRequestRequest) SetCateId(v int64) *ListCateContentRequestRequest {
	s.CateId = &v
	return s
}

func (s *ListCateContentRequestRequest) SetIsAlbum(v bool) *ListCateContentRequestRequest {
	s.IsAlbum = &v
	return s
}

func (s *ListCateContentRequestRequest) SetPageNum(v int32) *ListCateContentRequestRequest {
	s.PageNum = &v
	return s
}

func (s *ListCateContentRequestRequest) SetPageSize(v int32) *ListCateContentRequestRequest {
	s.PageSize = &v
	return s
}

func (s *ListCateContentRequestRequest) SetSortBy(v string) *ListCateContentRequestRequest {
	s.SortBy = &v
	return s
}

func (s *ListCateContentRequestRequest) SetSortOrder(v string) *ListCateContentRequestRequest {
	s.SortOrder = &v
	return s
}

func (s *ListCateContentRequestRequest) Validate() error {
	return dara.Validate(s)
}

type ListCateContentRequestUserInfo struct {
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
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListCateContentRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListCateContentRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListCateContentRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListCateContentRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListCateContentRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListCateContentRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListCateContentRequestUserInfo) SetEncodeKey(v string) *ListCateContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetEncodeType(v string) *ListCateContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetId(v string) *ListCateContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetIdType(v string) *ListCateContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListCateContentRequestUserInfo) SetOrganizationId(v string) *ListCateContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListCateContentRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
