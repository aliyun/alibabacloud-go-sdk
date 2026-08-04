// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecommendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *MobileRecommendRequest
	GetBotId() *string
	SetCount(v string) *MobileRecommendRequest
	GetCount() *string
	SetDeviceInfo(v *MobileRecommendRequestDeviceInfo) *MobileRecommendRequest
	GetDeviceInfo() *MobileRecommendRequestDeviceInfo
	SetStyle(v string) *MobileRecommendRequest
	GetStyle() *string
	SetType(v string) *MobileRecommendRequest
	GetType() *string
	SetUserInfo(v *MobileRecommendRequestUserInfo) *MobileRecommendRequest
	GetUserInfo() *MobileRecommendRequestUserInfo
}

type MobileRecommendRequest struct {
	// Bot ID.
	//
	// example:
	//
	// 10
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Quantity of recommended Result
	//
	// example:
	//
	// 6
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// Device identification information.
	//
	// This parameter is required.
	DeviceInfo *MobileRecommendRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Required when the request type is STYLE.
	//
	// example:
	//
	// 轻音乐
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// Request Type: Obtain daily recommendations, hot songs, or genre-based playlists.
	//
	// example:
	//
	// DAILY_REC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// User information – userId
	//
	// This parameter is required.
	UserInfo *MobileRecommendRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s MobileRecommendRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendRequest) GoString() string {
	return s.String()
}

func (s *MobileRecommendRequest) GetBotId() *string {
	return s.BotId
}

func (s *MobileRecommendRequest) GetCount() *string {
	return s.Count
}

func (s *MobileRecommendRequest) GetDeviceInfo() *MobileRecommendRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *MobileRecommendRequest) GetStyle() *string {
	return s.Style
}

func (s *MobileRecommendRequest) GetType() *string {
	return s.Type
}

func (s *MobileRecommendRequest) GetUserInfo() *MobileRecommendRequestUserInfo {
	return s.UserInfo
}

func (s *MobileRecommendRequest) SetBotId(v string) *MobileRecommendRequest {
	s.BotId = &v
	return s
}

func (s *MobileRecommendRequest) SetCount(v string) *MobileRecommendRequest {
	s.Count = &v
	return s
}

func (s *MobileRecommendRequest) SetDeviceInfo(v *MobileRecommendRequestDeviceInfo) *MobileRecommendRequest {
	s.DeviceInfo = v
	return s
}

func (s *MobileRecommendRequest) SetStyle(v string) *MobileRecommendRequest {
	s.Style = &v
	return s
}

func (s *MobileRecommendRequest) SetType(v string) *MobileRecommendRequest {
	s.Type = &v
	return s
}

func (s *MobileRecommendRequest) SetUserInfo(v *MobileRecommendRequestUserInfo) *MobileRecommendRequest {
	s.UserInfo = v
	return s
}

func (s *MobileRecommendRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MobileRecommendRequestDeviceInfo struct {
	// Value corresponding to the encoding type. Enter the Project ID of the project to which the product belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1414895629783187053
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
	// fjwZiYQdtkaI95fHaLNjYcaOA/mxUPzxxw2J5iBiTBnjUCWKwER4TSHCqkBnNOYvGJ4bRZA9KzBB2naS4r/Am0lSe8ECDAAOcJ9QKLFF6DM=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID.
	//
	// OPEN_ID: Default device ID.
	//
	// UNION_ID: Organization-level device ID. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required when **IdType*	- is set to **UNION_ID**.
	//
	// example:
	//
	// 暂无
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s MobileRecommendRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *MobileRecommendRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *MobileRecommendRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *MobileRecommendRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *MobileRecommendRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *MobileRecommendRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *MobileRecommendRequestDeviceInfo) SetEncodeKey(v string) *MobileRecommendRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *MobileRecommendRequestDeviceInfo) SetEncodeType(v string) *MobileRecommendRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *MobileRecommendRequestDeviceInfo) SetId(v string) *MobileRecommendRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *MobileRecommendRequestDeviceInfo) SetIdType(v string) *MobileRecommendRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *MobileRecommendRequestDeviceInfo) SetOrganizationId(v string) *MobileRecommendRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *MobileRecommendRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type MobileRecommendRequestUserInfo struct {
	// The value corresponding to the encoding Type. Enter the Project ID of the project to which this product belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1414895629783187053
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// ID value
	//
	// This parameter is required.
	//
	// example:
	//
	// R457Av3qg/OXTwVnFt12z6MwNe0HAS699V6n63OaLdu+VmwvhcNfMzBd+la553wWJhj3kBMjgHq2Y2dyCFoDBg==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// ID Type
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is Required when IdType is set to UNION_ID.
	//
	// example:
	//
	// 暂无
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s MobileRecommendRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendRequestUserInfo) GoString() string {
	return s.String()
}

func (s *MobileRecommendRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *MobileRecommendRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *MobileRecommendRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *MobileRecommendRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *MobileRecommendRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *MobileRecommendRequestUserInfo) SetEncodeKey(v string) *MobileRecommendRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *MobileRecommendRequestUserInfo) SetEncodeType(v string) *MobileRecommendRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *MobileRecommendRequestUserInfo) SetId(v string) *MobileRecommendRequestUserInfo {
	s.Id = &v
	return s
}

func (s *MobileRecommendRequestUserInfo) SetIdType(v string) *MobileRecommendRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *MobileRecommendRequestUserInfo) SetOrganizationId(v string) *MobileRecommendRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *MobileRecommendRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
