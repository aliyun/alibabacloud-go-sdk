// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// Specifies whether to automatically bind newly added assets for host and container protection. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	PostPaidHostAutoBind *int32 `json:"PostPaidHostAutoBind,omitempty" xml:"PostPaidHostAutoBind,omitempty"`
	// The version to which newly added assets are automatically bound for host and container protection. Valid values:
	//
	// - **1**: Free Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **5**: Advanced Edition.
	//
	// - **6**: Anti-virus Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// example:
	//
	// 3
	PostPaidHostAutoBindVersion *int32 `json:"PostPaidHostAutoBindVersion,omitempty" xml:"PostPaidHostAutoBindVersion,omitempty"`
	// The pay-as-you-go instance ID. This parameter is required.
	//
	// > Invoke the [DescribeVersionConfig](~~DescribeVersionConfig~~) operation to obtain this parameter.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// The switch status of pay-as-you-go modules in JSON string format. Valid values:
	//
	// - Key:
	//
	//   - **VUL**: vulnerability fix module
	//
	//   - **CSPM**: Cloud Security Posture Management (CSPM) module
	//
	//   - **AGENTLESS**: agentless detection module
	//
	//   - **SERVERLESS**: serverless security module
	//
	//   - **CTDR**: threat detection and response module
	//
	//   - **POST_HOST**: host and container security module
	//
	//   - **SDK**: malicious file detection SDK module
	//
	//   - **RASP**: application protection module
	//
	//   - **CTDR_STORAGE**: log management module
	//
	//   - **ANTI_RANSOMWARE**: anti-ransomware management
	//
	// - Value: 0 indicates disabled. 1 indicates enabled.
	//
	// > Modules for which no value is specified remain unchanged.
	//
	// <notice>This parameter has the same meaning as PostPayModuleSwitchObj. If both parameters are specified, the value of PostPayModuleSwitch takes precedence..
	//
	// example:
	//
	// {"VUL":1,"CSPM":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The pay-as-you-go module switch.
	//
	// 	Notice: This parameter has the same meaning as PostPayModuleSwitch. If both parameters are specified, the value of PostPayModuleSwitch takes precedence..
	PostPayModuleSwitchObjShrink *string `json:"PostPayModuleSwitchObj,omitempty" xml:"PostPayModuleSwitchObj,omitempty"`
}

func (s ModifyPostPayModuleSwitchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchShrinkRequest) GoString() string {
	return s.String()
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
