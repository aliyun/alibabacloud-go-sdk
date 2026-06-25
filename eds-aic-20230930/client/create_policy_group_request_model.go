// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraRedirect(v string) *CreatePolicyGroupRequest
	GetCameraRedirect() *string
	SetClipboard(v string) *CreatePolicyGroupRequest
	GetClipboard() *string
	SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest
	GetHtml5FileTransfer() *string
	SetLocalDrive(v string) *CreatePolicyGroupRequest
	GetLocalDrive() *string
	SetLockResolution(v string) *CreatePolicyGroupRequest
	GetLockResolution() *string
	SetNetRedirectPolicy(v *CreatePolicyGroupRequestNetRedirectPolicy) *CreatePolicyGroupRequest
	GetNetRedirectPolicy() *CreatePolicyGroupRequestNetRedirectPolicy
	SetPolicyGroupName(v string) *CreatePolicyGroupRequest
	GetPolicyGroupName() *string
	SetPolicyType(v string) *CreatePolicyGroupRequest
	GetPolicyType() *string
	SetResolutionHeight(v int32) *CreatePolicyGroupRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *CreatePolicyGroupRequest
	GetResolutionWidth() *int32
	SetWatermark(v *CreatePolicyGroupRequestWatermark) *CreatePolicyGroupRequest
	GetWatermark() *CreatePolicyGroupRequestWatermark
}

type CreatePolicyGroupRequest struct {
	// Specifies whether to enable local camera redirection.
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The clipboard permission.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy for the web client.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// Network redirection.
	NetRedirectPolicy *CreatePolicyGroupRequestNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The policy name.
	//
	// example:
	//
	// Default policy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The policy type.
	//
	// example:
	//
	// Instance
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The resolution height, in pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The resolution width, in pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// Screen watermark.
	Watermark *CreatePolicyGroupRequestWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s CreatePolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *CreatePolicyGroupRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *CreatePolicyGroupRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *CreatePolicyGroupRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *CreatePolicyGroupRequest) GetLockResolution() *string {
	return s.LockResolution
}

func (s *CreatePolicyGroupRequest) GetNetRedirectPolicy() *CreatePolicyGroupRequestNetRedirectPolicy {
	return s.NetRedirectPolicy
}

func (s *CreatePolicyGroupRequest) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *CreatePolicyGroupRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolicyGroupRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *CreatePolicyGroupRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *CreatePolicyGroupRequest) GetWatermark() *CreatePolicyGroupRequestWatermark {
	return s.Watermark
}

func (s *CreatePolicyGroupRequest) SetCameraRedirect(v string) *CreatePolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLockResolution(v string) *CreatePolicyGroupRequest {
	s.LockResolution = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetNetRedirectPolicy(v *CreatePolicyGroupRequestNetRedirectPolicy) *CreatePolicyGroupRequest {
	s.NetRedirectPolicy = v
	return s
}

func (s *CreatePolicyGroupRequest) SetPolicyGroupName(v string) *CreatePolicyGroupRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPolicyType(v string) *CreatePolicyGroupRequest {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetResolutionHeight(v int32) *CreatePolicyGroupRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetResolutionWidth(v int32) *CreatePolicyGroupRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v *CreatePolicyGroupRequestWatermark) *CreatePolicyGroupRequest {
	s.Watermark = v
	return s
}

func (s *CreatePolicyGroupRequest) Validate() error {
	if s.NetRedirectPolicy != nil {
		if err := s.NetRedirectPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.Watermark != nil {
		if err := s.Watermark.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyGroupRequestNetRedirectPolicy struct {
	// Specifies whether to manually configure a transparent proxy.
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The IP address of the transparent proxy. The IP address must be in IPv4 format.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The port of the transparent proxy. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1145
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The proxy password. The password must be 1 to 256 characters in length. It cannot contain Chinese characters or whitespace characters.
	//
	// example:
	//
	// password
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The proxy protocol type.
	//
	// example:
	//
	// socks5
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The proxy username. The username must be 1 to 256 characters in length. It cannot contain Chinese characters or whitespace characters.
	//
	// example:
	//
	// username
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	// The list of proxy rules. You can specify up to 100 rules.
	Rules []*CreatePolicyGroupRequestNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreatePolicyGroupRequestNetRedirectPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetCustomProxy() *string {
	return s.CustomProxy
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetHostAddr() *string {
	return s.HostAddr
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetPort() *string {
	return s.Port
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetProxyPassword() *string {
	return s.ProxyPassword
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetProxyUserName() *string {
	return s.ProxyUserName
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) GetRules() []*CreatePolicyGroupRequestNetRedirectPolicyRules {
	return s.Rules
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetCustomProxy(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetHostAddr(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetNetRedirect(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetPort(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyPassword(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyType(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetProxyUserName(v string) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) SetRules(v []*CreatePolicyGroupRequestNetRedirectPolicyRules) *CreatePolicyGroupRequestNetRedirectPolicy {
	s.Rules = v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicy) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePolicyGroupRequestNetRedirectPolicyRules struct {
	// The rule type.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The application package name or domain name.
	//
	// example:
	//
	// *.baidu.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s CreatePolicyGroupRequestNetRedirectPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) GetRuleType() *string {
	return s.RuleType
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) GetTarget() *string {
	return s.Target
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) SetRuleType(v string) *CreatePolicyGroupRequestNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) SetTarget(v string) *CreatePolicyGroupRequestNetRedirectPolicyRules {
	s.Target = &v
	return s
}

func (s *CreatePolicyGroupRequestNetRedirectPolicyRules) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyGroupRequestWatermark struct {
	// The font color of the watermark. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// The custom text for the watermark. The text can be up to 10 characters long and cannot contain emojis.
	//
	// example:
	//
	// custom text
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// The font size of the watermark. Valid values: 10 to 20.
	//
	// example:
	//
	// 12
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// Specifies whether to enable the screen watermark.
	//
	// example:
	//
	// off
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The opacity of the watermark. A larger value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 25
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The screen watermark content.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s CreatePolicyGroupRequestWatermark) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupRequestWatermark) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *CreatePolicyGroupRequestWatermark) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkColor(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkCustomText(v string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkFontSize(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkSwitch(v string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkTransparencyValue(v int32) *CreatePolicyGroupRequestWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) SetWatermarkTypes(v []*string) *CreatePolicyGroupRequestWatermark {
	s.WatermarkTypes = v
	return s
}

func (s *CreatePolicyGroupRequestWatermark) Validate() error {
	return dara.Validate(s)
}
