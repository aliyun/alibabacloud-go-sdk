// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScriptResponseBody
	GetRequestId() *string
	SetScript(v *CreateScriptResponseBodyScript) *CreateScriptResponseBody
	GetScript() *CreateScriptResponseBodyScript
	SetSuccess(v bool) *CreateScriptResponseBody
	GetSuccess() *bool
}

type CreateScriptResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Script information
	Script *CreateScriptResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptResponseBody) GetScript() *CreateScriptResponseBodyScript {
	return s.Script
}

func (s *CreateScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScriptResponseBody) SetCode(v string) *CreateScriptResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptResponseBody) SetHttpStatusCode(v int32) *CreateScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptResponseBody) SetMessage(v string) *CreateScriptResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptResponseBody) SetRequestId(v string) *CreateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptResponseBody) SetScript(v *CreateScriptResponseBodyScript) *CreateScriptResponseBody {
	s.Script = v
	return s
}

func (s *CreateScriptResponseBody) SetSuccess(v bool) *CreateScriptResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScriptResponseBody) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScriptResponseBodyScript struct {
	// Script debug status
	//
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// Industry
	//
	// example:
	//
	// 教育
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the script is a debug draft
	//
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// Indicates whether the script is a draft.
	//
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// NLU access method (applicable only to LLM scenarios). Enumeration: Managed—uses an Alibaba public account for access. This field is empty for non-LLM scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// NLU engine (applicable only to LLM scenarios). Enumeration: Prompts—used for LLM scenarios. This field is empty for non-LLM scenarios.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Associated scenario
	//
	// example:
	//
	// 回访
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Script description
	//
	// example:
	//
	// 课程满意度回访
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Script ID
	//
	// example:
	//
	// 8c58d3e0-bf27-4685-a5a5-65872ec5abc4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Script name
	//
	// example:
	//
	// 课程满意度回访
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Job status. Valid values:
	//
	// - **DRAFTED**: Draft.
	//
	// - **INITIALIZE_IN_PROGRESS**: Initializing.
	//
	// - **PUBLISHED**: Published.
	//
	// - **PUBLISH_IN_PROGRESS**: Publishing.
	//
	// - **ROLLBACK_IN_PROGRESS**: Rolling back.
	//
	// - **EXAMINE_IN_PROGRESS**: Pending review.
	//
	// - **PUBLISHED_AND_EXAMINE_IN_PROGRESS**: Published and pending review.
	//
	// - **PUBLISH_FAILED**: Publish failed.
	//
	// - **ROLLBACK_FAILED**: Rollback failed.
	//
	// - **IMPORT_IN_PROGRESS**: Importing.
	//
	// - **IMPORT_FAILED**: Import failed.
	//
	// example:
	//
	// DRAFTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Script update time (in milliseconds)
	//
	// example:
	//
	// 1578474045152
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateScriptResponseBodyScript) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponseBodyScript) GoString() string {
	return s.String()
}

func (s *CreateScriptResponseBodyScript) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *CreateScriptResponseBodyScript) GetIndustry() *string {
	return s.Industry
}

func (s *CreateScriptResponseBodyScript) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *CreateScriptResponseBodyScript) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *CreateScriptResponseBodyScript) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *CreateScriptResponseBodyScript) GetNluEngine() *string {
	return s.NluEngine
}

func (s *CreateScriptResponseBodyScript) GetScene() *string {
	return s.Scene
}

func (s *CreateScriptResponseBodyScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *CreateScriptResponseBodyScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateScriptResponseBodyScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateScriptResponseBodyScript) GetStatus() *string {
	return s.Status
}

func (s *CreateScriptResponseBodyScript) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateScriptResponseBodyScript) SetDebugStatus(v string) *CreateScriptResponseBodyScript {
	s.DebugStatus = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetIndustry(v string) *CreateScriptResponseBodyScript {
	s.Industry = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetIsDebugDrafted(v bool) *CreateScriptResponseBodyScript {
	s.IsDebugDrafted = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetIsDrafted(v bool) *CreateScriptResponseBodyScript {
	s.IsDrafted = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetNluAccessType(v string) *CreateScriptResponseBodyScript {
	s.NluAccessType = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetNluEngine(v string) *CreateScriptResponseBodyScript {
	s.NluEngine = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetScene(v string) *CreateScriptResponseBodyScript {
	s.Scene = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetScriptDescription(v string) *CreateScriptResponseBodyScript {
	s.ScriptDescription = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetScriptId(v string) *CreateScriptResponseBodyScript {
	s.ScriptId = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetScriptName(v string) *CreateScriptResponseBodyScript {
	s.ScriptName = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetStatus(v string) *CreateScriptResponseBodyScript {
	s.Status = &v
	return s
}

func (s *CreateScriptResponseBodyScript) SetUpdateTime(v int64) *CreateScriptResponseBodyScript {
	s.UpdateTime = &v
	return s
}

func (s *CreateScriptResponseBodyScript) Validate() error {
	return dara.Validate(s)
}
