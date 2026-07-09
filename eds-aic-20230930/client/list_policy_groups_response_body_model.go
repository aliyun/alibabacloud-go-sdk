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
	// The pagination token that indicates the position up to which data has been read in the current call. An empty value indicates that all data has been read.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy information.
	PolicyGroupModel []*ListPolicyGroupsResponseBodyPolicyGroupModel `json:"PolicyGroupModel,omitempty" xml:"PolicyGroupModel,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
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
	if s.PolicyGroupModel != nil {
		for _, item := range s.PolicyGroupModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyGroupsResponseBodyPolicyGroupModel struct {
	AccessPolicies []*ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies `json:"AccessPolicies,omitempty" xml:"AccessPolicies,omitempty" type:"Repeated"`
	// Indicates whether local camera redirection is enabled.
	//
	// example:
	//
	// on
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The clipboard permission.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-06-04 10:28:54
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The file transfer policy for the HTML5 client.
	//
	// example:
	//
	// download
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// The locked resolution.
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection settings.
	NetRedirectPolicy *ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty" type:"Struct"`
	// The policy ID.
	//
	// example:
	//
	// pg-9q6o8qpiy8opkj****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Default policy.
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The resources associated with the policy.
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
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	// The screen watermark settings.
	Watermark *ListPolicyGroupsResponseBodyPolicyGroupModelWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModel) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) GetAccessPolicies() []*ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	return s.AccessPolicies
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

func (s *ListPolicyGroupsResponseBodyPolicyGroupModel) SetAccessPolicies(v []*ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) *ListPolicyGroupsResponseBodyPolicyGroupModel {
	s.AccessPolicies = v
	return s
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
	if s.AccessPolicies != nil {
		for _, item := range s.AccessPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetRedirectPolicy != nil {
		if err := s.NetRedirectPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyRelatedResources != nil {
		if err := s.PolicyRelatedResources.Validate(); err != nil {
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

type ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies struct {
	AccessPolicyRuleId *int64  `json:"AccessPolicyRuleId,omitempty" xml:"AccessPolicyRuleId,omitempty"`
	CidrIp             *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Policy             *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GetAccessPolicyRuleId() *int64 {
	return s.AccessPolicyRuleId
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GetCidrIp() *string {
	return s.CidrIp
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GetPolicy() *string {
	return s.Policy
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) SetAccessPolicyRuleId(v int64) *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	s.AccessPolicyRuleId = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) SetCidrIp(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	s.CidrIp = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) SetDescription(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	s.Description = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) SetPolicy(v string) *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	s.Policy = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) SetPriority(v int32) *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies {
	s.Priority = &v
	return s
}

func (s *ListPolicyGroupsResponseBodyPolicyGroupModelAccessPolicies) Validate() error {
	return dara.Validate(s)
}

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicy struct {
	// Indicates whether a transparent proxy is manually configured.
	//
	// example:
	//
	// off
	CustomProxy *string `json:"CustomProxy,omitempty" xml:"CustomProxy,omitempty"`
	// The proxy IP address of the transparent proxy. The value must be in IPv4 format.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddr *string `json:"HostAddr,omitempty" xml:"HostAddr,omitempty"`
	// Indicates whether network redirection is enabled. After this feature is enabled, traffic is redirected to the client-side network by default.
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
	// The proxy password. The value must be 1 to 256 characters in length and cannot contain Chinese characters or whitespace characters.
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
	// The proxy username. The value must be 1 to 256 characters in length and cannot contain Chinese characters or whitespace characters.
	//
	// example:
	//
	// username
	ProxyUserName *string `json:"ProxyUserName,omitempty" xml:"ProxyUserName,omitempty"`
	// The list of proxy rules.
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

type ListPolicyGroupsResponseBodyPolicyGroupModelNetRedirectPolicyRules struct {
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
	// The list of instance group IDs.
	AndroidInstanceGroupIds []*string `json:"AndroidInstanceGroupIds,omitempty" xml:"AndroidInstanceGroupIds,omitempty" type:"Repeated"`
	// The list of matrix IDs.
	CloudPhoneMatrixIds []*string `json:"CloudPhoneMatrixIds,omitempty" xml:"CloudPhoneMatrixIds,omitempty" type:"Repeated"`
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
	// The watermark font color. Valid values: 0 to 16777215.
	//
	// example:
	//
	// 0
	WatermarkColor *int32 `json:"WatermarkColor,omitempty" xml:"WatermarkColor,omitempty"`
	// The custom watermark content. The value can be up to 10 characters in length and does not support emoji characters.
	//
	// example:
	//
	// custom text
	WatermarkCustomText *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	// The watermark font size. Valid values: 10 to 20.
	//
	// example:
	//
	// 12
	WatermarkFontSize *int32 `json:"WatermarkFontSize,omitempty" xml:"WatermarkFontSize,omitempty"`
	// The screen watermark switch.
	//
	// example:
	//
	// off
	WatermarkSwitch *string `json:"WatermarkSwitch,omitempty" xml:"WatermarkSwitch,omitempty"`
	// The watermark opacity. A larger value indicates lower transparency. Valid values: 10 to 100.
	//
	// example:
	//
	// 25
	WatermarkTransparencyValue *int32 `json:"WatermarkTransparencyValue,omitempty" xml:"WatermarkTransparencyValue,omitempty"`
	// The screen watermark content.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
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
