// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMusicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListMusicRequestDeviceInfo) *ListMusicRequest
	GetDeviceInfo() *ListMusicRequestDeviceInfo
	SetPayload(v *ListMusicRequestPayload) *ListMusicRequest
	GetPayload() *ListMusicRequestPayload
	SetUserInfo(v *ListMusicRequestUserInfo) *ListMusicRequest
	GetUserInfo() *ListMusicRequestUserInfo
}

type ListMusicRequest struct {
	// This parameter is required.
	DeviceInfo *ListMusicRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *ListMusicRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListMusicRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListMusicRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMusicRequest) GoString() string {
	return s.String()
}

func (s *ListMusicRequest) GetDeviceInfo() *ListMusicRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListMusicRequest) GetPayload() *ListMusicRequestPayload {
	return s.Payload
}

func (s *ListMusicRequest) GetUserInfo() *ListMusicRequestUserInfo {
	return s.UserInfo
}

func (s *ListMusicRequest) SetDeviceInfo(v *ListMusicRequestDeviceInfo) *ListMusicRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListMusicRequest) SetPayload(v *ListMusicRequestPayload) *ListMusicRequest {
	s.Payload = v
	return s
}

func (s *ListMusicRequest) SetUserInfo(v *ListMusicRequestUserInfo) *ListMusicRequest {
	s.UserInfo = v
	return s
}

func (s *ListMusicRequest) Validate() error {
	return dara.Validate(s)
}

type ListMusicRequestDeviceInfo struct {
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

func (s ListMusicRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMusicRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListMusicRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListMusicRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListMusicRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListMusicRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListMusicRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMusicRequestDeviceInfo) SetEncodeKey(v string) *ListMusicRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetEncodeType(v string) *ListMusicRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetId(v string) *ListMusicRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetIdType(v string) *ListMusicRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) SetOrganizationId(v string) *ListMusicRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListMusicRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListMusicRequestPayload struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1
	MusicId   *int64  `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// This parameter is required.
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMusicRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListMusicRequestPayload) GoString() string {
	return s.String()
}

func (s *ListMusicRequestPayload) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMusicRequestPayload) GetMusicId() *int64 {
	return s.MusicId
}

func (s *ListMusicRequestPayload) GetMusicName() *string {
	return s.MusicName
}

func (s *ListMusicRequestPayload) GetMusicType() *int64 {
	return s.MusicType
}

func (s *ListMusicRequestPayload) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *ListMusicRequestPayload) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMusicRequestPayload) SetCurrentPage(v int32) *ListMusicRequestPayload {
	s.CurrentPage = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicId(v int64) *ListMusicRequestPayload {
	s.MusicId = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicName(v string) *ListMusicRequestPayload {
	s.MusicName = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicType(v int64) *ListMusicRequestPayload {
	s.MusicType = &v
	return s
}

func (s *ListMusicRequestPayload) SetMusicTypeName(v string) *ListMusicRequestPayload {
	s.MusicTypeName = &v
	return s
}

func (s *ListMusicRequestPayload) SetPageSize(v int32) *ListMusicRequestPayload {
	s.PageSize = &v
	return s
}

func (s *ListMusicRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListMusicRequestUserInfo struct {
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

func (s ListMusicRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMusicRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListMusicRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListMusicRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListMusicRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListMusicRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListMusicRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMusicRequestUserInfo) SetEncodeKey(v string) *ListMusicRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetEncodeType(v string) *ListMusicRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetId(v string) *ListMusicRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetIdType(v string) *ListMusicRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListMusicRequestUserInfo) SetOrganizationId(v string) *ListMusicRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListMusicRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
