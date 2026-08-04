// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeforeTime(v string) *ListUserMessageRequest
	GetBeforeTime() *string
	SetUserInfo(v *ListUserMessageRequestUserInfo) *ListUserMessageRequest
	GetUserInfo() *ListUserMessageRequestUserInfo
	SetLimit(v int32) *ListUserMessageRequest
	GetLimit() *int32
}

type ListUserMessageRequest struct {
	// After a specific point in time
	//
	// example:
	//
	// 2022-07-27 14:06:55.984
	BeforeTime *string `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *ListUserMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// Number of records to query
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s ListUserMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageRequest) GoString() string {
	return s.String()
}

func (s *ListUserMessageRequest) GetBeforeTime() *string {
	return s.BeforeTime
}

func (s *ListUserMessageRequest) GetUserInfo() *ListUserMessageRequestUserInfo {
	return s.UserInfo
}

func (s *ListUserMessageRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUserMessageRequest) SetBeforeTime(v string) *ListUserMessageRequest {
	s.BeforeTime = &v
	return s
}

func (s *ListUserMessageRequest) SetUserInfo(v *ListUserMessageRequestUserInfo) *ListUserMessageRequest {
	s.UserInfo = v
	return s
}

func (s *ListUserMessageRequest) SetLimit(v int32) *ListUserMessageRequest {
	s.Limit = &v
	return s
}

func (s *ListUserMessageRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserMessageRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILLID, this value is the application\\"s Skill ID. When the encoding type is PACKAGENAME, this value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type: - PACKAGENAME: APK package name, used for Android application client links - SKILLID: Skill ID, used for cloud-based links
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of user ID: - OPENID: Default user ID identifier - UNIONID: Organization-level user ID identifier, available only after an organization has been registered on the Maojing skill application Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListUserMessageRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListUserMessageRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListUserMessageRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListUserMessageRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListUserMessageRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListUserMessageRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListUserMessageRequestUserInfo) SetEncodeKey(v string) *ListUserMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetEncodeType(v string) *ListUserMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetId(v string) *ListUserMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetIdType(v string) *ListUserMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) SetOrganizationId(v string) *ListUserMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListUserMessageRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
