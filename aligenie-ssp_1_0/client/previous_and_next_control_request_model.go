// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviousAndNextControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PreviousAndNextControlRequestDeviceInfo) *PreviousAndNextControlRequest
	GetDeviceInfo() *PreviousAndNextControlRequestDeviceInfo
	SetOpenControlPlayingListRequest(v *PreviousAndNextControlRequestOpenControlPlayingListRequest) *PreviousAndNextControlRequest
	GetOpenControlPlayingListRequest() *PreviousAndNextControlRequestOpenControlPlayingListRequest
	SetUserInfo(v *PreviousAndNextControlRequestUserInfo) *PreviousAndNextControlRequest
	GetUserInfo() *PreviousAndNextControlRequestUserInfo
}

type PreviousAndNextControlRequest struct {
	// This parameter is required.
	DeviceInfo *PreviousAndNextControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenControlPlayingListRequest *PreviousAndNextControlRequestOpenControlPlayingListRequest `json:"OpenControlPlayingListRequest,omitempty" xml:"OpenControlPlayingListRequest,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *PreviousAndNextControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PreviousAndNextControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequest) GetDeviceInfo() *PreviousAndNextControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PreviousAndNextControlRequest) GetOpenControlPlayingListRequest() *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	return s.OpenControlPlayingListRequest
}

func (s *PreviousAndNextControlRequest) GetUserInfo() *PreviousAndNextControlRequestUserInfo {
	return s.UserInfo
}

func (s *PreviousAndNextControlRequest) SetDeviceInfo(v *PreviousAndNextControlRequestDeviceInfo) *PreviousAndNextControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PreviousAndNextControlRequest) SetOpenControlPlayingListRequest(v *PreviousAndNextControlRequestOpenControlPlayingListRequest) *PreviousAndNextControlRequest {
	s.OpenControlPlayingListRequest = v
	return s
}

func (s *PreviousAndNextControlRequest) SetUserInfo(v *PreviousAndNextControlRequestUserInfo) *PreviousAndNextControlRequest {
	s.UserInfo = v
	return s
}

func (s *PreviousAndNextControlRequest) Validate() error {
	return dara.Validate(s)
}

type PreviousAndNextControlRequestDeviceInfo struct {
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

func (s PreviousAndNextControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PreviousAndNextControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PreviousAndNextControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PreviousAndNextControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PreviousAndNextControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetEncodeKey(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetEncodeType(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetId(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetIdType(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) SetOrganizationId(v string) *PreviousAndNextControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PreviousAndNextControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PreviousAndNextControlRequestOpenControlPlayingListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NEXT
	Cmd        *string                `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// false
	IsFromDevice *bool `json:"IsFromDevice,omitempty" xml:"IsFromDevice,omitempty"`
}

func (s PreviousAndNextControlRequestOpenControlPlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlRequestOpenControlPlayingListRequest) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) GetCmd() *string {
	return s.Cmd
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) GetIsFromDevice() *bool {
	return s.IsFromDevice
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetCmd(v string) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.Cmd = &v
	return s
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetExtendInfo(v map[string]interface{}) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.ExtendInfo = v
	return s
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) SetIsFromDevice(v bool) *PreviousAndNextControlRequestOpenControlPlayingListRequest {
	s.IsFromDevice = &v
	return s
}

func (s *PreviousAndNextControlRequestOpenControlPlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type PreviousAndNextControlRequestUserInfo struct {
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

func (s PreviousAndNextControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PreviousAndNextControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PreviousAndNextControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PreviousAndNextControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PreviousAndNextControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PreviousAndNextControlRequestUserInfo) SetEncodeKey(v string) *PreviousAndNextControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetEncodeType(v string) *PreviousAndNextControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetId(v string) *PreviousAndNextControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetIdType(v string) *PreviousAndNextControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) SetOrganizationId(v string) *PreviousAndNextControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PreviousAndNextControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
