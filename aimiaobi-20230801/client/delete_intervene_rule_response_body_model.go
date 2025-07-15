// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterveneRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInterveneRuleResponseBody
	GetCode() *string
	SetData(v *DeleteInterveneRuleResponseBodyData) *DeleteInterveneRuleResponseBody
	GetData() *DeleteInterveneRuleResponseBodyData
	SetHttpStatusCode(v int32) *DeleteInterveneRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInterveneRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInterveneRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInterveneRuleResponseBody
	GetSuccess() *bool
}

type DeleteInterveneRuleResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInterveneRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInterveneRuleResponseBody) GetData() *DeleteInterveneRuleResponseBodyData {
	return s.Data
}

func (s *DeleteInterveneRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInterveneRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInterveneRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInterveneRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInterveneRuleResponseBody) SetCode(v string) *DeleteInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetData(v *DeleteInterveneRuleResponseBodyData) *DeleteInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetHttpStatusCode(v int32) *DeleteInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetMessage(v string) *DeleteInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetRequestId(v string) *DeleteInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetSuccess(v bool) *DeleteInterveneRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteInterveneRuleResponseBodyData struct {
	Code       *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// dt-s50ntwtywb4y
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteInterveneRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *DeleteInterveneRuleResponseBodyData) GetFailIdList() []*string {
	return s.FailIdList
}

func (s *DeleteInterveneRuleResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteInterveneRuleResponseBodyData) SetCode(v int32) *DeleteInterveneRuleResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteInterveneRuleResponseBodyData) SetFailIdList(v []*string) *DeleteInterveneRuleResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *DeleteInterveneRuleResponseBodyData) SetTaskId(v string) *DeleteInterveneRuleResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteInterveneRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
