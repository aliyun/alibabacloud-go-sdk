// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumIsAddedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbumIdList(v []*string) *ListAlbumIsAddedRequest
	GetAlbumIdList() []*string
	SetDeviceInfo(v *ListAlbumIsAddedRequestDeviceInfo) *ListAlbumIsAddedRequest
	GetDeviceInfo() *ListAlbumIsAddedRequestDeviceInfo
	SetUserInfo(v *ListAlbumIsAddedRequestUserInfo) *ListAlbumIsAddedRequest
	GetUserInfo() *ListAlbumIsAddedRequestUserInfo
}

type ListAlbumIsAddedRequest struct {
	AlbumIdList []*string                          `json:"AlbumIdList,omitempty" xml:"AlbumIdList,omitempty" type:"Repeated"`
	DeviceInfo  *ListAlbumIsAddedRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo    *ListAlbumIsAddedRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListAlbumIsAddedRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequest) GetAlbumIdList() []*string {
	return s.AlbumIdList
}

func (s *ListAlbumIsAddedRequest) GetDeviceInfo() *ListAlbumIsAddedRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListAlbumIsAddedRequest) GetUserInfo() *ListAlbumIsAddedRequestUserInfo {
	return s.UserInfo
}

func (s *ListAlbumIsAddedRequest) SetAlbumIdList(v []*string) *ListAlbumIsAddedRequest {
	s.AlbumIdList = v
	return s
}

func (s *ListAlbumIsAddedRequest) SetDeviceInfo(v *ListAlbumIsAddedRequestDeviceInfo) *ListAlbumIsAddedRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListAlbumIsAddedRequest) SetUserInfo(v *ListAlbumIsAddedRequestUserInfo) *ListAlbumIsAddedRequest {
	s.UserInfo = v
	return s
}

func (s *ListAlbumIsAddedRequest) Validate() error {
	return dara.Validate(s)
}

type ListAlbumIsAddedRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlbumIsAddedRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlbumIsAddedRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlbumIsAddedRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListAlbumIsAddedRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlbumIsAddedRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetEncodeKey(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetEncodeType(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetId(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetIdType(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) SetOrganizationId(v string) *ListAlbumIsAddedRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlbumIsAddedRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlbumIsAddedRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListAlbumIsAddedRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlbumIsAddedRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlbumIsAddedRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListAlbumIsAddedRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlbumIsAddedRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlbumIsAddedRequestUserInfo) SetEncodeKey(v string) *ListAlbumIsAddedRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetEncodeType(v string) *ListAlbumIsAddedRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetId(v string) *ListAlbumIsAddedRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetIdType(v string) *ListAlbumIsAddedRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) SetOrganizationId(v string) *ListAlbumIsAddedRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlbumIsAddedRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
