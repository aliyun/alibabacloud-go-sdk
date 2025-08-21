// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudPlayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurPlayIndex(v int32) *CloudPlayerRequest
	GetCurPlayIndex() *int32
	SetDeviceInfo(v *CloudPlayerRequestDeviceInfo) *CloudPlayerRequest
	GetDeviceInfo() *CloudPlayerRequestDeviceInfo
	SetPlayMode(v string) *CloudPlayerRequest
	GetPlayMode() *string
	SetSongId(v string) *CloudPlayerRequest
	GetSongId() *string
	SetSongIdList(v []*string) *CloudPlayerRequest
	GetSongIdList() []*string
	SetSource(v string) *CloudPlayerRequest
	GetSource() *string
	SetUserInfo(v *CloudPlayerRequestUserInfo) *CloudPlayerRequest
	GetUserInfo() *CloudPlayerRequestUserInfo
}

type CloudPlayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurPlayIndex *int32 `json:"CurPlayIndex,omitempty" xml:"CurPlayIndex,omitempty"`
	// This parameter is required.
	DeviceInfo *CloudPlayerRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
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
	SongIdList []*string `json:"SongIdList,omitempty" xml:"SongIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// KG
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	UserInfo *CloudPlayerRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CloudPlayerRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerRequest) GoString() string {
	return s.String()
}

func (s *CloudPlayerRequest) GetCurPlayIndex() *int32 {
	return s.CurPlayIndex
}

func (s *CloudPlayerRequest) GetDeviceInfo() *CloudPlayerRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CloudPlayerRequest) GetPlayMode() *string {
	return s.PlayMode
}

func (s *CloudPlayerRequest) GetSongId() *string {
	return s.SongId
}

func (s *CloudPlayerRequest) GetSongIdList() []*string {
	return s.SongIdList
}

func (s *CloudPlayerRequest) GetSource() *string {
	return s.Source
}

func (s *CloudPlayerRequest) GetUserInfo() *CloudPlayerRequestUserInfo {
	return s.UserInfo
}

func (s *CloudPlayerRequest) SetCurPlayIndex(v int32) *CloudPlayerRequest {
	s.CurPlayIndex = &v
	return s
}

func (s *CloudPlayerRequest) SetDeviceInfo(v *CloudPlayerRequestDeviceInfo) *CloudPlayerRequest {
	s.DeviceInfo = v
	return s
}

func (s *CloudPlayerRequest) SetPlayMode(v string) *CloudPlayerRequest {
	s.PlayMode = &v
	return s
}

func (s *CloudPlayerRequest) SetSongId(v string) *CloudPlayerRequest {
	s.SongId = &v
	return s
}

func (s *CloudPlayerRequest) SetSongIdList(v []*string) *CloudPlayerRequest {
	s.SongIdList = v
	return s
}

func (s *CloudPlayerRequest) SetSource(v string) *CloudPlayerRequest {
	s.Source = &v
	return s
}

func (s *CloudPlayerRequest) SetUserInfo(v *CloudPlayerRequestUserInfo) *CloudPlayerRequest {
	s.UserInfo = v
	return s
}

func (s *CloudPlayerRequest) Validate() error {
	return dara.Validate(s)
}

type CloudPlayerRequestDeviceInfo struct {
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
	// 1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CloudPlayerRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CloudPlayerRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CloudPlayerRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CloudPlayerRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CloudPlayerRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CloudPlayerRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CloudPlayerRequestDeviceInfo) SetEncodeKey(v string) *CloudPlayerRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CloudPlayerRequestDeviceInfo) SetEncodeType(v string) *CloudPlayerRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CloudPlayerRequestDeviceInfo) SetId(v string) *CloudPlayerRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CloudPlayerRequestDeviceInfo) SetIdType(v string) *CloudPlayerRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CloudPlayerRequestDeviceInfo) SetOrganizationId(v string) *CloudPlayerRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CloudPlayerRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CloudPlayerRequestUserInfo struct {
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
	// 1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CloudPlayerRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CloudPlayerRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CloudPlayerRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CloudPlayerRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CloudPlayerRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CloudPlayerRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CloudPlayerRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CloudPlayerRequestUserInfo) SetEncodeKey(v string) *CloudPlayerRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CloudPlayerRequestUserInfo) SetEncodeType(v string) *CloudPlayerRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CloudPlayerRequestUserInfo) SetId(v string) *CloudPlayerRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CloudPlayerRequestUserInfo) SetIdType(v string) *CloudPlayerRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CloudPlayerRequestUserInfo) SetOrganizationId(v string) *CloudPlayerRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CloudPlayerRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
