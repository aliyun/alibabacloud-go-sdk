// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *QueryMusicTypeRequestDeviceInfo) *QueryMusicTypeRequest
	GetDeviceInfo() *QueryMusicTypeRequestDeviceInfo
	SetPayload(v *QueryMusicTypeRequestPayload) *QueryMusicTypeRequest
	GetPayload() *QueryMusicTypeRequestPayload
	SetUserInfo(v *QueryMusicTypeRequestUserInfo) *QueryMusicTypeRequest
	GetUserInfo() *QueryMusicTypeRequestUserInfo
}

type QueryMusicTypeRequest struct {
	// This parameter is required.
	DeviceInfo *QueryMusicTypeRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *QueryMusicTypeRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *QueryMusicTypeRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryMusicTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequest) GetDeviceInfo() *QueryMusicTypeRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *QueryMusicTypeRequest) GetPayload() *QueryMusicTypeRequestPayload {
	return s.Payload
}

func (s *QueryMusicTypeRequest) GetUserInfo() *QueryMusicTypeRequestUserInfo {
	return s.UserInfo
}

func (s *QueryMusicTypeRequest) SetDeviceInfo(v *QueryMusicTypeRequestDeviceInfo) *QueryMusicTypeRequest {
	s.DeviceInfo = v
	return s
}

func (s *QueryMusicTypeRequest) SetPayload(v *QueryMusicTypeRequestPayload) *QueryMusicTypeRequest {
	s.Payload = v
	return s
}

func (s *QueryMusicTypeRequest) SetUserInfo(v *QueryMusicTypeRequestUserInfo) *QueryMusicTypeRequest {
	s.UserInfo = v
	return s
}

func (s *QueryMusicTypeRequest) Validate() error {
	return dara.Validate(s)
}

type QueryMusicTypeRequestDeviceInfo struct {
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

func (s QueryMusicTypeRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryMusicTypeRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryMusicTypeRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *QueryMusicTypeRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *QueryMusicTypeRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeKey(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetIdType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetOrganizationId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type QueryMusicTypeRequestPayload struct {
}

func (s QueryMusicTypeRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestPayload) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestPayload) Validate() error {
	return dara.Validate(s)
}

type QueryMusicTypeRequestUserInfo struct {
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

func (s QueryMusicTypeRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryMusicTypeRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryMusicTypeRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *QueryMusicTypeRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *QueryMusicTypeRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeKey(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeType(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetId(v string) *QueryMusicTypeRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetIdType(v string) *QueryMusicTypeRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetOrganizationId(v string) *QueryMusicTypeRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
