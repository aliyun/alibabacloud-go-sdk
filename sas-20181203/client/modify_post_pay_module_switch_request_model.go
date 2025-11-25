// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPostPayModuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchRequest
	GetPostPayInstanceId() *string
	SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchRequest
	GetPostPayModuleSwitch() *string
}

type ModifyPostPayModuleSwitchRequest struct {
	// The ID of the pay-as-you-go instance. This parameter is required.
	//
	// >  You can call the [DescribeVersionConfig](https://help.aliyun.com/document_detail/421770.html) operation to obtain the ID.
	//
	// example:
	//
	// postpay-sas-**
	PostPayInstanceId *string `json:"PostPayInstanceId,omitempty" xml:"PostPayInstanceId,omitempty"`
	// The switch status of the pay-as-you-go module. The value is a JSON string. Valid values:
	//
	// 	- Key:
	//
	//     	- **VUL**: vulnerability fixing module
	//
	//     	- **CSPM**: cloud service configuration check module
	//
	//     	- **AGENTLESS**: agentless detection module
	//
	//     	- **SERVERLESS**: serverless asset module
	//
	//     	- **CTDR**: Threat Analysis and Response Module
	//
	//     	- **POST_HOST**: Host and Container Security Module
	//
	//     	- **SDK**: Malicious File Detection SDK Module
	//
	//     	- **RASP**: Application Protection Module
	//
	// 	- Value: A value of 0 specifies disabled. A value of 1 specifies enabled.
	//
	// >  If you do not specify a value for a module, the original value of the module is retained.
	//
	// example:
	//
	// {"VUL":1,"CSPM":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
}

func (s ModifyPostPayModuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPostPayModuleSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPayInstanceId() *string {
	return s.PostPayInstanceId
}

func (s *ModifyPostPayModuleSwitchRequest) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPayInstanceId(v string) *ModifyPostPayModuleSwitchRequest {
	s.PostPayInstanceId = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) SetPostPayModuleSwitch(v string) *ModifyPostPayModuleSwitchRequest {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *ModifyPostPayModuleSwitchRequest) Validate() error {
	return dara.Validate(s)
}
