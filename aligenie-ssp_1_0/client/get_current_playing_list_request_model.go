// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetCurrentPlayingListRequestDeviceInfo) *GetCurrentPlayingListRequest
	GetDeviceInfo() *GetCurrentPlayingListRequestDeviceInfo
	SetOpenQueryPlayListRequest(v *GetCurrentPlayingListRequestOpenQueryPlayListRequest) *GetCurrentPlayingListRequest
	GetOpenQueryPlayListRequest() *GetCurrentPlayingListRequestOpenQueryPlayListRequest
	SetUserInfo(v *GetCurrentPlayingListRequestUserInfo) *GetCurrentPlayingListRequest
	GetUserInfo() *GetCurrentPlayingListRequestUserInfo
}

type GetCurrentPlayingListRequest struct {
	// This parameter is required.
	DeviceInfo *GetCurrentPlayingListRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	OpenQueryPlayListRequest *GetCurrentPlayingListRequestOpenQueryPlayListRequest `json:"OpenQueryPlayListRequest,omitempty" xml:"OpenQueryPlayListRequest,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetCurrentPlayingListRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCurrentPlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequest) GetDeviceInfo() *GetCurrentPlayingListRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetCurrentPlayingListRequest) GetOpenQueryPlayListRequest() *GetCurrentPlayingListRequestOpenQueryPlayListRequest {
	return s.OpenQueryPlayListRequest
}

func (s *GetCurrentPlayingListRequest) GetUserInfo() *GetCurrentPlayingListRequestUserInfo {
	return s.UserInfo
}

func (s *GetCurrentPlayingListRequest) SetDeviceInfo(v *GetCurrentPlayingListRequestDeviceInfo) *GetCurrentPlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetCurrentPlayingListRequest) SetOpenQueryPlayListRequest(v *GetCurrentPlayingListRequestOpenQueryPlayListRequest) *GetCurrentPlayingListRequest {
	s.OpenQueryPlayListRequest = v
	return s
}

func (s *GetCurrentPlayingListRequest) SetUserInfo(v *GetCurrentPlayingListRequestUserInfo) *GetCurrentPlayingListRequest {
	s.UserInfo = v
	return s
}

func (s *GetCurrentPlayingListRequest) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingListRequestDeviceInfo struct {
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

func (s GetCurrentPlayingListRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingListRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingListRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingListRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingListRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetEncodeKey(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetEncodeType(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetId(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetIdType(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) SetOrganizationId(v string) *GetCurrentPlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingListRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingListRequestOpenQueryPlayListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCurrentPlayingListRequestOpenQueryPlayListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListRequestOpenQueryPlayListRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) SetPageNum(v int32) *GetCurrentPlayingListRequestOpenQueryPlayListRequest {
	s.PageNum = &v
	return s
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) SetPageSize(v int32) *GetCurrentPlayingListRequestOpenQueryPlayListRequest {
	s.PageSize = &v
	return s
}

func (s *GetCurrentPlayingListRequestOpenQueryPlayListRequest) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingListRequestUserInfo struct {
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

func (s GetCurrentPlayingListRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingListRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingListRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingListRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingListRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingListRequestUserInfo) SetEncodeKey(v string) *GetCurrentPlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetEncodeType(v string) *GetCurrentPlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetId(v string) *GetCurrentPlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetIdType(v string) *GetCurrentPlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) SetOrganizationId(v string) *GetCurrentPlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingListRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
