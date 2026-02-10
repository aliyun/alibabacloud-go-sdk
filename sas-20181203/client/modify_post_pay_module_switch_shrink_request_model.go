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
	// Automatic binding switch for new assets in host and container protection. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	PostPaidHostAutoBind *int32 `json:"PostPaidHostAutoBind,omitempty" xml:"PostPaidHostAutoBind,omitempty"`
	// Version for automatic binding of new assets in host and container protection. Values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Antivirus Edition
	//
	// - **7**: Flagship Edition
	//
	// example:
	//
	// 3
	PostPaidHostAutoBindVersion *int32 `json:"PostPaidHostAutoBindVersion,omitempty" xml:"PostPaidHostAutoBindVersion,omitempty"`
	// Pay-as-you-go instance ID, which must be filled in.
	//
	// > Call the [DescribeVersionConfig](~~DescribeVersionConfig~~) interface to obtain this parameter.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// Status of the pay-as-you-go module switch, in JsonString format. Values:
	//
	// - Key:
	//
	//   - **VUL**: Vulnerability Repair Module
	//
	//   - **CSPM**: Cloud Security Posture Management Module
	//
	//   - **AGENTLESS**: Agentless Detection Module
	//
	//   - **SERVERLESS**: Serverless Security Module
	//
	//   - **CTDR**: Threat Analysis and Response Module
	//
	//   - **POST_HOST**: Host and Container Security Module
	//
	//   - **SDK**: Malicious File Detection SDK Module
	//
	//   - **RASP**: Application Protection Module
	//
	//   - **CTDR_STORAGE**: Log Management Module
	//
	//   - **ANTI_RANSOMWARE**: Anti-Ransomware Management
	//
	// - Value: 0 means off, 1 means on
	//
	// > The values of modules not passed will not change.
	//
	// <notice>The meaning is the same as the PostPayModuleSwitchObj field. When both exist, the value of PostPayModuleSwitch takes precedence.
	//
	// example:
	//
	// {"VUL":1,"CSPM":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// Pay-as-you-go module switch.
	//
	// 	Notice:  The meaning is the same as the PostPayModuleSwitch field. When both exist, the value of PostPayModuleSwitch takes precedence.
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
