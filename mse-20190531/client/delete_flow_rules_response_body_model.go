// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteFlowRulesResponseBody
	GetCode() *int32
	SetData(v []*int64) *DeleteFlowRulesResponseBody
	GetData() []*int64
	SetHttpStatusCode(v int32) *DeleteFlowRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFlowRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFlowRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFlowRulesResponseBody
	GetSuccess() *bool
}

type DeleteFlowRulesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the rules that were deleted.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FAF577DD-8E8E-5BE6-80CC-380E19******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteFlowRulesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *DeleteFlowRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFlowRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFlowRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlowRulesResponseBody) SetCode(v int32) *DeleteFlowRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowRulesResponseBody) SetData(v []*int64) *DeleteFlowRulesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFlowRulesResponseBody) SetHttpStatusCode(v int32) *DeleteFlowRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFlowRulesResponseBody) SetMessage(v string) *DeleteFlowRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowRulesResponseBody) SetRequestId(v string) *DeleteFlowRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowRulesResponseBody) SetSuccess(v bool) *DeleteFlowRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
