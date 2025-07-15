// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraRedirect(v string) *ModifyPolicyGroupRequest
	GetCameraRedirect() *string
	SetClipboard(v string) *ModifyPolicyGroupRequest
	GetClipboard() *string
	SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest
	GetHtml5FileTransfer() *string
	SetLocalDrive(v string) *ModifyPolicyGroupRequest
	GetLocalDrive() *string
	SetLockResolution(v string) *ModifyPolicyGroupRequest
	GetLockResolution() *string
	SetNetRedirectPolicy(v *ModifyPolicyGroupRequestNetRedirectPolicy) *ModifyPolicyGroupRequest
	GetNetRedirectPolicy() *ModifyPolicyGroupRequestNetRedirectPolicy
	SetPolicyGroupId(v string) *ModifyPolicyGroupRequest
	GetPolicyGroupId() *string
	SetPolicyGroupName(v string) *ModifyPolicyGroupRequest
	GetPolicyGroupName() *string
	SetResolutionHeight(v int32) *ModifyPolicyGroupRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *ModifyPolicyGroupRequest
	GetResolutionWidth() *int32
	SetWatermark(v *ModifyPolicyGroupRequestWatermark) *ModifyPolicyGroupRequest
	GetWatermark() *ModifyPolicyGroupRequestWatermark
}

type ModifyPolicyGroupRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicy *ModifyPolicyGroupRequestNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-4bi18ebi9tfjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicyGroup
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32                             `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	Watermark       *ModifyPolicyGroupRequestWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *ModifyPolicyGroupRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *ModifyPolicyGroupRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ModifyPolicyGroupRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *ModifyPolicyGroupRequest) GetLockResolution() *string {
	return s.LockResolution
}

func (s *ModifyPolicyGroupRequest) GetNetRedirectPolicy() *ModifyPolicyGroupRequestNetRedirectPolicy {
	return s.NetRedirectPolicy
}

func (s *ModifyPolicyGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyPolicyGroupRequest) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *ModifyPolicyGroupRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *ModifyPolicyGroupRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *ModifyPolicyGroupRequest) GetWatermark() *ModifyPolicyGroupRequestWatermark {
	return s.Watermark
}

func (s *ModifyPolicyGroupRequest) SetCameraRedirect(v string) *ModifyPolicyGroupRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLockResolution(v string) *ModifyPolicyGroupRequest {
	s.LockResolution = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetNetRedirectPolicy(v *ModifyPolicyGroupRequestNetRedirectPolicy) *ModifyPolicyGroupRequest {
	s.NetRedirectPolicy = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupName(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetResolutionHeight(v int32) *ModifyPolicyGroupRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetResolutionWidth(v int32) *ModifyPolicyGroupRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v *ModifyPolicyGroupRequestWatermark) *ModifyPolicyGroupRequest {
	s.Watermark = v
	return s
}

func (s *ModifyPolicyGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestNetRedirectPolicy struct {
	// Specifies whether to manually configure a custom proxy.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The IPv4 address of the custom proxy.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Specifies whether to enable network redirection.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	NetRedirect *string `json:"NetRedirect,omitempty" xml:"NetRedirect,omitempty"`
	// The port of the custom proxy. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1145
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The password of the proxy. The password must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// password
	ProxyPassword *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	// The type of the proxy protocol.
	//
	// Valid values:
	//
	// 	- socks5.
	//
	// example:
	//
	// socks5
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The username of the proxy. The name must be 1 to 256 in length and cannot contain Chinese character or space characters.
	//
	// example:
	//
	// username
	ProxyUserName *string                                           `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	Rules         []*ModifyPolicyGroupRequestNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyPolicyGroupRequestNetRedirectPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetCustomProxy() *string {
	return s.CustomProxy
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetHostAddr() *string {
	return s.HostAddr
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetPort() *string {
	return s.Port
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetProxyPassword() *string {
	return s.ProxyPassword
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetProxyUserName() *string {
	return s.ProxyUserName
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) GetRules() []*ModifyPolicyGroupRequestNetRedirectPolicyRules {
	return s.Rules
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetCustomProxy(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetHostAddr(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetNetRedirect(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetPort(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyPassword(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyType(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetProxyUserName(v string) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) SetRules(v []*ModifyPolicyGroupRequestNetRedirectPolicyRules) *ModifyPolicyGroupRequestNetRedirectPolicy {
	s.Rules = v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestNetRedirectPolicyRules struct {
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyPolicyGroupRequestNetRedirectPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) GetTarget() *string {
	return s.Target
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) SetRuleType(v string) *ModifyPolicyGroupRequestNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) SetTarget(v string) *ModifyPolicyGroupRequestNetRedirectPolicyRules {
	s.Target = &v
	return s
}

func (s *ModifyPolicyGroupRequestNetRedirectPolicyRules) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyGroupRequestWatermark struct {
	WatermarkColor             *int32    `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	WatermarkCustomText        *string   `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkFontSize          *int32    `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	WatermarkSwitch            *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTransparencyValue *int32    `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	WatermarkTypes             []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ModifyPolicyGroupRequestWatermark) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupRequestWatermark) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *ModifyPolicyGroupRequestWatermark) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkColor(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkFontSize(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkSwitch(v string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkTransparencyValue(v int32) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) SetWatermarkTypes(v []*string) *ModifyPolicyGroupRequestWatermark {
	s.WatermarkTypes = v
	return s
}

func (s *ModifyPolicyGroupRequestWatermark) Validate() error {
	return dara.Validate(s)
}
