// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCallSummaryResponseBody
	GetCode() *string
	SetData(v string) *CreateCallSummaryResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateCallSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCallSummaryResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateCallSummaryResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateCallSummaryResponseBody
	GetRequestId() *string
}

type CreateCallSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// c58b9719-3bc3-441d-a4d3-fc0309ef7066
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCallSummaryResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCallSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCallSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCallSummaryResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateCallSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallSummaryResponseBody) SetCode(v string) *CreateCallSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallSummaryResponseBody) SetData(v string) *CreateCallSummaryResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallSummaryResponseBody) SetHttpStatusCode(v int32) *CreateCallSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCallSummaryResponseBody) SetMessage(v string) *CreateCallSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCallSummaryResponseBody) SetParams(v []*string) *CreateCallSummaryResponseBody {
	s.Params = v
	return s
}

func (s *CreateCallSummaryResponseBody) SetRequestId(v string) *CreateCallSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
