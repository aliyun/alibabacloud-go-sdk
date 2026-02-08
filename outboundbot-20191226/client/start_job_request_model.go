// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *StartJobRequest
	GetCallingNumber() []*string
	SetInstanceId(v string) *StartJobRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *StartJobRequest
	GetJobGroupId() *string
	SetJobJson(v string) *StartJobRequest
	GetJobJson() *string
	SetScenarioId(v string) *StartJobRequest
	GetScenarioId() *string
	SetScriptId(v string) *StartJobRequest
	GetScriptId() *string
}

type StartJobRequest struct {
	// List of calling numbers
	//
	// example:
	//
	// 10086
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID
	//
	// example:
	//
	// 4f21446e-324e-46f2-bf62-7f341fb004ea
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job data, including the following content:
	//
	// extras // Optional additional information
	//
	// - key // Key of the custom variable content
	//
	// - value // Value of the custom variable content
	//
	// contacts // Required field; callee contact information
	//
	// - phoneNumber // Required field; callee phone number
	//
	// - referenceId // Required field; contact ID (user-defined), primarily used to avoid duplicate names
	//
	// - role // Optional field; role of this contact in the outbound call scenario; if unspecified, no distinction is made
	//
	// - honorific // Optional field; honorific title for the contact
	//
	// - name // Optional field; contact name
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"extras": [{
	//
	// 		"value": "测试",
	//
	// 		"key": "测试"
	//
	// 	}, {
	//
	// 		"key": "ServiceId",
	//
	// 		"value": null
	//
	// 	}, {
	//
	// 		"key": "TenantId",
	//
	// 		"value": null
	//
	// 	}],
	//
	// 	"contacts": [{
	//
	// 		"phonenumber": "198********",
	//
	// 		"name": "张三",
	//
	// 		"priority": 5,
	//
	// 		"referenceId": "C01"
	//
	// 	}]
	//
	// }
	JobJson *string `json:"JobJson,omitempty" xml:"JobJson,omitempty"`
	// Scenario ID, used for compatibility with the legacy outbound calling feature. This parameter is optional.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Scenario ID (Required)
	//
	// example:
	//
	// 1eeae263-54ca-4a52-ae14-62da58f3c473
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s StartJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartJobRequest) GoString() string {
	return s.String()
}

func (s *StartJobRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *StartJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartJobRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *StartJobRequest) GetJobJson() *string {
	return s.JobJson
}

func (s *StartJobRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *StartJobRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *StartJobRequest) SetCallingNumber(v []*string) *StartJobRequest {
	s.CallingNumber = v
	return s
}

func (s *StartJobRequest) SetInstanceId(v string) *StartJobRequest {
	s.InstanceId = &v
	return s
}

func (s *StartJobRequest) SetJobGroupId(v string) *StartJobRequest {
	s.JobGroupId = &v
	return s
}

func (s *StartJobRequest) SetJobJson(v string) *StartJobRequest {
	s.JobJson = &v
	return s
}

func (s *StartJobRequest) SetScenarioId(v string) *StartJobRequest {
	s.ScenarioId = &v
	return s
}

func (s *StartJobRequest) SetScriptId(v string) *StartJobRequest {
	s.ScriptId = &v
	return s
}

func (s *StartJobRequest) Validate() error {
	return dara.Validate(s)
}
