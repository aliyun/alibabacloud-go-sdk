// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
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
	PostPayModuleSwitchObj *ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj `json:"PostPayModuleSwitchObj,omitempty" xml:"PostPayModuleSwitchObj,omitempty" type:"Struct"`
}

func (s ModifyPostPayModuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchRequest) GoString() string {
	return s.String()
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
	if s.PostPayModuleSwitchObj != nil {
		if err := s.PostPayModuleSwitchObj.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPostPayModuleSwitchRequestPostPayModuleSwitchObj struct {
	// Agentless Detection Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Agentless *int32 `json:"Agentless,omitempty" xml:"Agentless,omitempty"`
	// Anti-Ransomware Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	AntiRansomware *int32 `json:"AntiRansomware,omitempty" xml:"AntiRansomware,omitempty"`
	// Basic service module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// 	Notice: The basic service module switch does not support active modification. When other modules are on, this module is also on. If all other modules are off, then this module is off.
	//
	// example:
	//
	// 1
	BasicService *int32 `json:"BasicService,omitempty" xml:"BasicService,omitempty"`
	// Cloud Security Configuration Check Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Cspm *int32 `json:"Cspm,omitempty" xml:"Cspm,omitempty"`
	// Threat Analysis and Response Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Ctdr *int32 `json:"Ctdr,omitempty" xml:"Ctdr,omitempty"`
	// Log Management Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	CtdrStorage *int32 `json:"CtdrStorage,omitempty" xml:"CtdrStorage,omitempty"`
	// Host and Container Security Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	PostHost *int32 `json:"PostHost,omitempty" xml:"PostHost,omitempty"`
	// Application Protection Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Rasp *int32 `json:"Rasp,omitempty" xml:"Rasp,omitempty"`
	// Malicious File Detection SDK Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Sdk *int32 `json:"Sdk,omitempty" xml:"Sdk,omitempty"`
	// Serverless Security Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Serverless *int32 `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// Vulnerability Repair Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	Vul *int32 `json:"Vul,omitempty" xml:"Vul,omitempty"`
	// File Tamper Protection Module. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
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
