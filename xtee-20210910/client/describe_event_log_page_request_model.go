// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventLogPageRequest
	GetLang() *string
	SetAccountIdPRP(v string) *DescribeEventLogPageRequest
	GetAccountIdPRP() *string
	SetBeginTime(v int64) *DescribeEventLogPageRequest
	GetBeginTime() *int64
	SetCondition1AL(v string) *DescribeEventLogPageRequest
	GetCondition1AL() *string
	SetCondition2AL(v string) *DescribeEventLogPageRequest
	GetCondition2AL() *string
	SetCondition3AL(v string) *DescribeEventLogPageRequest
	GetCondition3AL() *string
	SetCurrentPage(v int32) *DescribeEventLogPageRequest
	GetCurrentPage() *int32
	SetDeviceTypeLRP(v string) *DescribeEventLogPageRequest
	GetDeviceTypeLRP() *string
	SetEmailPRP(v string) *DescribeEventLogPageRequest
	GetEmailPRP() *string
	SetEndTime(v int64) *DescribeEventLogPageRequest
	GetEndTime() *int64
	SetFailReasonLRP(v string) *DescribeEventLogPageRequest
	GetFailReasonLRP() *string
	SetIpPRP(v string) *DescribeEventLogPageRequest
	GetIpPRP() *string
	SetLoginResultARP(v string) *DescribeEventLogPageRequest
	GetLoginResultARP() *string
	SetLoginTypeLRP(v string) *DescribeEventLogPageRequest
	GetLoginTypeLRP() *string
	SetMacPRP(v string) *DescribeEventLogPageRequest
	GetMacPRP() *string
	SetMobilePRP(v string) *DescribeEventLogPageRequest
	GetMobilePRP() *string
	SetNickNamePRP(v string) *DescribeEventLogPageRequest
	GetNickNamePRP() *string
	SetOperateSourceLRP(v string) *DescribeEventLogPageRequest
	GetOperateSourceLRP() *string
	SetPageSize(v int32) *DescribeEventLogPageRequest
	GetPageSize() *int32
	SetReferPRP(v string) *DescribeEventLogPageRequest
	GetReferPRP() *string
	SetRegId(v string) *DescribeEventLogPageRequest
	GetRegId() *string
	SetRegisterIpPRP(v string) *DescribeEventLogPageRequest
	GetRegisterIpPRP() *string
	SetReqIdPBS(v string) *DescribeEventLogPageRequest
	GetReqIdPBS() *string
	SetScoreEBS(v int32) *DescribeEventLogPageRequest
	GetScoreEBS() *int32
	SetScoreSBS(v int32) *DescribeEventLogPageRequest
	GetScoreSBS() *int32
	SetServiceABS(v string) *DescribeEventLogPageRequest
	GetServiceABS() *string
	SetTagsLBS(v string) *DescribeEventLogPageRequest
	GetTagsLBS() *string
	SetUmidPDI(v string) *DescribeEventLogPageRequest
	GetUmidPDI() *string
	SetUserAgentPRP(v string) *DescribeEventLogPageRequest
	GetUserAgentPRP() *string
	SetUserNameTypeLRP(v string) *DescribeEventLogPageRequest
	GetUserNameTypeLRP() *string
}

type DescribeEventLogPageRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The account ID (request_param.accountId). The value can be up to 50 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 180650758xxxxxxx
	AccountIdPRP *string `json:"accountIdPRP,omitempty" xml:"accountIdPRP,omitempty"`
	// The start timestamp of the log. Unit: milliseconds.
	//
	// example:
	//
	// 1737101348000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The first full-text match condition. The value can be up to 30 characters in length.
	//
	// example:
	//
	// rm0102
	Condition1AL *string `json:"condition1AL,omitempty" xml:"condition1AL,omitempty"`
	// The second full-text match condition. The value can be up to 30 characters in length.
	//
	// example:
	//
	// rm0102
	Condition2AL *string `json:"condition2AL,omitempty" xml:"condition2AL,omitempty"`
	// The third full-text match condition. The value can be up to 30 characters in length.
	//
	// example:
	//
	// rm0102
	Condition3AL *string `json:"condition3AL,omitempty" xml:"condition3AL,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The device type (request_param.deviceType). Example values: 1: PC. 2: MOBILE.
	//
	// example:
	//
	// PC
	DeviceTypeLRP *string `json:"deviceTypeLRP,omitempty" xml:"deviceTypeLRP,omitempty"`
	// The email address (request_param.email). The value can be up to 100 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// xxxx@123.com
	EmailPRP *string `json:"emailPRP,omitempty" xml:"emailPRP,omitempty"`
	// The end time. Unit: milliseconds.
	//
	// example:
	//
	// 1746669075000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The logon failure reason (-request_param.failReason).
	//
	// example:
	//
	// wrongPassword
	FailReasonLRP *string `json:"failReasonLRP,omitempty" xml:"failReasonLRP,omitempty"`
	// The IP address (request_param.ip). The value can be up to 20 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 168.168.168.168
	IpPRP *string `json:"ipPRP,omitempty" xml:"ipPRP,omitempty"`
	// The logon success flag (request_param.loginResult).
	//
	// example:
	//
	// SUCCESS
	LoginResultARP *string `json:"loginResultARP,omitempty" xml:"loginResultARP,omitempty"`
	// The logon authentication method (-request_param.loginType).
	//
	// example:
	//
	// PASSWORD
	LoginTypeLRP *string `json:"loginTypeLRP,omitempty" xml:"loginTypeLRP,omitempty"`
	// The device MAC address (-request_param.mac). The value can be up to 30 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 00-1C-F0-1D-A7-81
	MacPRP *string `json:"macPRP,omitempty" xml:"macPRP,omitempty"`
	// The phone number (supports MD5: request_param.mobile/request_param.mobileMd5). The value can be up to 30 characters in length and supports the "\\*" and "?" wildcards. The search is performed based on the mobile and mobileMd5 fields.
	//
	// example:
	//
	// 17600000000
	MobilePRP *string `json:"mobilePRP,omitempty" xml:"mobilePRP,omitempty"`
	// The account nickname (request_param.nickName). The value can be up to 50 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 测试xx
	NickNamePRP *string `json:"nickNamePRP,omitempty" xml:"nickNamePRP,omitempty"`
	// The operation source (request_param.operateSource). Example values: 1: PC. 2: H5. 3: App.
	//
	// example:
	//
	// PC
	OperateSourceLRP *string `json:"operateSourceLRP,omitempty" xml:"operateSourceLRP,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The referer (-request_param.refer). The value can be up to 50 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// refer
	ReferPRP *string `json:"referPRP,omitempty" xml:"referPRP,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The account registration IP address (request_param.registerIp). The value can be up to 20 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 168.168.168.168
	RegisterIpPRP *string `json:"registerIpPRP,omitempty" xml:"registerIpPRP,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD6B08EC-1B44-5378-8838-C76A36415C55
	ReqIdPBS *string `json:"reqIdPBS,omitempty" xml:"reqIdPBS,omitempty"`
	// The end value of the score range (score). Only non-negative integers are allowed. The end value must be greater than the start value. Both boundaries are inclusive.
	//
	// example:
	//
	// 2
	ScoreEBS *int32 `json:"scoreEBS,omitempty" xml:"scoreEBS,omitempty"`
	// The start value of the score range (score). Only non-negative integers are allowed. The end value must be greater than the start value. Both boundaries are inclusive.
	//
	// example:
	//
	// 1
	ScoreSBS *int32 `json:"scoreSBS,omitempty" xml:"scoreSBS,omitempty"`
	// The event name (instance_id).
	//
	// example:
	//
	// de_afghcf6411
	ServiceABS *string `json:"serviceABS,omitempty" xml:"serviceABS,omitempty"`
	// The risk label (tags). The data is obtained from DescribeTagsList.
	//
	// example:
	//
	// rg0001
	TagsLBS *string `json:"tagsLBS,omitempty" xml:"tagsLBS,omitempty"`
	// The device ID (device_info.umid).
	//
	// example:
	//
	// 设备ID
	UmidPDI *string `json:"umidPDI,omitempty" xml:"umidPDI,omitempty"`
	// The user agent (-request_param.userAgent). The value can be up to 50 characters in length and supports the "\\*" and "?" wildcards.
	//
	// example:
	//
	// 00-1C-F0-1D-A7-81
	UserAgentPRP *string `json:"userAgentPRP,omitempty" xml:"userAgentPRP,omitempty"`
	// The account name type for the logon scenario (-request_param.userNameType).
	//
	// example:
	//
	// type
	UserNameTypeLRP *string `json:"userNameTypeLRP,omitempty" xml:"userNameTypeLRP,omitempty"`
}

func (s DescribeEventLogPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventLogPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventLogPageRequest) GetAccountIdPRP() *string {
	return s.AccountIdPRP
}

func (s *DescribeEventLogPageRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeEventLogPageRequest) GetCondition1AL() *string {
	return s.Condition1AL
}

func (s *DescribeEventLogPageRequest) GetCondition2AL() *string {
	return s.Condition2AL
}

func (s *DescribeEventLogPageRequest) GetCondition3AL() *string {
	return s.Condition3AL
}

func (s *DescribeEventLogPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventLogPageRequest) GetDeviceTypeLRP() *string {
	return s.DeviceTypeLRP
}

func (s *DescribeEventLogPageRequest) GetEmailPRP() *string {
	return s.EmailPRP
}

func (s *DescribeEventLogPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeEventLogPageRequest) GetFailReasonLRP() *string {
	return s.FailReasonLRP
}

