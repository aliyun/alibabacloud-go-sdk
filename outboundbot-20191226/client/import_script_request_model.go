// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ImportScriptRequest
	GetInstanceId() *string
	SetNluEngine(v string) *ImportScriptRequest
	GetNluEngine() *string
	SetSignatureUrl(v string) *ImportScriptRequest
	GetSignatureUrl() *string
}

type ImportScriptRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00b37342-e759-4fe5-b296-aef775933af0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// NLU engine. Applies only to Large Language Model (LLM) scenarios.
	//
	// Valid values:
	//
	// - Prompts: LLM scenario.
	//
	// - Empty: Non-LLM scenario.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// URL of the scenario JSON file to import.
	//
	// > This value is the Folder parameter returned by the GetJobDataUploadParams operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPLOADED/SCRIPT/136a055e-***-46c6-853a-731b3a2973de/18dc6d79-338f-413c-a5a8-dcce5f05b41a_9bd7916d-a340-4925-a911-92390cbe0f33.json
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
}

func (s ImportScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportScriptRequest) GoString() string {
	return s.String()
}

func (s *ImportScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ImportScriptRequest) GetSignatureUrl() *string {
	return s.SignatureUrl
}

func (s *ImportScriptRequest) SetInstanceId(v string) *ImportScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportScriptRequest) SetNluEngine(v string) *ImportScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *ImportScriptRequest) SetSignatureUrl(v string) *ImportScriptRequest {
	s.SignatureUrl = &v
	return s
}

func (s *ImportScriptRequest) Validate() error {
	return dara.Validate(s)
}
