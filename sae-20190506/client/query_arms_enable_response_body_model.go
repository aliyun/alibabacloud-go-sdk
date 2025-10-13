// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryArmsEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryArmsEnableResponseBody
	GetCode() *string
	SetData(v *QueryArmsEnableResponseBodyData) *QueryArmsEnableResponseBody
	GetData() *QueryArmsEnableResponseBodyData
	SetErrorCode(v string) *QueryArmsEnableResponseBody
	GetErrorCode() *string
	SetMessage(v string) *QueryArmsEnableResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryArmsEnableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryArmsEnableResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryArmsEnableResponseBody
	GetTraceId() *string
}

type QueryArmsEnableResponseBody struct {
	// example:
	//
	// 200
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryArmsEnableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s QueryArmsEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryArmsEnableResponseBody) GoString() string {
	return s.String()
}

func (s *QueryArmsEnableResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryArmsEnableResponseBody) GetData() *QueryArmsEnableResponseBodyData {
	return s.Data
}

func (s *QueryArmsEnableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryArmsEnableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryArmsEnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryArmsEnableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryArmsEnableResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryArmsEnableResponseBody) SetCode(v string) *QueryArmsEnableResponseBody {
	s.Code = &v
	return s
}

func (s *QueryArmsEnableResponseBody) SetData(v *QueryArmsEnableResponseBodyData) *QueryArmsEnableResponseBody {
	s.Data = v
	return s
}

func (s *QueryArmsEnableResponseBody) SetErrorCode(v string) *QueryArmsEnableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryArmsEnableResponseBody) SetMessage(v string) *QueryArmsEnableResponseBody {
	s.Message = &v
	return s
}

func (s *QueryArmsEnableResponseBody) SetRequestId(v string) *QueryArmsEnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryArmsEnableResponseBody) SetSuccess(v bool) *QueryArmsEnableResponseBody {
	s.Success = &v
	return s
}

func (s *QueryArmsEnableResponseBody) SetTraceId(v string) *QueryArmsEnableResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryArmsEnableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryArmsEnableResponseBodyData struct {
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s QueryArmsEnableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryArmsEnableResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryArmsEnableResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *QueryArmsEnableResponseBodyData) SetEnable(v bool) *QueryArmsEnableResponseBodyData {
	s.Enable = &v
	return s
}

func (s *QueryArmsEnableResponseBodyData) Validate() error {
	return dara.Validate(s)
}
