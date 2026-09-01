// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyPostPayModuleSwitchRequest
	GetClientToken() *string
	SetEdrModuleSwitch(v *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) *ModifyPostPayModuleSwitchRequest
	GetEdrModuleSwitch() *ModifyPostPayModuleSwitchRequestEdrModuleSwitch
	SetPostPaidHostAutoBind(v int32) *ModifyPostPayModuleSwitchRequest
	GetPostPaidHostAutoBind() *int32
	SetPostPaidHostAutoBindVersion(v int32) *ModifyPostPayModuleSwitchRequest
	GetPostPaidHostAutoBindVersion() *int32
	SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchRequest
	GetPostPayInstanceId() *string
	SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchRequest
	GetPostPayModuleSwitch() *string
	SetPostPayModuleSwitchObj(v *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) *ModifyPostPayModuleSwitchRequest
	GetPostPayModuleSwitchObj() *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj
}

type ModifyPostPayModuleSwitchRequest struct {
	// The client token that is used to ensure the idempotence of the request. Use a different token for each request. Only ASCII characters are supported. The token can be up to 64 characters in length.
	ClientToken     *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EdrModuleSwitch *ModifyPostPayModuleSwitchRequestEdrModuleSwitch `json:"EdrModuleSwitch,omitempty" xml:"EdrModuleSwitch,omitempty" type:"Struct"`
	// Specifies whether to automatically bind new assets for host and container protection. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	PostPaidHostAutoBind *int32 `json:"PostPaidHostAutoBind,omitempty" xml:"PostPaidHostAutoBind,omitempty"`
	// The edition to which new assets are automatically bound for host and container protection. Valid values:
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
	//   - **VUL**: vulnerability management module
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
	// <notice>This parameter has the same meaning as PostPayModuleSwitchObj. If both parameters are specified, the value of PostPayModuleSwitch takes precedence.</notice>
	//
	// example:
	//
	// {"VUL":1,"CSPM":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// The pay-as-you-go module switch.
	//
	// 	Notice: This parameter has the same meaning as PostPayModuleSwitch. If both parameters are specified, the value of PostPayModuleSwitch takes precedence.</notice>
	PostPayModuleSwitchObj *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj `json:"PostPayModuleSwitchObj,omitempty" xml:"PostPayModuleSwitchObj,omitempty" type:"Struct"`
}

