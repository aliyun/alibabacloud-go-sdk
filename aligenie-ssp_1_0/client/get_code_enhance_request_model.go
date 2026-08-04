// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeEnhanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelInfo(v *GetCodeEnhanceRequestChannelInfo) *GetCodeEnhanceRequest
	GetChannelInfo() *GetCodeEnhanceRequestChannelInfo
	SetUserInfo(v *GetCodeEnhanceRequestUserInfo) *GetCodeEnhanceRequest
	GetUserInfo() *GetCodeEnhanceRequestUserInfo
}

type GetCodeEnhanceRequest struct {
	// Activation Channel, such as WeChat mini program or third-party app.
	//
	// This parameter is required.
	ChannelInfo *GetCodeEnhanceRequestChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Struct"`
	// List of User Identifier information.
	//
	// This parameter is required.
	UserInfo *GetCodeEnhanceRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCodeEnhanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceRequest) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequest) GetChannelInfo() *GetCodeEnhanceRequestChannelInfo {
	return s.ChannelInfo
}

func (s *GetCodeEnhanceRequest) GetUserInfo() *GetCodeEnhanceRequestUserInfo {
	return s.UserInfo
}

func (s *GetCodeEnhanceRequest) SetChannelInfo(v *GetCodeEnhanceRequestChannelInfo) *GetCodeEnhanceRequest {
	s.ChannelInfo = v
	return s
}

func (s *GetCodeEnhanceRequest) SetUserInfo(v *GetCodeEnhanceRequestUserInfo) *GetCodeEnhanceRequest {
	s.UserInfo = v
	return s
}

func (s *GetCodeEnhanceRequest) Validate() error {
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

type GetCodeEnhanceRequestChannelInfo struct {
	// Activation Channel, such as WeChat or ThirdApp.
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

func (s GetCodeEnhanceRequestChannelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceRequestChannelInfo) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequestChannelInfo) GetChannel() *string {
	return s.Channel
}

func (s *GetCodeEnhanceRequestChannelInfo) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *GetCodeEnhanceRequestChannelInfo) SetChannel(v string) *GetCodeEnhanceRequestChannelInfo {
	s.Channel = &v
	return s
}

func (s *GetCodeEnhanceRequestChannelInfo) SetExtInfo(v string) *GetCodeEnhanceRequestChannelInfo {
	s.ExtInfo = &v
	return s
}

func (s *GetCodeEnhanceRequestChannelInfo) Validate() error {
	return dara.Validate(s)
}

type GetCodeEnhanceRequestUserInfo struct {
	// The value corresponding to the encoding type. Enter the Project ID of the project to which the product belongs. You can view it in the Tmall Genie AI Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding Type. Enter PROJECT_ID here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier. Enter the value of userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID:
	//
	// OPENID: The default User ID identifier.
	//
	// UNIONID: The organization-dimension User ID identifier. You must request an organization in advance on the Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 1
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCodeEnhanceRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCodeEnhanceRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCodeEnhanceRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetCodeEnhanceRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCodeEnhanceRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCodeEnhanceRequestUserInfo) SetEncodeKey(v string) *GetCodeEnhanceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetEncodeType(v string) *GetCodeEnhanceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetId(v string) *GetCodeEnhanceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetIdType(v string) *GetCodeEnhanceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) SetOrganizationId(v string) *GetCodeEnhanceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCodeEnhanceRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
