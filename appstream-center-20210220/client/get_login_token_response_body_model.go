// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *GetLoginTokenResponseBody
	GetAccessType() *string
	SetAccountType(v string) *GetLoginTokenResponseBody
	GetAccountType() *string
	SetAdDomain(v string) *GetLoginTokenResponseBody
	GetAdDomain() *string
	SetEmail(v string) *GetLoginTokenResponseBody
	GetEmail() *string
	SetEndUserId(v string) *GetLoginTokenResponseBody
	GetEndUserId() *string
	SetIdpId(v string) *GetLoginTokenResponseBody
	GetIdpId() *string
	SetIndustry(v string) *GetLoginTokenResponseBody
	GetIndustry() *string
	SetKeepAliveToken(v string) *GetLoginTokenResponseBody
	GetKeepAliveToken() *string
	SetLabel(v string) *GetLoginTokenResponseBody
	GetLabel() *string
	SetLoginToken(v string) *GetLoginTokenResponseBody
	GetLoginToken() *string
	SetMfaTypeList(v []*GetLoginTokenResponseBodyMfaTypeList) *GetLoginTokenResponseBody
	GetMfaTypeList() []*GetLoginTokenResponseBodyMfaTypeList
	SetNextStage(v string) *GetLoginTokenResponseBody
	GetNextStage() *string
	SetNickName(v string) *GetLoginTokenResponseBody
	GetNickName() *string
	SetOfficeSites(v []*string) *GetLoginTokenResponseBody
	GetOfficeSites() []*string
	SetPasswordStrategy(v *GetLoginTokenResponseBodyPasswordStrategy) *GetLoginTokenResponseBody
	GetPasswordStrategy() *GetLoginTokenResponseBodyPasswordStrategy
	SetPhone(v string) *GetLoginTokenResponseBody
	GetPhone() *string
	SetProps(v map[string]*string) *GetLoginTokenResponseBody
	GetProps() map[string]*string
	SetQrCodePng(v string) *GetLoginTokenResponseBody
	GetQrCodePng() *string
	SetReason(v string) *GetLoginTokenResponseBody
	GetReason() *string
	SetRequestId(v string) *GetLoginTokenResponseBody
	GetRequestId() *string
	SetRiskVerifyInfo(v *GetLoginTokenResponseBodyRiskVerifyInfo) *GetLoginTokenResponseBody
	GetRiskVerifyInfo() *GetLoginTokenResponseBodyRiskVerifyInfo
	SetSecret(v string) *GetLoginTokenResponseBody
	GetSecret() *string
	SetSessionId(v string) *GetLoginTokenResponseBody
	GetSessionId() *string
	SetTenantAlias(v string) *GetLoginTokenResponseBody
	GetTenantAlias() *string
	SetTenantId(v int64) *GetLoginTokenResponseBody
	GetTenantId() *int64
	SetTenantInfos(v []*GetLoginTokenResponseBodyTenantInfos) *GetLoginTokenResponseBody
	GetTenantInfos() []*GetLoginTokenResponseBodyTenantInfos
	SetVpcRegionId(v string) *GetLoginTokenResponseBody
	GetVpcRegionId() *string
	SetWindowDisplayMode(v string) *GetLoginTokenResponseBody
	GetWindowDisplayMode() *string
	SetWyId(v string) *GetLoginTokenResponseBody
	GetWyId() *string
}

type GetLoginTokenResponseBody struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// ad
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// easthp***.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// example:
	//
	// alice***@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// idp-7ttvs4ove8bo5***
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// edu
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// 006YwvYMsesWWsDBZnVB+Wq9AvJDVIqOY3YCktvtb7+KxMb3ClnNlV8+l/knhZYrXUmeP06IzkjF+IgcZ3vZKOyMprDyFHjCy1r27FRE/U7+geWCl8iQ+yF8GaCRHfJEkC2+ROs93HkT4tfHxyY1J8W7O7ZQGUC/cdCvm+cCP6FIy73IUuPuVR6PcKYXIp***
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// example:
	//
	// test:wuying
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d****
	LoginToken  *string                                 `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	MfaTypeList []*GetLoginTokenResponseBodyMfaTypeList `json:"MfaTypeList,omitempty" xml:"MfaTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// MFABind
	NextStage        *string                                    `json:"NextStage,omitempty" xml:"NextStage,omitempty"`
	NickName         *string                                    `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OfficeSites      []*string                                  `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	PasswordStrategy *GetLoginTokenResponseBodyPasswordStrategy `json:"PasswordStrategy,omitempty" xml:"PasswordStrategy,omitempty" type:"Struct"`
	// example:
	//
	// 1826717****
	Phone *string            `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Props map[string]*string `json:"Props,omitempty" xml:"Props,omitempty"`
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWY******
	QrCodePng *string `json:"QrCodePng,omitempty" xml:"QrCodePng,omitempty"`
	// example:
	//
	// PasswordExpired
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// EEA72491-B731-53D6-83ED-209769D6****
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskVerifyInfo *GetLoginTokenResponseBodyRiskVerifyInfo `json:"RiskVerifyInfo,omitempty" xml:"RiskVerifyInfo,omitempty" type:"Struct"`
	// example:
	//
	// 4JZNSDHDM3T6AZ4G2O5OWXBLLE4P****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// example:
	//
	// cc15c91c-821b-4edd-9af2-6df66cc****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// Rq201***
	TenantAlias *string `json:"TenantAlias,omitempty" xml:"TenantAlias,omitempty"`
	// example:
	//
	// 13747924304****
	TenantId    *int64                                  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantInfos []*GetLoginTokenResponseBodyTenantInfos `json:"TenantInfos,omitempty" xml:"TenantInfos,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// example:
	//
	// Windowed
	WindowDisplayMode *string `json:"WindowDisplayMode,omitempty" xml:"WindowDisplayMode,omitempty"`
	// example:
	//
	// 0aba1403b337a***
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s GetLoginTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) GetAccessType() *string {
	return s.AccessType
}