func (s *DescribeEventLogPageRequest) GetIpPRP() *string {
	return s.IpPRP
}

func (s *DescribeEventLogPageRequest) GetLoginResultARP() *string {
	return s.LoginResultARP
}

func (s *DescribeEventLogPageRequest) GetLoginTypeLRP() *string {
	return s.LoginTypeLRP
}

func (s *DescribeEventLogPageRequest) GetMacPRP() *string {
	return s.MacPRP
}

func (s *DescribeEventLogPageRequest) GetMobilePRP() *string {
	return s.MobilePRP
}

func (s *DescribeEventLogPageRequest) GetNickNamePRP() *string {
	return s.NickNamePRP
}

func (s *DescribeEventLogPageRequest) GetOperateSourceLRP() *string {
	return s.OperateSourceLRP
}

func (s *DescribeEventLogPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventLogPageRequest) GetReferPRP() *string {
	return s.ReferPRP
}

func (s *DescribeEventLogPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventLogPageRequest) GetRegisterIpPRP() *string {
	return s.RegisterIpPRP
}

func (s *DescribeEventLogPageRequest) GetReqIdPBS() *string {
	return s.ReqIdPBS
}

func (s *DescribeEventLogPageRequest) GetScoreEBS() *int32 {
	return s.ScoreEBS
}

func (s *DescribeEventLogPageRequest) GetScoreSBS() *int32 {
	return s.ScoreSBS
}

func (s *DescribeEventLogPageRequest) GetServiceABS() *string {
	return s.ServiceABS
}

func (s *DescribeEventLogPageRequest) GetTagsLBS() *string {
	return s.TagsLBS
}

func (s *DescribeEventLogPageRequest) GetUmidPDI() *string {
	return s.UmidPDI
}

func (s *DescribeEventLogPageRequest) GetUserAgentPRP() *string {
	return s.UserAgentPRP
}

func (s *DescribeEventLogPageRequest) GetUserNameTypeLRP() *string {
	return s.UserNameTypeLRP
}

func (s *DescribeEventLogPageRequest) SetLang(v string) *DescribeEventLogPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetAccountIdPRP(v string) *DescribeEventLogPageRequest {
	s.AccountIdPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetBeginTime(v int64) *DescribeEventLogPageRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetCondition1AL(v string) *DescribeEventLogPageRequest {
	s.Condition1AL = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetCondition2AL(v string) *DescribeEventLogPageRequest {
	s.Condition2AL = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetCondition3AL(v string) *DescribeEventLogPageRequest {
	s.Condition3AL = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetCurrentPage(v int32) *DescribeEventLogPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetDeviceTypeLRP(v string) *DescribeEventLogPageRequest {
	s.DeviceTypeLRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetEmailPRP(v string) *DescribeEventLogPageRequest {
	s.EmailPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetEndTime(v int64) *DescribeEventLogPageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetFailReasonLRP(v string) *DescribeEventLogPageRequest {
	s.FailReasonLRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetIpPRP(v string) *DescribeEventLogPageRequest {
	s.IpPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetLoginResultARP(v string) *DescribeEventLogPageRequest {
	s.LoginResultARP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetLoginTypeLRP(v string) *DescribeEventLogPageRequest {
	s.LoginTypeLRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetMacPRP(v string) *DescribeEventLogPageRequest {
	s.MacPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetMobilePRP(v string) *DescribeEventLogPageRequest {
	s.MobilePRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetNickNamePRP(v string) *DescribeEventLogPageRequest {
	s.NickNamePRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetOperateSourceLRP(v string) *DescribeEventLogPageRequest {
	s.OperateSourceLRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetPageSize(v int32) *DescribeEventLogPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetReferPRP(v string) *DescribeEventLogPageRequest {
	s.ReferPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetRegId(v string) *DescribeEventLogPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetRegisterIpPRP(v string) *DescribeEventLogPageRequest {
	s.RegisterIpPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetReqIdPBS(v string) *DescribeEventLogPageRequest {
	s.ReqIdPBS = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetScoreEBS(v int32) *DescribeEventLogPageRequest {
	s.ScoreEBS = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetScoreSBS(v int32) *DescribeEventLogPageRequest {
	s.ScoreSBS = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetServiceABS(v string) *DescribeEventLogPageRequest {
	s.ServiceABS = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetTagsLBS(v string) *DescribeEventLogPageRequest {
	s.TagsLBS = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetUmidPDI(v string) *DescribeEventLogPageRequest {
	s.UmidPDI = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetUserAgentPRP(v string) *DescribeEventLogPageRequest {
	s.UserAgentPRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) SetUserNameTypeLRP(v string) *DescribeEventLogPageRequest {
	s.UserNameTypeLRP = &v
	return s
}

func (s *DescribeEventLogPageRequest) Validate() error {
	return dara.Validate(s)
}
