// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *GetLoginTokenResponseBody
	GetEmail() *string
	SetEndUserId(v string) *GetLoginTokenResponseBody
	GetEndUserId() *string
	SetIndustry(v string) *GetLoginTokenResponseBody
	GetIndustry() *string
	SetKeepAliveToken(v string) *GetLoginTokenResponseBody
	GetKeepAliveToken() *string
	SetLabel(v string) *GetLoginTokenResponseBody
	GetLabel() *string
	SetLoginToken(v string) *GetLoginTokenResponseBody
	GetLoginToken() *string
	SetNextStage(v string) *GetLoginTokenResponseBody
	GetNextStage() *string
	SetNickName(v string) *GetLoginTokenResponseBody
	GetNickName() *string
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
	SetTenantId(v int64) *GetLoginTokenResponseBody
	GetTenantId() *int64
	SetWindowDisplayMode(v string) *GetLoginTokenResponseBody
	GetWindowDisplayMode() *string
}

type GetLoginTokenResponseBody struct {
	// The email address of the user. The system returns the email address in the return value of the LoginToken parameter after the user logs on to the client.
	//
	// 	- For a convenience user, the return value is the email address specified when the administrator creates the convenience user.
	//
	// 	- For an AD user, the return value is in the following format: `Username@Name of the AD domain`.
	//
	// example:
	//
	// alice
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The account of the convenience user or the AD user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// edu
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// The token used to keep the user logged on. After the user logs on to the client and select the Keep Logon option, `KeepAliveToken` is returned when you call the operation. If the user does not select the Keep Logon option, null is returned.
	//
	// example:
	//
	// 006YwvYMsesWWsDBZnVB+Wq9AvJDVIqOY3YCktvtb7+KxMb3ClnNlV8+l/knhZYrXUmeP06IzkjF+IgcZ3vZKOyMprDyFHjCy1r27FRE/U7+geWCl8iQ+yF8GaCRHfJEkC2+ROs93HkT4tfHxyY1J8W7O7ZQGUC/cdCvm+cCP6FIy73IUuPuVR6PcKYXIpEZPW
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// The attribute of the convenience user. For an AD user, null is returned.
	//
	// example:
	//
	// test:sample
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The next stage that is expected to enter. For example, an administrator enables MFA in the EDS console. When an end user enters the password, that is, the end user completes the `ADPassword` stage, this parameter returns `MFAVerify`. This indicates that MFA is required.
	//
	// >  For more information about the authentication stages, see the `CurrentStage` parameter.
	//
	// example:
	//
	// MFAVerify
	NextStage *string `json:"NextStage,omitempty" xml:"NextStage,omitempty"`
	NickName  *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// > This is a parameter only for internal use.
	PasswordStrategy *GetLoginTokenResponseBodyPasswordStrategy `json:"PasswordStrategy,omitempty" xml:"PasswordStrategy,omitempty" type:"Struct"`
	// Enter the mobile number of the convenience user. For an AD user, null is returned.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// > This is a parameter only for internal use.
	Props map[string]*string `json:"Props,omitempty" xml:"Props,omitempty"`
	// The QR code that is generated when the virtual MFA device is bound. The value is encoded in Base64. This parameter can be empty. This parameter is required only when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWY******
	QrCodePng *string `json:"QrCodePng,omitempty" xml:"QrCodePng,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// null
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Risk identification information regarding the signin process.
	RiskVerifyInfo *GetLoginTokenResponseBodyRiskVerifyInfo `json:"RiskVerifyInfo,omitempty" xml:"RiskVerifyInfo,omitempty" type:"Struct"`
	// The key that is generated when you bind the virtual MFA device. This parameter is required when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// 5OCLLKKOJU5HPBX66H3QCTWYI7MH****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The ID of the session. The ID is returned the first time you call the `GetLoginToken` operation in the session. If MFA is required, you must specify this parameter in subsequent stages.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	//
	// example:
	//
	// d6ec166d-ab93-4286-bf7f-a18bb929****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the Alibaba Cloud account. The ID is used for hardware client authentication.
	//
	// example:
	//
	// 166353906220****
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// mode
	WindowDisplayMode *string `json:"WindowDisplayMode,omitempty" xml:"WindowDisplayMode,omitempty"`
}

func (s GetLoginTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) GetEmail() *string {
	return s.Email
}

func (s *GetLoginTokenResponseBody) GetEndUserId() *string {
	return s.EndUserId
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

func (s *GetLoginTokenResponseBody) GetNextStage() *string {
	return s.NextStage
}

func (s *GetLoginTokenResponseBody) GetNickName() *string {
	return s.NickName
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

func (s *GetLoginTokenResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetLoginTokenResponseBody) GetWindowDisplayMode() *string {
	return s.WindowDisplayMode
}

func (s *GetLoginTokenResponseBody) SetEmail(v string) *GetLoginTokenResponseBody {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEndUserId(v string) *GetLoginTokenResponseBody {
	s.EndUserId = &v
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

func (s *GetLoginTokenResponseBody) SetNextStage(v string) *GetLoginTokenResponseBody {
	s.NextStage = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetNickName(v string) *GetLoginTokenResponseBody {
	s.NickName = &v
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

func (s *GetLoginTokenResponseBody) SetTenantId(v int64) *GetLoginTokenResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWindowDisplayMode(v string) *GetLoginTokenResponseBody {
	s.WindowDisplayMode = &v
	return s
}

func (s *GetLoginTokenResponseBody) Validate() error {
	if s.PasswordStrategy != nil {
		if err := s.PasswordStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.RiskVerifyInfo != nil {
		if err := s.RiskVerifyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginTokenResponseBodyPasswordStrategy struct {
	// > This is a parameter only for internal use.
	TenantAlternativeChars []*string `json:"TenantAlternativeChars,omitempty" xml:"TenantAlternativeChars,omitempty" type:"Repeated"`
	// > This is a parameter only for internal use.
	//
	// example:
	//
	// null
	TenantPasswordLength *string `json:"TenantPasswordLength,omitempty" xml:"TenantPasswordLength,omitempty"`
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

func (s *GetLoginTokenResponseBodyPasswordStrategy) GetTenantPasswordLength() *string {
	return s.TenantPasswordLength
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantAlternativeChars(v []*string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantAlternativeChars = v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) SetTenantPasswordLength(v string) *GetLoginTokenResponseBodyPasswordStrategy {
	s.TenantPasswordLength = &v
	return s
}

func (s *GetLoginTokenResponseBodyPasswordStrategy) Validate() error {
	return dara.Validate(s)
}

type GetLoginTokenResponseBodyRiskVerifyInfo struct {
	// The email used for authentication.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The duration of the lock.
	//
	// example:
	//
	// 1713749778
	LastLockDuration *int64 `json:"LastLockDuration,omitempty" xml:"LastLockDuration,omitempty"`
	// Whether the account is locked or not.
	//
	// example:
	//
	// true
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The mobile number used for authentication.
	//
	// example:
	//
	// 1388888****
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

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) GetLocked() *string {
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

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLocked(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
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
