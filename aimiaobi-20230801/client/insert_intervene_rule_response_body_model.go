// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertInterveneRuleResponseBody
	GetCode() *string
	SetData(v *InsertInterveneRuleResponseBodyData) *InsertInterveneRuleResponseBody
	GetData() *InsertInterveneRuleResponseBodyData
	SetHttpStatusCode(v int32) *InsertInterveneRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InsertInterveneRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertInterveneRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsertInterveneRuleResponseBody
	GetSuccess() *bool
}

type InsertInterveneRuleResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DD656AF9-0839-521A-A3D2-F320009F9C87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertInterveneRuleResponseBody) GetData() *InsertInterveneRuleResponseBodyData {
	return s.Data
}

func (s *InsertInterveneRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InsertInterveneRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertInterveneRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertInterveneRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertInterveneRuleResponseBody) SetCode(v string) *InsertInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetData(v *InsertInterveneRuleResponseBodyData) *InsertInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetHttpStatusCode(v int32) *InsertInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetMessage(v string) *InsertInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetRequestId(v string) *InsertInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetSuccess(v bool) *InsertInterveneRuleResponseBody {
	s.Success = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsertInterveneRuleResponseBodyData struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s InsertInterveneRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *InsertInterveneRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *InsertInterveneRuleResponseBodyData) SetCode(v int32) *InsertInterveneRuleResponseBodyData {
	s.Code = &v
	return s
}

func (s *InsertInterveneRuleResponseBodyData) SetRuleId(v int64) *InsertInterveneRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *InsertInterveneRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
