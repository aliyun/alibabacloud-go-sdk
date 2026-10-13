// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyPostPayModuleSwitchShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyPostPayModuleSwitchShrinkRequest
	GetDryRun() *bool
	SetEdrModuleSwitchShrink(v string) *ModifyPostPayModuleSwitchShrinkRequest
	GetEdrModuleSwitchShrink() *string
	SetPostPaidHostAutoBind(v int32) *ModifyPostPayModuleSwitchShrinkRequest
	GetPostPaidHostAutoBind() *int32
	SetPostPaidHostAutoBindVersion(v int32) *ModifyPostPayModuleSwitchShrinkRequest
	GetPostPaidHostAutoBindVersion() *int32
	SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchShrinkRequest
	GetPostPayInstanceId() *string
	SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchShrinkRequest
	GetPostPayModuleSwitch() *string
	SetPostPayModuleSwitchObjShrink(v string) *ModifyPostPayModuleSwitchShrinkRequest
	GetPostPayModuleSwitchObjShrink() *string
}

type ModifyPostPayModuleSwitchShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: Performs a dry run. The system checks the request but does not perform the actual operation.
	//
	// - false: Performs the request normally.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The switch configuration of the EDR module.
	EdrModuleSwitchShrink *string `json:"EdrModuleSwitch,omitempty" xml:"EdrModuleSwitch,omitempty"`
	// Specifies whether to automatically bind new assets to the host and container protection feature. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	PostPaidHostAutoBind *int32 `json:"PostPaidHostAutoBind,omitempty" xml:"PostPaidHostAutoBind,omitempty"`
	// The edition to which new assets are automatically bound for host and container protection. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// example:
	//
	// 3
	PostPaidHostAutoBindVersion *int32 `json:"PostPaidHostAutoBindVersion,omitempty" xml:"PostPaidHostAutoBindVersion,omitempty"`
	// The pay-as-you-go instance ID. This parameter is required.
	//
	// > Invoke the [DescribeVersionConfig](~~DescribeVersionConfig~~) operation to obtain the value of this parameter.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// The switch status of pay-as-you-go modules. The value is a JSON string. Valid values:
	//
	// - Key:
	//
	//   - **VUL**: Vulnerability fixing module
	//
	//   - **CSPM**: Cloud Security Posture Management module
	//
	//   - **AGENTLESS**: Agentless detection module
	//
	//   - **SERVERLESS**: Serverless security module
	//
	//   - **CTDR**: Cloud Threat Detection and Response module
	//
	//   - **POST_HOST**: Host and container security module
	//
	//   - **SDK**: Malicious file detection SDK module
	//
	//   - **RASP**: Application protection module
	//
	//   - **CTDR_STORAGE**: Log management module
	//
	//   - **ANTI_RANSOMWARE**: Anti-ransomware management
	//
	// - Value: 0 indicates disabled. 1 indicates enabled.
	//
	// > The values of modules that are not specified remain unchanged.
	//
	// <notice>This parameter has the same meaning as the PostPayModuleSwitchObj parameter. If both parameters are specified, the value of PostPayModuleSwitch takes precedence.</notice>
	//
	// example:
	//
	// {"VUL":1,"CSPM":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The switches of pay-as-you-go modules.
	//
	// 	Notice:  This parameter has the same meaning as the PostPayModuleSwitch parameter. If both parameters are specified, the value of PostPayModuleSwitch takes precedence.</notice>
	PostPayModuleSwitchObjShrink *string `json:"PostPayModuleSwitchObj,omitempty" xml:"PostPayModuleSwitchObj,omitempty"`
}

func (s ModifyPostPayModuleSwitchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetEdrModuleSwitchShrink() *string {
	return s.EdrModuleSwitchShrink
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetPostPaidHostAutoBind() *int32 {
	return s.PostPaidHostAutoBind
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetPostPaidHostAutoBindVersion() *int32 {
	return s.PostPaidHostAutoBindVersion
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetPostPayInstanceId() *string {
	return s.PostPayInstanceId
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) GetPostPayModuleSwitchObjShrink() *string {
	return s.PostPayModuleSwitchObjShrink
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetClientToken(v string) *ModifyPostPayModuleSwitchShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetDryRun(v bool) *ModifyPostPayModuleSwitchShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetEdrModuleSwitchShrink(v string) *ModifyPostPayModuleSwitchShrinkRequest {
	s.EdrModuleSwitchShrink = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetPostPaidHostAutoBind(v int32) *ModifyPostPayModuleSwitchShrinkRequest {
	s.PostPaidHostAutoBind = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetPostPaidHostAutoBindVersion(v int32) *ModifyPostPayModuleSwitchShrinkRequest {
	s.PostPaidHostAutoBindVersion = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchShrinkRequest {
	s.PostPayInstanceId = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchShrinkRequest {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) SetPostPayModuleSwitchObjShrink(v string) *ModifyPostPayModuleSwitchShrinkRequest {
	s.PostPayModuleSwitchObjShrink = &v
	return s
}

func (s *ModifyPostPayModuleSwitchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
