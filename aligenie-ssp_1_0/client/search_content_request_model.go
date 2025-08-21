// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *SearchContentRequestDeviceInfo) *SearchContentRequest
	GetDeviceInfo() *SearchContentRequestDeviceInfo
	SetRequest(v *SearchContentRequestRequest) *SearchContentRequest
	GetRequest() *SearchContentRequestRequest
	SetUserInfo(v *SearchContentRequestUserInfo) *SearchContentRequest
	GetUserInfo() *SearchContentRequestUserInfo
}

type SearchContentRequest struct {
	// This parameter is required.
	DeviceInfo *SearchContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Request *SearchContentRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *SearchContentRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SearchContentRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchContentRequest) GoString() string {
	return s.String()
}

func (s *SearchContentRequest) GetDeviceInfo() *SearchContentRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *SearchContentRequest) GetRequest() *SearchContentRequestRequest {
	return s.Request
}

func (s *SearchContentRequest) GetUserInfo() *SearchContentRequestUserInfo {
	return s.UserInfo
}

func (s *SearchContentRequest) SetDeviceInfo(v *SearchContentRequestDeviceInfo) *SearchContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *SearchContentRequest) SetRequest(v *SearchContentRequestRequest) *SearchContentRequest {
	s.Request = v
	return s
}

func (s *SearchContentRequest) SetUserInfo(v *SearchContentRequestUserInfo) *SearchContentRequest {
	s.UserInfo = v
	return s
}

func (s *SearchContentRequest) Validate() error {
	return dara.Validate(s)
}

type SearchContentRequestDeviceInfo struct {
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

func (s SearchContentRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SearchContentRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SearchContentRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SearchContentRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *SearchContentRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *SearchContentRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SearchContentRequestDeviceInfo) SetEncodeKey(v string) *SearchContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetEncodeType(v string) *SearchContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetId(v string) *SearchContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetIdType(v string) *SearchContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) SetOrganizationId(v string) *SearchContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *SearchContentRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type SearchContentRequestRequest struct {
	// example:
	//
	// music
	Cate *string `json:"Cate,omitempty" xml:"Cate,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// false
	QueryAlbum *bool `json:"QueryAlbum,omitempty" xml:"QueryAlbum,omitempty"`
	// example:
	//
	// singer
	SubCate *string `json:"SubCate,omitempty" xml:"SubCate,omitempty"`
}

func (s SearchContentRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchContentRequestRequest) GoString() string {
	return s.String()
}

func (s *SearchContentRequestRequest) GetCate() *string {
	return s.Cate
}

func (s *SearchContentRequestRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *SearchContentRequestRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchContentRequestRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchContentRequestRequest) GetQueryAlbum() *bool {
	return s.QueryAlbum
}

func (s *SearchContentRequestRequest) GetSubCate() *string {
	return s.SubCate
}

func (s *SearchContentRequestRequest) SetCate(v string) *SearchContentRequestRequest {
	s.Cate = &v
	return s
}

func (s *SearchContentRequestRequest) SetPageNum(v int32) *SearchContentRequestRequest {
	s.PageNum = &v
	return s
}

func (s *SearchContentRequestRequest) SetPageSize(v int32) *SearchContentRequestRequest {
	s.PageSize = &v
	return s
}

func (s *SearchContentRequestRequest) SetQuery(v string) *SearchContentRequestRequest {
	s.Query = &v
	return s
}

func (s *SearchContentRequestRequest) SetQueryAlbum(v bool) *SearchContentRequestRequest {
	s.QueryAlbum = &v
	return s
}

func (s *SearchContentRequestRequest) SetSubCate(v string) *SearchContentRequestRequest {
	s.SubCate = &v
	return s
}

func (s *SearchContentRequestRequest) Validate() error {
	return dara.Validate(s)
}

type SearchContentRequestUserInfo struct {
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

func (s SearchContentRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SearchContentRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SearchContentRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SearchContentRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *SearchContentRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *SearchContentRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SearchContentRequestUserInfo) SetEncodeKey(v string) *SearchContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetEncodeType(v string) *SearchContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetId(v string) *SearchContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetIdType(v string) *SearchContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SearchContentRequestUserInfo) SetOrganizationId(v string) *SearchContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *SearchContentRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