func (s *GetLoginTokenResponseBody) GetAccountType() *string {
	return s.AccountType
}

func (s *GetLoginTokenResponseBody) GetAdDomain() *string {
	return s.AdDomain
}

func (s *GetLoginTokenResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetLoginTokenResponseBody) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetLoginTokenResponseBody) GetIdpId() *string {
	return s.IdpId
}

func (s *GetLoginTokenResponseBody) GetIndustry() *string {
	return s.Industry
}

func (s *GetLoginTokenResponseBody) GetKeepAliveToken() *string {
	return s.KeepAliveToken
}

func (s *GetLoginTokenResponseBody) GetLabel() *string {
	return s.Label
}

func (s *GetLoginTokenResponseBody) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GetLoginTokenResponseBody) GetMfaTypeList() []*GetLoginTokenResponseBodyMfaTypeList {
	return s.MfaTypeList
}

func (s *GetLoginTokenResponseBody) GetNextStage() *string {
	return s.NextStage
}

func (s *GetLoginTokenResponseBody) GetNickName() *string {
	return s.NickName
}

func (s *GetLoginTokenResponseBody) GetOfficeSites() []*string {
	return s.OfficeSites
}

func (s *GetLoginTokenResponseBody) GetPasswordStrategy() *GetLoginTokenResponseBodyPasswordStrategy {
	return s.PasswordStrategy
}

func (s *GetLoginTokenResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetLoginTokenResponseBody) GetProps() map[string]*string {
	return s.Props
}

func (s *GetLoginTokenResponseBody) GetQrCodePng() *string {
	return s.QrCodePng
}

func (s *GetLoginTokenResponseBody) GetReason() *string {
	return s.Reason
}

func (s *GetLoginTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginTokenResponseBody) GetRiskVerifyInfo() *GetLoginTokenResponseBodyRiskVerifyInfo {
	return s.RiskVerifyInfo
}

func (s *GetLoginTokenResponseBody) GetSecret() *string {
	return s.Secret
}

func (s *GetLoginTokenResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLoginTokenResponseBody) GetTenantAlias() *string {
	return s.TenantAlias
}

func (s *GetLoginTokenResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetLoginTokenResponseBody) GetTenantInfos() []*GetLoginTokenResponseBodyTenantInfos {
	return s.TenantInfos
}

func (s *GetLoginTokenResponseBody) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *GetLoginTokenResponseBody) GetWindowDisplayMode() *string {
	return s.WindowDisplayMode
}

func (s *GetLoginTokenResponseBody) GetWyId() *string {
	return s.WyId
}

