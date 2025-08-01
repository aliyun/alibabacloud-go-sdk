// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFlowRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteWebFlowRulesResponseBody
	GetCode() *int32
	SetData(v []*int64) *DeleteWebFlowRulesResponseBody
	GetData() []*int64
	SetHttpStatusCode(v int32) *DeleteWebFlowRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteWebFlowRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWebFlowRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWebFlowRulesResponseBody
	GetSuccess() *bool
}

type DeleteWebFlowRulesResponseBody struct {
	// example:
	//
	// 200
	Code *int32   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 354FF159-E9FB-5FBA-BDD5-E99EE440A88D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWebFlowRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFlowRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebFlowRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteWebFlowRulesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *DeleteWebFlowRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWebFlowRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWebFlowRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebFlowRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWebFlowRulesResponseBody) SetCode(v int32) *DeleteWebFlowRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) SetData(v []*int64) *DeleteWebFlowRulesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) SetHttpStatusCode(v int32) *DeleteWebFlowRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) SetMessage(v string) *DeleteWebFlowRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) SetRequestId(v string) *DeleteWebFlowRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) SetSuccess(v bool) *DeleteWebFlowRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWebFlowRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
