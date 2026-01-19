// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDemoInstanceExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckDemoInstanceExistsResponseBody
	GetCode() *string
	SetData(v bool) *CheckDemoInstanceExistsResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckDemoInstanceExistsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckDemoInstanceExistsResponseBody
	GetMessage() *string
	SetParams(v []*string) *CheckDemoInstanceExistsResponseBody
	GetParams() []*string
	SetRequestId(v string) *CheckDemoInstanceExistsResponseBody
	GetRequestId() *string
}

type CheckDemoInstanceExistsResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool     `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDemoInstanceExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDemoInstanceExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDemoInstanceExistsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckDemoInstanceExistsResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckDemoInstanceExistsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckDemoInstanceExistsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckDemoInstanceExistsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CheckDemoInstanceExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDemoInstanceExistsResponseBody) SetCode(v string) *CheckDemoInstanceExistsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetData(v bool) *CheckDemoInstanceExistsResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetHttpStatusCode(v int32) *CheckDemoInstanceExistsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetMessage(v string) *CheckDemoInstanceExistsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetParams(v []*string) *CheckDemoInstanceExistsResponseBody {
	s.Params = v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetRequestId(v string) *CheckDemoInstanceExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
