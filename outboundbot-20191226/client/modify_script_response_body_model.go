// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyScriptResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScriptResponseBody
	GetRequestId() *string
	SetScript(v *ModifyScriptResponseBodyScript) *ModifyScriptResponseBody
	GetScript() *ModifyScriptResponseBodyScript
	SetSuccess(v bool) *ModifyScriptResponseBody
	GetSuccess() *bool
}

type ModifyScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Script    *ModifyScriptResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScriptResponseBody) GetScript() *ModifyScriptResponseBodyScript {
	return s.Script
}

func (s *ModifyScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScriptResponseBody) SetCode(v string) *ModifyScriptResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScriptResponseBody) SetHttpStatusCode(v int32) *ModifyScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyScriptResponseBody) SetMessage(v string) *ModifyScriptResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScriptResponseBody) SetRequestId(v string) *ModifyScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScriptResponseBody) SetScript(v *ModifyScriptResponseBodyScript) *ModifyScriptResponseBody {
	s.Script = v
	return s
}

func (s *ModifyScriptResponseBody) SetSuccess(v bool) *ModifyScriptResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScriptResponseBody) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyScriptResponseBodyScript struct {
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// example:
	//
	// true
	IsDrafted         *bool   `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	Scene             *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// example:
	//
	// c153d0d8-ba04-41c0-8632-453944c9dd0b
	ScriptId   *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1578881227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ModifyScriptResponseBodyScript) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptResponseBodyScript) GoString() string {
	return s.String()
}

func (s *ModifyScriptResponseBodyScript) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *ModifyScriptResponseBodyScript) GetIndustry() *string {
	return s.Industry
}

func (s *ModifyScriptResponseBodyScript) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *ModifyScriptResponseBodyScript) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *ModifyScriptResponseBodyScript) GetScene() *string {
	return s.Scene
}

func (s *ModifyScriptResponseBodyScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *ModifyScriptResponseBodyScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyScriptResponseBodyScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *ModifyScriptResponseBodyScript) GetStatus() *string {
	return s.Status
}

func (s *ModifyScriptResponseBodyScript) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ModifyScriptResponseBodyScript) SetDebugStatus(v string) *ModifyScriptResponseBodyScript {
	s.DebugStatus = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetIndustry(v string) *ModifyScriptResponseBodyScript {
	s.Industry = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetIsDebugDrafted(v bool) *ModifyScriptResponseBodyScript {
	s.IsDebugDrafted = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetIsDrafted(v bool) *ModifyScriptResponseBodyScript {
	s.IsDrafted = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetScene(v string) *ModifyScriptResponseBodyScript {
	s.Scene = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetScriptDescription(v string) *ModifyScriptResponseBodyScript {
	s.ScriptDescription = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetScriptId(v string) *ModifyScriptResponseBodyScript {
	s.ScriptId = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetScriptName(v string) *ModifyScriptResponseBodyScript {
	s.ScriptName = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetStatus(v string) *ModifyScriptResponseBodyScript {
	s.Status = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) SetUpdateTime(v int64) *ModifyScriptResponseBodyScript {
	s.UpdateTime = &v
	return s
}

func (s *ModifyScriptResponseBodyScript) Validate() error {
	return dara.Validate(s)
}
