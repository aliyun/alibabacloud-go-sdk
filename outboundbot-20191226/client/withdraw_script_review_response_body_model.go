// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawScriptReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WithdrawScriptReviewResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *WithdrawScriptReviewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *WithdrawScriptReviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *WithdrawScriptReviewResponseBody
	GetRequestId() *string
	SetScript(v *WithdrawScriptReviewResponseBodyScript) *WithdrawScriptReviewResponseBody
	GetScript() *WithdrawScriptReviewResponseBodyScript
	SetSuccess(v bool) *WithdrawScriptReviewResponseBody
	GetSuccess() *bool
}

type WithdrawScriptReviewResponseBody struct {
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Script    *WithdrawScriptReviewResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WithdrawScriptReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawScriptReviewResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawScriptReviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *WithdrawScriptReviewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WithdrawScriptReviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WithdrawScriptReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawScriptReviewResponseBody) GetScript() *WithdrawScriptReviewResponseBodyScript {
	return s.Script
}

func (s *WithdrawScriptReviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WithdrawScriptReviewResponseBody) SetCode(v string) *WithdrawScriptReviewResponseBody {
	s.Code = &v
	return s
}

func (s *WithdrawScriptReviewResponseBody) SetHttpStatusCode(v int32) *WithdrawScriptReviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WithdrawScriptReviewResponseBody) SetMessage(v string) *WithdrawScriptReviewResponseBody {
	s.Message = &v
	return s
}

func (s *WithdrawScriptReviewResponseBody) SetRequestId(v string) *WithdrawScriptReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawScriptReviewResponseBody) SetScript(v *WithdrawScriptReviewResponseBodyScript) *WithdrawScriptReviewResponseBody {
	s.Script = v
	return s
}

func (s *WithdrawScriptReviewResponseBody) SetSuccess(v bool) *WithdrawScriptReviewResponseBody {
	s.Success = &v
	return s
}

func (s *WithdrawScriptReviewResponseBody) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WithdrawScriptReviewResponseBodyScript struct {
	// example:
	//
	// PUBLISHED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// false
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// example:
	//
	// false
	IsDrafted         *bool   `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	Scene             *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// example:
	//
	// e4e2a770-b97b-465a-80d8-06dca008c503
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

func (s WithdrawScriptReviewResponseBodyScript) String() string {
	return dara.Prettify(s)
}

func (s WithdrawScriptReviewResponseBodyScript) GoString() string {
	return s.String()
}

func (s *WithdrawScriptReviewResponseBodyScript) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *WithdrawScriptReviewResponseBodyScript) GetIndustry() *string {
	return s.Industry
}

func (s *WithdrawScriptReviewResponseBodyScript) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *WithdrawScriptReviewResponseBodyScript) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *WithdrawScriptReviewResponseBodyScript) GetScene() *string {
	return s.Scene
}

func (s *WithdrawScriptReviewResponseBodyScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *WithdrawScriptReviewResponseBodyScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *WithdrawScriptReviewResponseBodyScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *WithdrawScriptReviewResponseBodyScript) GetStatus() *string {
	return s.Status
}

func (s *WithdrawScriptReviewResponseBodyScript) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *WithdrawScriptReviewResponseBodyScript) SetDebugStatus(v string) *WithdrawScriptReviewResponseBodyScript {
	s.DebugStatus = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetIndustry(v string) *WithdrawScriptReviewResponseBodyScript {
	s.Industry = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetIsDebugDrafted(v bool) *WithdrawScriptReviewResponseBodyScript {
	s.IsDebugDrafted = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetIsDrafted(v bool) *WithdrawScriptReviewResponseBodyScript {
	s.IsDrafted = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetScene(v string) *WithdrawScriptReviewResponseBodyScript {
	s.Scene = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetScriptDescription(v string) *WithdrawScriptReviewResponseBodyScript {
	s.ScriptDescription = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetScriptId(v string) *WithdrawScriptReviewResponseBodyScript {
	s.ScriptId = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetScriptName(v string) *WithdrawScriptReviewResponseBodyScript {
	s.ScriptName = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetStatus(v string) *WithdrawScriptReviewResponseBodyScript {
	s.Status = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) SetUpdateTime(v int64) *WithdrawScriptReviewResponseBodyScript {
	s.UpdateTime = &v
	return s
}

func (s *WithdrawScriptReviewResponseBodyScript) Validate() error {
	return dara.Validate(s)
}
