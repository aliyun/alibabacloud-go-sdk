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
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *CreatePlayingListOAuth2RequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
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
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenCreatePlayingListRequest != nil {
		if err := s.OpenCreatePlayingListRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePlayingListOAuth2RequestDeviceInfo struct {
	// The value corresponding to the encoding type. Enter the Project ID of the project to which the product belongs. You can view it in the Tmall Genie AI Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Enter PROJECT_ID here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID. Enter the value of deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of device ID:
	//
	// OPEN_ID: The default device ID.
	//
	// UNION_ID: The organization-level device ID. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID.
	//
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
	// Playback objects
	//
	// This parameter is required.
	ContentList []*CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// Content type for playback
	//
	// Content: content; Album: album; Playlist: collect
	//
	// This parameter is required.
	//
	// example:
	//
	// content
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Extension information
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Index of the item to play
	//
	// Can be empty. Default is 0, which means playback starts from the beginning.
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Indicates whether album playback should continue from the last played episode. For example, if the last playback stopped at episode 5, whether to resume from episode 5. Default is true.
	//
	// example:
	//
	// true
	NeedAlbumContinued *bool `json:"NeedAlbumContinued,omitempty" xml:"NeedAlbumContinued,omitempty"`
	// Playback source, the unique identifier for configuring playback control capabilities.
	//
	// Optional. Default value is "default".
	//
	// example:
	//
	// default
	PlayFrom *string `json:"PlayFrom,omitempty" xml:"PlayFrom,omitempty"`
	// Playback pattern
	//
	// Repeat all: Repeat; Shuffle: Shuffle; Repeat one: RepeatOne; Play in order: Normal.
	//
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
	if s.ContentList != nil {
		for _, item := range s.ContentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePlayingListOAuth2RequestOpenCreatePlayingListRequestContentList struct {
	// Third-party ID.
	//
	// If the item is content, this is the content ID; if it is an album, this is the album ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Source
	//
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
