// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAlbumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListSubAlbumRequestDeviceInfo) *ListSubAlbumRequest
	GetDeviceInfo() *ListSubAlbumRequestDeviceInfo
	SetQuerySubscriptionAlbumRequest(v *ListSubAlbumRequestQuerySubscriptionAlbumRequest) *ListSubAlbumRequest
	GetQuerySubscriptionAlbumRequest() *ListSubAlbumRequestQuerySubscriptionAlbumRequest
	SetUserInfo(v *ListSubAlbumRequestUserInfo) *ListSubAlbumRequest
	GetUserInfo() *ListSubAlbumRequestUserInfo
}

type ListSubAlbumRequest struct {
	DeviceInfo *ListSubAlbumRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// request
	QuerySubscriptionAlbumRequest *ListSubAlbumRequestQuerySubscriptionAlbumRequest `json:"QuerySubscriptionAlbumRequest,omitempty" xml:"QuerySubscriptionAlbumRequest,omitempty" type:"Struct"`
	UserInfo                      *ListSubAlbumRequestUserInfo                      `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListSubAlbumRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequest) GetDeviceInfo() *ListSubAlbumRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListSubAlbumRequest) GetQuerySubscriptionAlbumRequest() *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	return s.QuerySubscriptionAlbumRequest
}

func (s *ListSubAlbumRequest) GetUserInfo() *ListSubAlbumRequestUserInfo {
	return s.UserInfo
}

func (s *ListSubAlbumRequest) SetDeviceInfo(v *ListSubAlbumRequestDeviceInfo) *ListSubAlbumRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListSubAlbumRequest) SetQuerySubscriptionAlbumRequest(v *ListSubAlbumRequestQuerySubscriptionAlbumRequest) *ListSubAlbumRequest {
	s.QuerySubscriptionAlbumRequest = v
	return s
}

func (s *ListSubAlbumRequest) SetUserInfo(v *ListSubAlbumRequestUserInfo) *ListSubAlbumRequest {
	s.UserInfo = v
	return s
}

func (s *ListSubAlbumRequest) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubAlbumRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListSubAlbumRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListSubAlbumRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListSubAlbumRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListSubAlbumRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSubAlbumRequestDeviceInfo) SetEncodeKey(v string) *ListSubAlbumRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetEncodeType(v string) *ListSubAlbumRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetId(v string) *ListSubAlbumRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetIdType(v string) *ListSubAlbumRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) SetOrganizationId(v string) *ListSubAlbumRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListSubAlbumRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumRequestQuerySubscriptionAlbumRequest struct {
	// example:
	//
	// 4476001
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80011
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	Page *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 睡前故事
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequest) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) GetAlbumId() *string {
	return s.AlbumId
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) GetPage() *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage {
	return s.Page
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) GetTitle() *string {
	return s.Title
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetAlbumId(v string) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.AlbumId = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetCategoryId(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.CategoryId = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetPage(v *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.Page = v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) SetTitle(v string) *ListSubAlbumRequestQuerySubscriptionAlbumRequest {
	s.Title = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequest) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumRequestQuerySubscriptionAlbumRequestPage struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) SetPageNum(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage {
	s.PageNum = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) SetPageSize(v int32) *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage {
	s.PageSize = &v
	return s
}

func (s *ListSubAlbumRequestQuerySubscriptionAlbumRequestPage) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubAlbumRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListSubAlbumRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListSubAlbumRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListSubAlbumRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListSubAlbumRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSubAlbumRequestUserInfo) SetEncodeKey(v string) *ListSubAlbumRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetEncodeType(v string) *ListSubAlbumRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetId(v string) *ListSubAlbumRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetIdType(v string) *ListSubAlbumRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) SetOrganizationId(v string) *ListSubAlbumRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListSubAlbumRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
