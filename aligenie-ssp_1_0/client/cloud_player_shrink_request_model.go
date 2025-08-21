// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPlayerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurPlayIndex(v int32) *CloudPlayerShrinkRequest
	GetCurPlayIndex() *int32
	SetDeviceInfoShrink(v string) *CloudPlayerShrinkRequest
	GetDeviceInfoShrink() *string
	SetPlayMode(v string) *CloudPlayerShrinkRequest
	GetPlayMode() *string
	SetSongId(v string) *CloudPlayerShrinkRequest
	GetSongId() *string
	SetSongIdListShrink(v string) *CloudPlayerShrinkRequest
	GetSongIdListShrink() *string
	SetSource(v string) *CloudPlayerShrinkRequest
	GetSource() *string
	SetUserInfoShrink(v string) *CloudPlayerShrinkRequest
	GetUserInfoShrink() *string
}

type CloudPlayerShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurPlayIndex *int32 `json:"CurPlayIndex,omitempty" xml:"CurPlayIndex,omitempty"`
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Normal
	PlayMode *string `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	// example:
	//
	// 123
	SongId *string `json:"SongId,omitempty" xml:"SongId,omitempty"`
	// This parameter is required.
	SongIdListShrink *string `json:"SongIdList,omitempty" xml:"SongIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KG
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CloudPlayerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloudPlayerShrinkRequest) GetCurPlayIndex() *int32 {
	return s.CurPlayIndex
}

func (s *CloudPlayerShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CloudPlayerShrinkRequest) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CloudPlayerShrinkRequest) GetSongId() *string {
	return s.SongId
}

func (s *CloudPlayerShrinkRequest) GetSongIdListShrink() *string {
	return s.SongIdListShrink
}

func (s *CloudPlayerShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CloudPlayerShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CloudPlayerShrinkRequest) SetCurPlayIndex(v int32) *CloudPlayerShrinkRequest {
	s.CurPlayIndex = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetDeviceInfoShrink(v string) *CloudPlayerShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetPlayMode(v string) *CloudPlayerShrinkRequest {
	s.PlayMode = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetSongId(v string) *CloudPlayerShrinkRequest {
	s.SongId = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetSongIdListShrink(v string) *CloudPlayerShrinkRequest {
	s.SongIdListShrink = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetSource(v string) *CloudPlayerShrinkRequest {
	s.Source = &v
	return s
}

func (s *CloudPlayerShrinkRequest) SetUserInfoShrink(v string) *CloudPlayerShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CloudPlayerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
