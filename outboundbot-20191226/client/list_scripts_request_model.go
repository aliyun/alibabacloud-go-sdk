// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptsRequest
	GetInstanceId() *string
	SetNluEngine(v string) *ListScriptsRequest
	GetNluEngine() *string
	SetPageNumber(v int32) *ListScriptsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptsRequest
	GetPageSize() *int32
	SetScriptName(v string) *ListScriptsRequest
	GetScriptName() *string
}

type ListScriptsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// NLU bot engine.
	//
	// - If left unspecified, queries scenarios using small models.
	//
	// - If set to "Prompts", queries text-input pattern under LLM scenarios.
	//
	// - If set to "SSE_FUNCTION", queries Function Compute pattern under LLM scenarios.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries displayed per page
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Script name
	//
	// example:
	//
	// 课程满意度回访
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
}

func (s ListScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsRequest) GoString() string {
	return s.String()
}

func (s *ListScriptsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptsRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListScriptsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListScriptsRequest) SetInstanceId(v string) *ListScriptsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptsRequest) SetNluEngine(v string) *ListScriptsRequest {
	s.NluEngine = &v
	return s
}

func (s *ListScriptsRequest) SetPageNumber(v int32) *ListScriptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsRequest) SetPageSize(v int32) *ListScriptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptsRequest) SetScriptName(v string) *ListScriptsRequest {
	s.ScriptName = &v
	return s
}

func (s *ListScriptsRequest) Validate() error {
	return dara.Validate(s)
}
