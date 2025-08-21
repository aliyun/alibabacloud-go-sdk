// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListOAuth2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *CreatePlayingListOAuth2RequestDeviceInfo) *CreatePlayingListOAuth2Request
	GetDeviceInfo() *CreatePlayingListOAuth2RequestDeviceInfo
	SetOpenCreatePlayingListRequest(v *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) *CreatePlayingListOAuth2Request
	GetOpenCreatePlayingListRequest() *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest
}

type CreatePlayingListOAuth2Request struct {
	// This parameter is required.
	DeviceInfo *CreatePlayingListOAuth2RequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenCreatePlayingListRequest *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest `json:"OpenCreatePlayingListRequest,omitempty" xml:"OpenCreatePlayingListRequest,omitempty" type:"Struct"`
}

func (s CreatePlayingListOAuth2Request) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2Request) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2Request) GetDeviceInfo() *CreatePlayingListOAuth2RequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CreatePlayingListOAuth2Request) GetOpenCreatePlayingListRequest() *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	return s.OpenCreatePlayingListRequest
}

func (s *CreatePlayingListOAuth2Request) SetDeviceInfo(v *CreatePlayingListOAuth2RequestDeviceInfo) *CreatePlayingListOAuth2Request {
	s.DeviceInfo = v
	return s
}

func (s *CreatePlayingListOAuth2Request) SetOpenCreatePlayingListRequest(v *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) *CreatePlayingListOAuth2Request {
	s.OpenCreatePlayingListRequest = v
	return s
}

func (s *CreatePlayingListOAuth2Request) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListOAuth2RequestDeviceInfo struct {
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
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreatePlayingListOAuth2RequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2RequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) SetEncodeKey(v string) *CreatePlayingListOAuth2RequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) SetEncodeType(v string) *CreatePlayingListOAuth2RequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) SetId(v string) *CreatePlayingListOAuth2RequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) SetIdType(v string) *CreatePlayingListOAuth2RequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) SetOrganizationId(v string) *CreatePlayingListOAuth2RequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest struct {
	// This parameter is required.
	ContentList []*CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
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

func (s CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetContentList() []*CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList {
	return s.ContentList
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetIndex() *int32 {
	return s.Index
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetNeedAlbumContinued() *bool {
	return s.NeedAlbumContinued
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetPlayFrom() *string {
	return s.PlayFrom
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetContentList(v []*CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.ContentList = v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetContentType(v string) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetExtendInfo(v map[string]interface{}) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.ExtendInfo = v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetIndex(v int32) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.Index = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetNeedAlbumContinued(v bool) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.NeedAlbumContinued = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetPlayFrom(v string) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.PlayFrom = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) SetPlayMode(v string) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest {
	s.PlayMode = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList struct {
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

func (s CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) GetRawId() *string {
	return s.RawId
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) GetSource() *string {
	return s.Source
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) SetRawId(v string) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList {
	s.RawId = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) SetSource(v string) *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList {
	s.Source = &v
	return s
}

func (s *CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList) Validate() error {
	return dara.Validate(s)
}