func (s ModifyPostPayModuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPostPayModuleSwitchRequest) GetEdrModuleSwitch() *ModifyPostPayModuleSwitchRequestEdrModuleSwitch {
	return s.EdrModuleSwitch
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPaidHostAutoBind() *int32 {
	return s.PostPaidHostAutoBind
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPaidHostAutoBindVersion() *int32 {
	return s.PostPaidHostAutoBindVersion
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPayInstanceId() *string {
	return s.PostPayInstanceId
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPayModuleSwitchObj() *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	return s.PostPayModuleSwitchObj
}

func (s *ModifyPostPayModuleSwitchRequest) SetClientToken(v string) *ModifyPostPayModuleSwitchRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetEdrModuleSwitch(v *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) *ModifyPostPayModuleSwitchRequest {
	s.EdrModuleSwitch = v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPaidHostAutoBind(v int32) *ModifyPostPayModuleSwitchRequest {
	s.PostPaidHostAutoBind = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPaidHostAutoBindVersion(v int32) *ModifyPostPayModuleSwitchRequest {
	s.PostPaidHostAutoBindVersion = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchRequest {
	s.PostPayInstanceId = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchRequest {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPayModuleSwitchObj(v *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) *ModifyPostPayModuleSwitchRequest {
	s.PostPayModuleSwitchObj = v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) Validate() error {
	if s.EdrModuleSwitch != nil {
		if err := s.EdrModuleSwitch.Validate(); err != nil {
			return err
		}
	}
	if s.PostPayModuleSwitchObj != nil {
		if err := s.PostPayModuleSwitchObj.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPostPayModuleSwitchRequestEdrModuleSwitch struct {
	AutoBind       *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	EDR_HOST_USAGE *int64 `json:"EDR_HOST_USAGE,omitempty" xml:"EDR_HOST_USAGE,omitempty"`
}

func (s ModifyPostPayModuleSwitchRequestEdrModuleSwitch) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchRequestEdrModuleSwitch) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) GetEDR_HOST_USAGE() *int64 {
	return s.EDR_HOST_USAGE
}

func (s *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) SetAutoBind(v int32) *ModifyPostPayModuleSwitchRequestEdrModuleSwitch {
	s.AutoBind = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) SetEDR_HOST_USAGE(v int64) *ModifyPostPayModuleSwitchRequestEdrModuleSwitch {
	s.EDR_HOST_USAGE = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestEdrModuleSwitch) Validate() error {
	return dara.Validate(s)
}

type ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj struct {
	// The agentless detection module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Agentless *int32 `json:"Agentless,omitempty" xml:"Agentless,omitempty"`
	// The AI digital human module.
	//
	// example:
	//
	// 1
	AiDigital *int32 `json:"AiDigital,omitempty" xml:"AiDigital,omitempty"`
	// The anti-ransomware module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	AntiRansomware *int32 `json:"AntiRansomware,omitempty" xml:"AntiRansomware,omitempty"`
	// The basic service module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// 	Notice: The basic service module switch cannot be manually modified. This module is enabled when any other module is enabled, and is disabled when all other modules are disabled.
	//
	// example:
	//
	// 1
	BasicService *int32 `json:"BasicService,omitempty" xml:"BasicService,omitempty"`
	// The cloud security configuration check module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Cspm *int32 `json:"Cspm,omitempty" xml:"Cspm,omitempty"`
	// The threat detection and response module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Ctdr *int32 `json:"Ctdr,omitempty" xml:"Ctdr,omitempty"`
	// The log management module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	CtdrStorage *int32 `json:"CtdrStorage,omitempty" xml:"CtdrStorage,omitempty"`
	// The host and container security module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	PostHost *int32 `json:"PostHost,omitempty" xml:"PostHost,omitempty"`
	// The application protection module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Rasp *int32 `json:"Rasp,omitempty" xml:"Rasp,omitempty"`
	// The malicious file detection SDK module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Sdk *int32 `json:"Sdk,omitempty" xml:"Sdk,omitempty"`
	// The serverless security module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Serverless *int32 `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// The vulnerability management module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Vul *int32 `json:"Vul,omitempty" xml:"Vul,omitempty"`
	// The file tamper-proofing module. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	WebLock *int32 `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
}

func (s ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetAgentless() *int32 {
	return s.Agentless
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetAiDigital() *int32 {
	return s.AiDigital
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetAntiRansomware() *int32 {
	return s.AntiRansomware
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetBasicService() *int32 {
	return s.BasicService
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetCspm() *int32 {
	return s.Cspm
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetCtdr() *int32 {
	return s.Ctdr
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetCtdrStorage() *int32 {
	return s.CtdrStorage
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetPostHost() *int32 {
	return s.PostHost
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetRasp() *int32 {
	return s.Rasp
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetSdk() *int32 {
	return s.Sdk
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetServerless() *int32 {
	return s.Serverless
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetVul() *int32 {
	return s.Vul
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) GetWebLock() *int32 {
	return s.WebLock
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetAgentless(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Agentless = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetAiDigital(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.AiDigital = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetAntiRansomware(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.AntiRansomware = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetBasicService(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.BasicService = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetCspm(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Cspm = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetCtdr(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Ctdr = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetCtdrStorage(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.CtdrStorage = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetPostHost(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.PostHost = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetRasp(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Rasp = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetSdk(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Sdk = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetServerless(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Serverless = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetVul(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.Vul = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) SetWebLock(v int32) *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj {
	s.WebLock = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj) Validate() error {
	return dara.Validate(s)
}
