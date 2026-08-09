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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The NLU robot engine.
	//
	// - If this parameter is left empty, small model scenarios are queried.
	//
	// - If this parameter is set to Prompts, the text input mode under large model scenarios is queried.
	//
	// - If this parameter is set to SSE_FUNCTION, the function computing mode under large model scenarios is queried.
	//
	// - If this parameter is set to BeeBot, the workflow configuration mode under large model scenarios is queried.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The script name.
	//
	// example:
	//
	// Course satisfaction follow-up
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
