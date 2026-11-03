// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeApplicationApmServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DowngradeApplicationApmServiceResponseBody
	GetCode() *string
	SetData(v *DowngradeApplicationApmServiceResponseBodyData) *DowngradeApplicationApmServiceResponseBody
	GetData() *DowngradeApplicationApmServiceResponseBodyData
	SetErrorCode(v string) *DowngradeApplicationApmServiceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DowngradeApplicationApmServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DowngradeApplicationApmServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DowngradeApplicationApmServiceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DowngradeApplicationApmServiceResponseBody
	GetTraceId() *string
}

type DowngradeApplicationApmServiceResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned results.
	Data *DowngradeApplicationApmServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// System.Upgrading
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether ARMS Premium Edition monitoring is successfully disabled. Valid values:
	//
	// - **true**: Disabled successfully.
	//
	// - **false**: Failed to disable.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to precisely query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DowngradeApplicationApmServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradeApplicationApmServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeApplicationApmServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DowngradeApplicationApmServiceResponseBody) GetData() *DowngradeApplicationApmServiceResponseBodyData {
	return s.Data
}

func (s *DowngradeApplicationApmServiceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DowngradeApplicationApmServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DowngradeApplicationApmServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradeApplicationApmServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DowngradeApplicationApmServiceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DowngradeApplicationApmServiceResponseBody) SetCode(v string) *DowngradeApplicationApmServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetData(v *DowngradeApplicationApmServiceResponseBodyData) *DowngradeApplicationApmServiceResponseBody {
	s.Data = v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetErrorCode(v string) *DowngradeApplicationApmServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetMessage(v string) *DowngradeApplicationApmServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetRequestId(v string) *DowngradeApplicationApmServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetSuccess(v bool) *DowngradeApplicationApmServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) SetTraceId(v string) *DowngradeApplicationApmServiceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DowngradeApplicationApmServiceResponseBodyData struct {
	// Indicates whether ARMS Premium Edition monitoring is successfully disabled. Valid values:
	//
	// - **true**: Disabled successfully.
	//
	// - **false**: Failed to disable.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DowngradeApplicationApmServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DowngradeApplicationApmServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DowngradeApplicationApmServiceResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *DowngradeApplicationApmServiceResponseBodyData) SetStatus(v bool) *DowngradeApplicationApmServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
