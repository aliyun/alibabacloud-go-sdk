// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptVoiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeScriptVoiceConfigRequest
	GetInstanceId() *string
	SetScriptId(v string) *DescribeScriptVoiceConfigRequest
	GetScriptId() *string
	SetScriptVoiceConfigId(v string) *DescribeScriptVoiceConfigRequest
	GetScriptVoiceConfigId() *string
}

type DescribeScriptVoiceConfigRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0fe7f71c-8771-42ef-9bb1-19aa16ae7120
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Scenario voice configuration ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a07b634-e15d-445f-bbcb-fc4ea2df7b87
	ScriptVoiceConfigId *string `json:"ScriptVoiceConfigId,omitempty" xml:"ScriptVoiceConfigId,omitempty"`
}

func (s DescribeScriptVoiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptVoiceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeScriptVoiceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScriptVoiceConfigRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeScriptVoiceConfigRequest) GetScriptVoiceConfigId() *string {
	return s.ScriptVoiceConfigId
}

func (s *DescribeScriptVoiceConfigRequest) SetInstanceId(v string) *DescribeScriptVoiceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeScriptVoiceConfigRequest) SetScriptId(v string) *DescribeScriptVoiceConfigRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeScriptVoiceConfigRequest) SetScriptVoiceConfigId(v string) *DescribeScriptVoiceConfigRequest {
	s.ScriptVoiceConfigId = &v
	return s
}

func (s *DescribeScriptVoiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
