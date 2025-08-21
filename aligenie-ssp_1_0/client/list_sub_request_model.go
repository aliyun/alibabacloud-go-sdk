// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListSubRequestDeviceInfo) *ListSubRequest
	GetDeviceInfo() *ListSubRequestDeviceInfo
	SetPage(v *ListSubRequestPage) *ListSubRequest
	GetPage() *ListSubRequestPage
	SetUserInfo(v *ListSubRequestUserInfo) *ListSubRequest
	GetUserInfo() *ListSubRequestUserInfo
}

type ListSubRequest struct {
	// This parameter is required.
	DeviceInfo *ListSubRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Page *ListSubRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListSubRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListSubRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubRequest) GoString() string {
	return s.String()
}

func (s *ListSubRequest) GetDeviceInfo() *ListSubRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListSubRequest) GetPage() *ListSubRequestPage {
	return s.Page
}

func (s *ListSubRequest) GetUserInfo() *ListSubRequestUserInfo {
	return s.UserInfo
}

func (s *ListSubRequest) SetDeviceInfo(v *ListSubRequestDeviceInfo) *ListSubRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListSubRequest) SetPage(v *ListSubRequestPage) *ListSubRequest {
	s.Page = v
	return s
}

func (s *ListSubRequest) SetUserInfo(v *ListSubRequestUserInfo) *ListSubRequest {
	s.UserInfo = v
	return s
}

func (s *ListSubRequest) Validate() error {
	return dara.Validate(s)
}

type ListSubRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListSubRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListSubRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListSubRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListSubRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListSubRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSubRequestDeviceInfo) SetEncodeKey(v string) *ListSubRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetEncodeType(v string) *ListSubRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetId(v string) *ListSubRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetIdType(v string) *ListSubRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListSubRequestDeviceInfo) SetOrganizationId(v string) *ListSubRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListSubRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListSubRequestPage struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListSubRequestPage) GoString() string {
	return s.String()
}

func (s *ListSubRequestPage) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListSubRequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubRequestPage) SetPageNum(v int32) *ListSubRequestPage {
	s.PageNum = &v
	return s
}

func (s *ListSubRequestPage) SetPageSize(v int32) *ListSubRequestPage {
	s.PageSize = &v
	return s
}

func (s *ListSubRequestPage) Validate() error {
	return dara.Validate(s)
}

type ListSubRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNION_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListSubRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListSubRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListSubRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListSubRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListSubRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListSubRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSubRequestUserInfo) SetEncodeKey(v string) *ListSubRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListSubRequestUserInfo) SetEncodeType(v string) *ListSubRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListSubRequestUserInfo) SetId(v string) *ListSubRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListSubRequestUserInfo) SetIdType(v string) *ListSubRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListSubRequestUserInfo) SetOrganizationId(v string) *ListSubRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListSubRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
