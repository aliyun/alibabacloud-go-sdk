// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPolicyGroupsResponseBody
	GetNextToken() *string
	SetPolicyGroupModel(v []*ListPolicyGroupsResponseBodyPolicyGroupModel) *ListPolicyGroupsResponseBody
	GetPolicyGroupModel() []*ListPolicyGroupsResponseBodyPolicyGroupModel
	SetRequestId(v string) *ListPolicyGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPolicyGroupsResponseBody
	GetTotalCount() *int32
}

type ListPolicyGroupsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policies.
	PolicyGroupModel []*ListPolicyGroupsResponseBodyPolicyGroupModel `json:"PolicyGroupModel,omitempty" xml:"PolicyGroupModel,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicyGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicyGroupsResponseBody) GetPolicyGroupModel() []*ListPolicyGroupsResponseBodyPolicyGroupModel {
	return s.PolicyGroupModel
}

func (s *ListPolicyGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPolicyGroupsResponseBody) SetNextToken(v string) *ListPolicyGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetPolicyGroupModel(v []*ListPolicyGroupsResponseBodyPolicyGroupModel) *ListPolicyGroupsResponseBody {
	s.PolicyGroupModel = v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetRequestId(v string) *ListPolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyGroupsResponseBody) SetTotalCount(v int32) *ListPolicyGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPolicyGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModel struct {
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
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: read and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2024-06-04 10:28:54
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The file transfer policy of the HTML5 client.
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
	// download
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write denied.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Identifies whether the resolution is locked.
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
	NetRedirectPolicy *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-9q6o8qpiy8opkj****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// Default Policy
	PolicyGroupName        *string                                                             `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	PolicyRelatedResources *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources `json:"PolicyRelatedResources,omitempty" xml:"PolicyRelatedResources,omitempty" type:"Struct"`
	// The height of the resolution.
	//
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// The width of the resolution.
	//
	// example:
	//
	// 1920
	SessionResolutionWidth *int32                                                 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	Watermark              *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetClipboard() *string {
	return s.Clipboard
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetLockResolution() *string {
	return s.LockResolution
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetNetRedirectPolicy() *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	return s.NetRedirectPolicy
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetPolicyRelatedResources() *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources {
	return s.PolicyRelatedResources
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetSessionResolutionHeight() *int32 {
	return s.SessionResolutionHeight
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetSessionResolutionWidth() *int32 {
	return s.SessionResolutionWidth
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetWatermark() *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	return s.Watermark
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetCameraRedirect(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.CameraRedirect = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetClipboard(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Clipboard = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetGmtCreate(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.GmtCreate = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetHtml5FileTransfer(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Html5FileTransfer = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetLocalDrive(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.LocalDrive = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetLockResolution(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.LockResolution = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetNetRedirectPolicy(v *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.NetRedirectPolicy = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyGroupId(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyGroupId = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyGroupName(v string) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyGroupName = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetPolicyRelatedResources(v *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.PolicyRelatedResources = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetSessionResolutionHeight(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.SessionResolutionHeight = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetSessionResolutionWidth(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.SessionResolutionWidth = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetWatermark(v *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.Watermark = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy struct {
	// Indicates whether a custom proxy is manually configured.
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
	// Indicates whether the network redirection feature is enabled. When this feature is enabled, network traffic is automatically redirected to the on-premises network by default.
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
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	// The proxy rules.
	Rules []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetCustomProxy() *string {
	return s.CustomProxy
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetHostAddr() *string {
	return s.HostAddr
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetNetRedirect() *string {
	return s.NetRedirect
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetPort() *string {
	return s.Port
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetProxyPassword() *string {
	return s.ProxyPassword
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetProxyType() *string {
	return s.ProxyType
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetProxyUserName() *string {
	return s.ProxyUserName
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) GetRules() []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules {
	return s.Rules
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetCustomProxy(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.CustomProxy = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetHostAddr(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.HostAddr = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetNetRedirect(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.NetRedirect = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetPort(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.Port = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyPassword(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyPassword = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyType(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyType = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetProxyUserName(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.ProxyUserName = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) SetRules(v []*ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy {
	s.Rules = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules struct {
	// The type of the rule.
	//
	// Valid values:
	//
	// 	- prc: an application package name.
	//
	// 	- domain: a domain name.
	//
	// example:
	//
	// domain
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The name of the application package or domain name.
	//
	// example:
	//
	// *.example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) GetRuleType() *string {
	return s.RuleType
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) GetTarget() *string {
	return s.Target
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) SetRuleType(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules {
	s.RuleType = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) SetTarget(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules {
	s.Target = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources struct {
	AndroidInstanceGroupIds []*string `json:"AndroidInstanceGroupIds,omitempty" xml:"AndroidInstanceGroupIds,omitempty" type:"Repeated"`
	CloudPhoneMatrixIds     []*string `json:"CloudPhoneMatrixIds,omitempty" xml:"CloudPhoneMatrixIds,omitempty" type:"Repeated"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) GetAndroidInstanceGroupIds() []*string {
	return s.AndroidInstanceGroupIds
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) GetCloudPhoneMatrixIds() []*string {
	return s.CloudPhoneMatrixIds
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) SetAndroidInstanceGroupIds(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources {
	s.AndroidInstanceGroupIds = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) SetCloudPhoneMatrixIds(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources {
	s.CloudPhoneMatrixIds = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelPolicyRelatedResources) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModelWatermark struct {
	WatermarkColor             *int32    `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	WatermarkCustomText        *string   `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkFontSize          *int32    `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	WatermarkSwitch            *string   `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	WatermarkTransparencyValue *int32    `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	WatermarkTypes             []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkColor() *int32 {
	return s.WatermarkColor
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkCustomText() *string {
	return s.WatermarkCustomText
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkFontSize() *int32 {
	return s.WatermarkFontSize
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkSwitch() *string {
	return s.WatermarkSwitch
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkTransparencyValue() *int32 {
	return s.WatermarkTransparencyValue
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkColor(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkColor = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkCustomText(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkCustomText = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkFontSize(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkFontSize = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkSwitch(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkSwitch = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkTransparencyValue(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkTransparencyValue = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) SetWatermarkTypes(v []*string) *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark {
	s.WatermarkTypes = v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark) Validate() error {
	return dara.Validate(s)
}