func (s *GetLoginTokenResponseBody) SetAccessType(v string) *GetLoginTokenResponseBody {
	s.AccessType = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetAccountType(v string) *GetLoginTokenResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetAdDomain(v string) *GetLoginTokenResponseBody {
	s.AdDomain = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEmail(v string) *GetLoginTokenResponseBody {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEndUserId(v string) *GetLoginTokenResponseBody {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIdpId(v string) *GetLoginTokenResponseBody {
	s.IdpId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIndustry(v string) *GetLoginTokenResponseBody {
	s.Industry = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetKeepAliveToken(v string) *GetLoginTokenResponseBody {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLabel(v string) *GetLoginTokenResponseBody {
	s.Label = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLoginToken(v string) *GetLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetMfaTypeList(v []*GetLoginTokenResponseBodyMfaTypeList) *GetLoginTokenResponseBody {
	s.MfaTypeList = v
	return s
}

func (s *GetLoginTokenResponseBody) SetNextStage(v string) *GetLoginTokenResponseBody {
	s.NextStage = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetNickName(v string) *GetLoginTokenResponseBody {
	s.NickName = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetOfficeSites(v []*string) *GetLoginTokenResponseBody {
	s.OfficeSites = v
	return s
}

func (s *GetLoginTokenResponseBody) SetPasswordStrategy(v *GetLoginTokenResponseBodyPasswordStrategy) *GetLoginTokenResponseBody {
	s.PasswordStrategy = v
	return s
}

func (s *GetLoginTokenResponseBody) SetPhone(v string) *GetLoginTokenResponseBody {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetProps(v map[string]*string) *GetLoginTokenResponseBody {
	s.Props = v
	return s
}

func (s *GetLoginTokenResponseBody) SetQrCodePng(v string) *GetLoginTokenResponseBody {
	s.QrCodePng = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetReason(v string) *GetLoginTokenResponseBody {
	s.Reason = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRiskVerifyInfo(v *GetLoginTokenResponseBodyRiskVerifyInfo) *GetLoginTokenResponseBody {
	s.RiskVerifyInfo = v
	return s
}

func (s *GetLoginTokenResponseBody) SetSecret(v string) *GetLoginTokenResponseBody {
	s.Secret = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetSessionId(v string) *GetLoginTokenResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantAlias(v string) *GetLoginTokenResponseBody {
	s.TenantAlias = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantId(v int64) *GetLoginTokenResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantInfos(v []*GetLoginTokenResponseBodyTenantInfos) *GetLoginTokenResponseBody {
	s.TenantInfos = v
	return s
}

func (s *GetLoginTokenResponseBody) SetVpcRegionId(v string) *GetLoginTokenResponseBody {
	s.VpcRegionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWindowDisplayMode(v string) *GetLoginTokenResponseBody {
	s.WindowDisplayMode = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWyId(v string) *GetLoginTokenResponseBody {
	s.WyId = &v
	return s
}

func (s *GetLoginTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLoginTokenResponseBodyMfaTypeList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s GetLoginTokenResponseBodyMfaTypeList) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBodyMfaTypeList) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyMfaTypeList) GetName() *string {
	return s.Name
}

func (s *GetLoginTokenResponseBodyMfaTypeList) GetStage() *string {
	return s.Stage
}

func (s *GetLoginTokenResponseBodyMfaTypeList) SetName(v string) *GetLoginTokenResponseBodyMfaTypeList {
	s.Name = &v
	return s
}

func (s *GetLoginTokenResponseBodyMfaTypeList) SetStage(v string) *GetLoginTokenResponseBodyMfaTypeList {
	s.Stage = &v
	return s
}

func (s *GetLoginTokenResponseBodyMfaTypeList) Validate() error {
	return dara.Validate(s)
}

type GetLoginTokenResponseBodyPasswordStrategy struct {
	TenantAlternativeChars []*string `json:"TenantAlternativeChars,omitempty" xml:"TenantAlternativeChars,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TenantPasswordLength *int32 `json:"TenantPasswordLength,omitempty" xml:"TenantPasswordLength,omitempty"`
}

func (s GetLoginTokenResponseBodyPasswordStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBodyPasswordStrategy) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) GetTenantAlternativeChars() []*string {
	return s.TenantAlternativeChars
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) GetTenantPasswordLength() *int32 {
	return s.TenantPasswordLength
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantAlternativeChars(v []*string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantAlternativeChars = v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantPasswordLength(v int32) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantPasswordLength = &v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) Validate() error {
	return dara.Validate(s)
}

type GetLoginTokenResponseBodyRiskVerifyInfo struct {
	// example:
	//
	// ppas***@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 12000
	LastLockDuration *int64 `json:"LastLockDuration,omitempty" xml:"LastLockDuration,omitempty"`
	// example:
	//
	// true
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// example:
	//
	// 1826717****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) GetEmail() *string {
	return s.Email
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) GetLastLockDuration() *int64 {
	return s.LastLockDuration
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) GetLocked() *bool {
	return s.Locked
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) GetPhone() *string {
	return s.Phone
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetEmail(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLastLockDuration(v int64) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.LastLockDuration = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLocked(v bool) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Locked = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetPhone(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) Validate() error {
	return dara.Validate(s)
}

type GetLoginTokenResponseBodyTenantInfos struct {
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// Up830***
	TenantAlias *string `json:"TenantAlias,omitempty" xml:"TenantAlias,omitempty"`
}

func (s GetLoginTokenResponseBodyTenantInfos) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBodyTenantInfos) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyTenantInfos) GetAccessType() *string {
	return s.AccessType
}

func (s *GetLoginTokenResponseBodyTenantInfos) GetTenantAlias() *string {
	return s.TenantAlias
}

func (s *GetLoginTokenResponseBodyTenantInfos) SetAccessType(v string) *GetLoginTokenResponseBodyTenantInfos {
	s.AccessType = &v
	return s
}

func (s *GetLoginTokenResponseBodyTenantInfos) SetTenantAlias(v string) *GetLoginTokenResponseBodyTenantInfos {
	s.TenantAlias = &v
	return s
}

func (s *GetLoginTokenResponseBodyTenantInfos) Validate() error {
	return dara.Validate(s)
}
