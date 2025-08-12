// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLogCount(v []*DescribeAlertLogCountResponseBodyAlertLogCount) *DescribeAlertLogCountResponseBody
	GetAlertLogCount() []*DescribeAlertLogCountResponseBodyAlertLogCount
	SetCode(v string) *DescribeAlertLogCountResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAlertLogCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertLogCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertLogCountResponseBody
	GetSuccess() *bool
}

type DescribeAlertLogCountResponseBody struct {
	// The statistics of alert logs.
	AlertLogCount []*DescribeAlertLogCountResponseBodyAlertLogCount `json:"AlertLogCount,omitempty" xml:"AlertLogCount,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1C4A3709-BF52-42EE-87B5-7435F0929585
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertLogCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBody) GetAlertLogCount() []*DescribeAlertLogCountResponseBodyAlertLogCount {
	return s.AlertLogCount
}

func (s *DescribeAlertLogCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertLogCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertLogCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertLogCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertLogCountResponseBody) SetAlertLogCount(v []*DescribeAlertLogCountResponseBodyAlertLogCount) *DescribeAlertLogCountResponseBody {
	s.AlertLogCount = v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetCode(v string) *DescribeAlertLogCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetMessage(v string) *DescribeAlertLogCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetRequestId(v string) *DescribeAlertLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetSuccess(v bool) *DescribeAlertLogCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogCountResponseBodyAlertLogCount struct {
	// The number of alert logs.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The details about alert logs.
	Logs []*DescribeAlertLogCountResponseBodyAlertLogCountLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s DescribeAlertLogCountResponseBodyAlertLogCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogCountResponseBodyAlertLogCount) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) GetLogs() []*DescribeAlertLogCountResponseBodyAlertLogCountLogs {
	return s.Logs
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) SetCount(v int32) *DescribeAlertLogCountResponseBodyAlertLogCount {
	s.Count = &v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) SetLogs(v []*DescribeAlertLogCountResponseBodyAlertLogCountLogs) *DescribeAlertLogCountResponseBodyAlertLogCount {
	s.Logs = v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertLogCountResponseBodyAlertLogCountLogs struct {
	// The name of the dimension field based on which alert logs are aggregated.
	//
	// example:
	//
	// product
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the dimension field based on which alert logs are aggregated.
	//
	// example:
	//
	// ECS
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAlertLogCountResponseBodyAlertLogCountLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogCountResponseBodyAlertLogCountLogs) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) GetName() *string {
	return s.Name
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) SetName(v string) *DescribeAlertLogCountResponseBodyAlertLogCountLogs {
	s.Name = &v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) SetValue(v string) *DescribeAlertLogCountResponseBodyAlertLogCountLogs {
	s.Value = &v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) Validate() error {
	return dara.Validate(s)
}
