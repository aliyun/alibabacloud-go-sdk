// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoAppReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *VideoAppReportRequestDeviceInfo) *VideoAppReportRequest
	GetDeviceInfo() *VideoAppReportRequestDeviceInfo
	SetPayload(v *VideoAppReportRequestPayload) *VideoAppReportRequest
	GetPayload() *VideoAppReportRequestPayload
	SetUserInfo(v *VideoAppReportRequestUserInfo) *VideoAppReportRequest
	GetUserInfo() *VideoAppReportRequestUserInfo
}

type VideoAppReportRequest struct {
	// Device identity information
	DeviceInfo *VideoAppReportRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// 业务参数
	Payload *VideoAppReportRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// 用户标识信息
	//
	// This parameter is required.
	UserInfo *VideoAppReportRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s VideoAppReportRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportRequest) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequest) GetDeviceInfo() *VideoAppReportRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *VideoAppReportRequest) GetPayload() *VideoAppReportRequestPayload {
	return s.Payload
}

func (s *VideoAppReportRequest) GetUserInfo() *VideoAppReportRequestUserInfo {
	return s.UserInfo
}

func (s *VideoAppReportRequest) SetDeviceInfo(v *VideoAppReportRequestDeviceInfo) *VideoAppReportRequest {
	s.DeviceInfo = v
	return s
}

func (s *VideoAppReportRequest) SetPayload(v *VideoAppReportRequestPayload) *VideoAppReportRequest {
	s.Payload = v
	return s
}

func (s *VideoAppReportRequest) SetUserInfo(v *VideoAppReportRequestUserInfo) *VideoAppReportRequest {
	s.UserInfo = v
	return s
}

func (s *VideoAppReportRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
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

type VideoAppReportRequestDeviceInfo struct {
	// 编码类型对应的值
	//
	// - 编码类型是SKILL_ID时，其值为应用的Skill ID。
	//
	// - 编码类型是PACKAGE_NAME时，其值为对应客户端App的packageName。
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID from Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used for Android application customer flows.
	//
	// - SKILL_ID: Skill ID, used for cloud-based flows.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, set to either deviceOpenId or deviceUnionId.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID.
	//
	// - OPEN_ID: The default device ID.
	//
	// - UNION_ID: Organization-level device ID. This value is available only after an organization has been registered on the Tmall Genie Skill Application Open Platform.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s VideoAppReportRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *VideoAppReportRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *VideoAppReportRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *VideoAppReportRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *VideoAppReportRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *VideoAppReportRequestDeviceInfo) SetEncodeKey(v string) *VideoAppReportRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetEncodeType(v string) *VideoAppReportRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetId(v string) *VideoAppReportRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetIdType(v string) *VideoAppReportRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetOrganizationId(v string) *VideoAppReportRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type VideoAppReportRequestPayload struct {
	// 会员的结束时间。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1652337963097
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 会员是否登录。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsLogin *bool `json:"isLogin,omitempty" xml:"isLogin,omitempty"`
	// 该账号是否为会员。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsVip *bool `json:"isVip,omitempty" xml:"isVip,omitempty"`
	// 账号的昵称。
	//
	// example:
	//
	// test
	LoginNick *string `json:"loginNick,omitempty" xml:"loginNick,omitempty"`
	// 明文uuid
	//
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
	// 账户的手机号码。
	//
	// example:
	//
	// 188*777
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 应用包名。
	//
	// This parameter is required.
	//
	// example:
	//
	// com.***.test
	PkgName *string `json:"pkgName,omitempty" xml:"pkgName,omitempty"`
	// 会员的开始时间。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1652337963097
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s VideoAppReportRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportRequestPayload) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestPayload) GetEndTime() *int64 {
	return s.EndTime
}

func (s *VideoAppReportRequestPayload) GetIsLogin() *bool {
	return s.IsLogin
}

func (s *VideoAppReportRequestPayload) GetIsVip() *bool {
	return s.IsVip
}

func (s *VideoAppReportRequestPayload) GetLoginNick() *string {
	return s.LoginNick
}

func (s *VideoAppReportRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *VideoAppReportRequestPayload) GetPhone() *string {
	return s.Phone
}

func (s *VideoAppReportRequestPayload) GetPkgName() *string {
	return s.PkgName
}

func (s *VideoAppReportRequestPayload) GetStartTime() *int64 {
	return s.StartTime
}

func (s *VideoAppReportRequestPayload) SetEndTime(v int64) *VideoAppReportRequestPayload {
	s.EndTime = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetIsLogin(v bool) *VideoAppReportRequestPayload {
	s.IsLogin = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetIsVip(v bool) *VideoAppReportRequestPayload {
	s.IsVip = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetLoginNick(v string) *VideoAppReportRequestPayload {
	s.LoginNick = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetOriginUuid(v string) *VideoAppReportRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetPhone(v string) *VideoAppReportRequestPayload {
	s.Phone = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetPkgName(v string) *VideoAppReportRequestPayload {
	s.PkgName = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetStartTime(v int64) *VideoAppReportRequestPayload {
	s.StartTime = &v
	return s
}

func (s *VideoAppReportRequestPayload) Validate() error {
	return dara.Validate(s)
}

type VideoAppReportRequestUserInfo struct {
	// 编码类型对应的值
	//
	// - 编码类型是SKILL_ID时，其值为应用的Skill ID。
	//
	// - 编码类型是PACKAGE_NAME时，其值为对应客户端App的packageName。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// 编码类型，获取天猫精灵的用户标识的途径有多种，不同途径对应不同的编码类型。
	//
	// - PACKAGE_NAME：apk包名，Android应用客户链路的编码类型。
	//
	// - SKILL_ID：技能ID，云端链路的编码类型。
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// 用户标识，设置为userOpenId或userUnionId。
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 用户Id的类型
	//
	// - OPEN_ID：默认的用户ID标识。
	//
	// - UNION_ID: 组织维度的用户ID标识，在天猫精灵技能应用开放平台申请过组织后才会有该值。
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// 组织ID，如果IdType为UNION_ID时必填。
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s VideoAppReportRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s VideoAppReportRequestUserInfo) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *VideoAppReportRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *VideoAppReportRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *VideoAppReportRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *VideoAppReportRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *VideoAppReportRequestUserInfo) SetEncodeKey(v string) *VideoAppReportRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetEncodeType(v string) *VideoAppReportRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetId(v string) *VideoAppReportRequestUserInfo {
	s.Id = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetIdType(v string) *VideoAppReportRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetOrganizationId(v string) *VideoAppReportRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
