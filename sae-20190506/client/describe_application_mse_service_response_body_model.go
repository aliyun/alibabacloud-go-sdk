// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMseServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationMseServiceResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationMseServiceResponseBodyData) *DescribeApplicationMseServiceResponseBody
	GetData() *DescribeApplicationMseServiceResponseBodyData
	SetErrorCode(v string) *DescribeApplicationMseServiceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationMseServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationMseServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationMseServiceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationMseServiceResponseBody
	GetTraceId() *string
}

type DescribeApplicationMseServiceResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationMseServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
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

func (s DescribeApplicationMseServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMseServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMseServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationMseServiceResponseBody) GetData() *DescribeApplicationMseServiceResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationMseServiceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationMseServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationMseServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationMseServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationMseServiceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationMseServiceResponseBody) SetCode(v string) *DescribeApplicationMseServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetData(v *DescribeApplicationMseServiceResponseBodyData) *DescribeApplicationMseServiceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetErrorCode(v string) *DescribeApplicationMseServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetMessage(v string) *DescribeApplicationMseServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetRequestId(v string) *DescribeApplicationMseServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetSuccess(v bool) *DescribeApplicationMseServiceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) SetTraceId(v string) *DescribeApplicationMseServiceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationMseServiceResponseBodyData struct {
	MseAppId        *string `json:"MseAppId,omitempty" xml:"MseAppId,omitempty"`
	MseAppName      *string `json:"MseAppName,omitempty" xml:"MseAppName,omitempty"`
	MseAppNameSpace *string `json:"MseAppNameSpace,omitempty" xml:"MseAppNameSpace,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApplicationMseServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMseServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMseServiceResponseBodyData) GetMseAppId() *string {
	return s.MseAppId
}

func (s *DescribeApplicationMseServiceResponseBodyData) GetMseAppName() *string {
	return s.MseAppName
}

func (s *DescribeApplicationMseServiceResponseBodyData) GetMseAppNameSpace() *string {
	return s.MseAppNameSpace
}

func (s *DescribeApplicationMseServiceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationMseServiceResponseBodyData) SetMseAppId(v string) *DescribeApplicationMseServiceResponseBodyData {
	s.MseAppId = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBodyData) SetMseAppName(v string) *DescribeApplicationMseServiceResponseBodyData {
	s.MseAppName = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBodyData) SetMseAppNameSpace(v string) *DescribeApplicationMseServiceResponseBodyData {
	s.MseAppNameSpace = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBodyData) SetStatus(v string) *DescribeApplicationMseServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeApplicationMseServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
