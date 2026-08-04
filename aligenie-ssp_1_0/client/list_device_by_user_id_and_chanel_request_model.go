// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdAndChanelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelInfo(v *ListDeviceByUserIdAndChanelRequestChannelInfo) *ListDeviceByUserIdAndChanelRequest
	GetChannelInfo() *ListDeviceByUserIdAndChanelRequestChannelInfo
	SetUserInfo(v *ListDeviceByUserIdAndChanelRequestUserInfo) *ListDeviceByUserIdAndChanelRequest
	GetUserInfo() *ListDeviceByUserIdAndChanelRequestUserInfo
}

type ListDeviceByUserIdAndChanelRequest struct {
	// Activation channel, such as WeChat mini program or third-party app.
	//
	// This parameter is required.
	ChannelInfo *ListDeviceByUserIdAndChanelRequestChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Struct"`
	// List of User Identifier information.
	//
	// This parameter is required.
	UserInfo *ListDeviceByUserIdAndChanelRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListDeviceByUserIdAndChanelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequest) GetChannelInfo() *ListDeviceByUserIdAndChanelRequestChannelInfo {
	return s.ChannelInfo
}

func (s *ListDeviceByUserIdAndChanelRequest) GetUserInfo() *ListDeviceByUserIdAndChanelRequestUserInfo {
	return s.UserInfo
}

func (s *ListDeviceByUserIdAndChanelRequest) SetChannelInfo(v *ListDeviceByUserIdAndChanelRequestChannelInfo) *ListDeviceByUserIdAndChanelRequest {
	s.ChannelInfo = v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequest) SetUserInfo(v *ListDeviceByUserIdAndChanelRequestUserInfo) *ListDeviceByUserIdAndChanelRequest {
	s.UserInfo = v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequest) Validate() error {
	if s.ChannelInfo != nil {
		if err := s.ChannelInfo.Validate(); err != nil {
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

type ListDeviceByUserIdAndChanelRequestChannelInfo struct {
	// Activation provisioning channel, such as WeChat or ThirdApp.
	//
	// This parameter is required.
	//
	// example:
	//
	// WeChat、ThirdApp
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Extension information.
	//
	// example:
	//
	// {}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s ListDeviceByUserIdAndChanelRequestChannelInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequestChannelInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) GetChannel() *string {
	return s.Channel
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) SetChannel(v string) *ListDeviceByUserIdAndChanelRequestChannelInfo {
	s.Channel = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) SetExtInfo(v string) *ListDeviceByUserIdAndChanelRequestChannelInfo {
	s.ExtInfo = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestChannelInfo) Validate() error {
	return dara.Validate(s)
}

type ListDeviceByUserIdAndChanelRequestUserInfo struct {
	// The value corresponding to the encoding type; enter the Project ID of the project to which the product belongs. You can view this in the Tmall Genie AI platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding Type; enter PROJECT_ID here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier; enter the value of userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Type of the User ID:
	//
	// OPEN_ID: The default User ID identity.
	//
	// UNION_ID: The User ID identity at the organization dimension, which requires you to request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID; required if IdType is UNION_ID.
	//
	// example:
	//
	// 1***2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdAndChanelRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetEncodeKey(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetEncodeType(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetId(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetIdType(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) SetOrganizationId(v string) *ListDeviceByUserIdAndChanelRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
