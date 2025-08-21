// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *CreatePlayingListRequestDeviceInfo) *CreatePlayingListRequest
	GetDeviceInfo() *CreatePlayingListRequestDeviceInfo
	SetOpenCreatePlayingListRequest(v *CreatePlayingListRequestOpenCreatePlayingListRequest) *CreatePlayingListRequest
	GetOpenCreatePlayingListRequest() *CreatePlayingListRequestOpenCreatePlayingListRequest
	SetUserInfo(v *CreatePlayingListRequestUserInfo) *CreatePlayingListRequest
	GetUserInfo() *CreatePlayingListRequestUserInfo
}

type CreatePlayingListRequest struct {
	// This parameter is required.
	DeviceInfo *CreatePlayingListRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenCreatePlayingListRequest *CreatePlayingListRequestOpenCreatePlayingListRequest `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *CreatePlayingListRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreatePlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequest) GetDeviceInfo() *CreatePlayingListRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CreatePlayingListRequest) GetOpenCreatePlayingListRequest() *CreatePlayingListRequestOpenCreatePlayingListRequest {
	return s.OpenCreatePlayingListRequest
}

func (s *CreatePlayingListRequest) GetUserInfo() *CreatePlayingListRequestUserInfo {
	return s.UserInfo
}

func (s *CreatePlayingListRequest) SetDeviceInfo(v *CreatePlayingListRequestDeviceInfo) *CreatePlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreatePlayingListRequest) SetOpenCreatePlayingListRequest(v *CreatePlayingListRequestOpenCreatePlayingListRequest) *CreatePlayingListRequest {
	s.OpenCreatePlayingListRequest = v
	return s
}

func (s *CreatePlayingListRequest) SetUserInfo(v *CreatePlayingListRequestUserInfo) *CreatePlayingListRequest {
	s.UserInfo = v
	return s
}

func (s *CreatePlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListRequestDeviceInfo struct {
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
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreatePlayingListRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreatePlayingListRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreatePlayingListRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CreatePlayingListRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreatePlayingListRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreatePlayingListRequestDeviceInfo) SetEncodeKey(v string) *CreatePlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetEncodeType(v string) *CreatePlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetId(v string) *CreatePlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetIdType(v string) *CreatePlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) SetOrganizationId(v string) *CreatePlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreatePlayingListRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListRequestOpenCreatePlayingListRequest struct {
	// This parameter is required.
	ContentList []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// content
	ContentType *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ExtendInfo  map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// true
	NeedAlbumContinued *bool `json:"NeedAlbumContinued,omitempty" xml:"NeedAlbumContinued,omitempty"`
	// example:
	//
	// default
	PlayFrom *string `json:"PlayFrom,omitempty" xml:"PlayFrom,omitempty"`
	// example:
	//
	// Repeat
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetContentList() []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList {
	return s.ContentList
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetIndex() *int32 {
	return s.Index
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetNeedAlbumContinued() *bool {
	return s.NeedAlbumContinued
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetPlayFrom() *string {
	return s.PlayFrom
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetContentList(v []*CreatePlayingListRequestOpenCreatePlayingListRequestContentList) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ContentList = v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetContentType(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetExtendInfo(v map[string]interface{}) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.ExtendInfo = v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetIndex(v int32) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.Index = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetNeedAlbumContinued(v bool) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.NeedAlbumContinued = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetPlayFrom(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.PlayFrom = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) SetPlayMode(v string) *CreatePlayingListRequestOpenCreatePlayingListRequest {
	s.PlayMode = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListRequestOpenCreatePlayingListRequestContentList struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xiami
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequestContentList) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListRequestOpenCreatePlayingListRequestContentList) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) GetRawId() *string {
	return s.RawId
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) GetSource() *string {
	return s.Source
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) SetRawId(v string) *CreatePlayingListRequestOpenCreatePlayingListRequestContentList {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) SetSource(v string) *CreatePlayingListRequestOpenCreatePlayingListRequestContentList {
	s.Source = &v
	return s
}

func (s *CreatePlayingListRequestOpenCreatePlayingListRequestContentList) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListRequestUserInfo struct {
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
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreatePlayingListRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreatePlayingListRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreatePlayingListRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreatePlayingListRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CreatePlayingListRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreatePlayingListRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreatePlayingListRequestUserInfo) SetEncodeKey(v string) *CreatePlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetEncodeType(v string) *CreatePlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetId(v string) *CreatePlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetIdType(v string) *CreatePlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) SetOrganizationId(v string) *CreatePlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreatePlayingListRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
