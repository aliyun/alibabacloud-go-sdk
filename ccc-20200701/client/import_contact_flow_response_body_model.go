// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportContactFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportContactFlowResponseBody
	GetMessage() *string
	SetParams(v []*string) *ImportContactFlowResponseBody
	GetParams() []*string
	SetRequestId(v string) *ImportContactFlowResponseBody
	GetRequestId() *string
}

type ImportContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 27DD30C4-CAE2-481A-97CC-D3C54625341D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ImportContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportContactFlowResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ImportContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportContactFlowResponseBody) SetCode(v string) *ImportContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ImportContactFlowResponseBody) SetHttpStatusCode(v int32) *ImportContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportContactFlowResponseBody) SetMessage(v string) *ImportContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ImportContactFlowResponseBody) SetParams(v []*string) *ImportContactFlowResponseBody {
	s.Params = v
	return s
}

func (s *ImportContactFlowResponseBody) SetRequestId(v string) *ImportContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
