// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PlayModeControlRequestDeviceInfo) *PlayModeControlRequest
	GetDeviceInfo() *PlayModeControlRequestDeviceInfo
	SetOpenPlayModeControlRequest(v *PlayModeControlRequestOpenPlayModeControlRequest) *PlayModeControlRequest
	GetOpenPlayModeControlRequest() *PlayModeControlRequestOpenPlayModeControlRequest
	SetUserInfo(v *PlayModeControlRequestUserInfo) *PlayModeControlRequest
	GetUserInfo() *PlayModeControlRequestUserInfo
}

type PlayModeControlRequest struct {
	// This parameter is required.
	DeviceInfo *PlayModeControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenPlayModeControlRequest *PlayModeControlRequestOpenPlayModeControlRequest `json:"OpenPlayModeControlRequest,omitempty" xml:"OpenPlayModeControlRequest,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *PlayModeControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PlayModeControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequest) GetDeviceInfo() *PlayModeControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PlayModeControlRequest) GetOpenPlayModeControlRequest() *PlayModeControlRequestOpenPlayModeControlRequest {
	return s.OpenPlayModeControlRequest
}

func (s *PlayModeControlRequest) GetUserInfo() *PlayModeControlRequestUserInfo {
	return s.UserInfo
}

func (s *PlayModeControlRequest) SetDeviceInfo(v *PlayModeControlRequestDeviceInfo) *PlayModeControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PlayModeControlRequest) SetOpenPlayModeControlRequest(v *PlayModeControlRequestOpenPlayModeControlRequest) *PlayModeControlRequest {
	s.OpenPlayModeControlRequest = v
	return s
}

func (s *PlayModeControlRequest) SetUserInfo(v *PlayModeControlRequestUserInfo) *PlayModeControlRequest {
	s.UserInfo = v
	return s
}

func (s *PlayModeControlRequest) Validate() error {
	return dara.Validate(s)
}

type PlayModeControlRequestDeviceInfo struct {
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

func (s PlayModeControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayModeControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayModeControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PlayModeControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayModeControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeKey(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeType(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetId(v string) *PlayModeControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetIdType(v string) *PlayModeControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetOrganizationId(v string) *PlayModeControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PlayModeControlRequestOpenPlayModeControlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Normal
	OpenPlayMode *string `json:"OpenPlayMode,omitempty" xml:"OpenPlayMode,omitempty"`
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) GetOpenPlayMode() *string {
	return s.OpenPlayMode
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) SetOpenPlayMode(v string) *PlayModeControlRequestOpenPlayModeControlRequest {
	s.OpenPlayMode = &v
	return s
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) Validate() error {
	return dara.Validate(s)
}

type PlayModeControlRequestUserInfo struct {
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

func (s PlayModeControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayModeControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayModeControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PlayModeControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayModeControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayModeControlRequestUserInfo) SetEncodeKey(v string) *PlayModeControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetEncodeType(v string) *PlayModeControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetId(v string) *PlayModeControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetIdType(v string) *PlayModeControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetOrganizationId(v string) *PlayModeControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
