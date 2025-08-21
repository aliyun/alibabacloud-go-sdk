// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *IndexControlPlayingListRequestDeviceInfo) *IndexControlPlayingListRequest
	GetDeviceInfo() *IndexControlPlayingListRequestDeviceInfo
	SetOpenIndexControlRequest(v *IndexControlPlayingListRequestOpenIndexControlRequest) *IndexControlPlayingListRequest
	GetOpenIndexControlRequest() *IndexControlPlayingListRequestOpenIndexControlRequest
	SetUserInfo(v *IndexControlPlayingListRequestUserInfo) *IndexControlPlayingListRequest
	GetUserInfo() *IndexControlPlayingListRequestUserInfo
}

type IndexControlPlayingListRequest struct {
	// This parameter is required.
	DeviceInfo *IndexControlPlayingListRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenIndexControlRequest *IndexControlPlayingListRequestOpenIndexControlRequest `json:"OpenIndexControlRequest,omitempty" xml:"OpenIndexControlRequest,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *IndexControlPlayingListRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s IndexControlPlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequest) GetDeviceInfo() *IndexControlPlayingListRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *IndexControlPlayingListRequest) GetOpenIndexControlRequest() *IndexControlPlayingListRequestOpenIndexControlRequest {
	return s.OpenIndexControlRequest
}

func (s *IndexControlPlayingListRequest) GetUserInfo() *IndexControlPlayingListRequestUserInfo {
	return s.UserInfo
}

func (s *IndexControlPlayingListRequest) SetDeviceInfo(v *IndexControlPlayingListRequestDeviceInfo) *IndexControlPlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *IndexControlPlayingListRequest) SetOpenIndexControlRequest(v *IndexControlPlayingListRequestOpenIndexControlRequest) *IndexControlPlayingListRequest {
	s.OpenIndexControlRequest = v
	return s
}

func (s *IndexControlPlayingListRequest) SetUserInfo(v *IndexControlPlayingListRequestUserInfo) *IndexControlPlayingListRequest {
	s.UserInfo = v
	return s
}

func (s *IndexControlPlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type IndexControlPlayingListRequestDeviceInfo struct {
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

func (s IndexControlPlayingListRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetIdType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type IndexControlPlayingListRequestOpenIndexControlRequest struct {
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// false
	NeedContentContinued *bool `json:"NeedContentContinued,omitempty" xml:"NeedContentContinued,omitempty"`
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetIndex() *int32 {
	return s.Index
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetNeedContentContinued() *bool {
	return s.NeedContentContinued
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetExtendInfo(v map[string]interface{}) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.ExtendInfo = v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetIndex(v int32) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.Index = &v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetNeedContentContinued(v bool) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.NeedContentContinued = &v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) Validate() error {
	return dara.Validate(s)
}

type IndexControlPlayingListRequestUserInfo struct {
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

func (s IndexControlPlayingListRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *IndexControlPlayingListRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *IndexControlPlayingListRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *IndexControlPlayingListRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *IndexControlPlayingListRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeType(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetId(v string) *IndexControlPlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetIdType(v string) *IndexControlPlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
