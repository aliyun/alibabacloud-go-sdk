// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeTrafficResponseBody
	GetCode() *string
	SetData(v *ResumeTrafficResponseBodyData) *ResumeTrafficResponseBody
	GetData() *ResumeTrafficResponseBodyData
	SetErrorCode(v string) *ResumeTrafficResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ResumeTrafficResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeTrafficResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ResumeTrafficResponseBody
	GetSuccess() *string
	SetTraceId(v string) *ResumeTrafficResponseBody
	GetTraceId() *string
}

type ResumeTrafficResponseBody struct {
	// example:
	//
	// 200
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ResumeTrafficResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ResumeTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTrafficResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeTrafficResponseBody) GetData() *ResumeTrafficResponseBodyData {
	return s.Data
}

func (s *ResumeTrafficResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResumeTrafficResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeTrafficResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ResumeTrafficResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ResumeTrafficResponseBody) SetCode(v string) *ResumeTrafficResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetData(v *ResumeTrafficResponseBodyData) *ResumeTrafficResponseBody {
	s.Data = v
	return s
}

func (s *ResumeTrafficResponseBody) SetErrorCode(v string) *ResumeTrafficResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetMessage(v string) *ResumeTrafficResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetRequestId(v string) *ResumeTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetSuccess(v string) *ResumeTrafficResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetTraceId(v string) *ResumeTrafficResponseBody {
	s.TraceId = &v
	return s
}

func (s *ResumeTrafficResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResumeTrafficResponseBodyData struct {
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResumeTrafficResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResumeTrafficResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *ResumeTrafficResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeTrafficResponseBodyData) SetMsg(v string) *ResumeTrafficResponseBodyData {
	s.Msg = &v
	return s
}

func (s *ResumeTrafficResponseBodyData) SetSuccess(v bool) *ResumeTrafficResponseBodyData {
	s.Success = &v
	return s
}

func (s *ResumeTrafficResponseBodyData) Validate() error {
	return dara.Validate(s)
}
