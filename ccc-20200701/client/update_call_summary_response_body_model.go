// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCallSummaryResponseBody
	GetCode() *string
	SetData(v interface{}) *UpdateCallSummaryResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *UpdateCallSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCallSummaryResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateCallSummaryResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateCallSummaryResponseBody
	GetRequestId() *string
}

type UpdateCallSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9FBA26B0-462B-4D77-B78F-AF35560DBC71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCallSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCallSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCallSummaryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateCallSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCallSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCallSummaryResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateCallSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCallSummaryResponseBody) SetCode(v string) *UpdateCallSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCallSummaryResponseBody) SetData(v interface{}) *UpdateCallSummaryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCallSummaryResponseBody) SetHttpStatusCode(v int32) *UpdateCallSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCallSummaryResponseBody) SetMessage(v string) *UpdateCallSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCallSummaryResponseBody) SetParams(v []*string) *UpdateCallSummaryResponseBody {
	s.Params = v
	return s
}

func (s *UpdateCallSummaryResponseBody) SetRequestId(v string) *UpdateCallSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCallSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
