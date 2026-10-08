// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPipelineIntegratedTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartPipelineIntegratedTaskResponseBody
	GetCode() *string
	SetData(v string) *StartPipelineIntegratedTaskResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *StartPipelineIntegratedTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartPipelineIntegratedTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartPipelineIntegratedTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartPipelineIntegratedTaskResponseBody
	GetSuccess() *bool
}

type StartPipelineIntegratedTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartPipelineIntegratedTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPipelineIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartPipelineIntegratedTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartPipelineIntegratedTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *StartPipelineIntegratedTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartPipelineIntegratedTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartPipelineIntegratedTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPipelineIntegratedTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartPipelineIntegratedTaskResponseBody) SetCode(v string) *StartPipelineIntegratedTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) SetData(v string) *StartPipelineIntegratedTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) SetHttpStatusCode(v int32) *StartPipelineIntegratedTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) SetMessage(v string) *StartPipelineIntegratedTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) SetRequestId(v string) *StartPipelineIntegratedTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) SetSuccess(v bool) *StartPipelineIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartPipelineIntegratedTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
