// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMonitorDataResponseBodyData) *DescribeMonitorDataResponseBody
	GetData() []*DescribeMonitorDataResponseBodyData
	SetMessage(v string) *DescribeMonitorDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitorDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitorDataResponseBody
	GetSuccess() *bool
}

type DescribeMonitorDataResponseBody struct {
	// The response data.
	Data []*DescribeMonitorDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponseBody) GetData() []*DescribeMonitorDataResponseBodyData {
	return s.Data
}

func (s *DescribeMonitorDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorDataResponseBody) SetData(v []*DescribeMonitorDataResponseBodyData) *DescribeMonitorDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetMessage(v string) *DescribeMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetRequestId(v string) *DescribeMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetSuccess(v bool) *DescribeMonitorDataResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorDataResponseBodyData struct {
	// The metric name.
	//
	// example:
	//
	// qps
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array of data points for the metric.
	Value []*DescribeMonitorDataResponseBodyDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeMonitorDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeMonitorDataResponseBodyData) GetValue() []*DescribeMonitorDataResponseBodyDataValue {
	return s.Value
}

func (s *DescribeMonitorDataResponseBodyData) SetName(v string) *DescribeMonitorDataResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeMonitorDataResponseBodyData) SetValue(v []*DescribeMonitorDataResponseBodyDataValue) *DescribeMonitorDataResponseBodyData {
	s.Value = v
	return s
}

func (s *DescribeMonitorDataResponseBodyData) Validate() error {
	if s.Value != nil {
		for _, item := range s.Value {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorDataResponseBodyDataValue struct {
	// The timestamp of the data point.
	//
	// example:
	//
	// 1774972800
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The value of the metric at the specified timestamp.
	//
	// example:
	//
	// 60
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorDataResponseBodyDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataResponseBodyDataValue) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponseBodyDataValue) GetTime() *string {
	return s.Time
}

func (s *DescribeMonitorDataResponseBodyDataValue) GetValue() *float64 {
	return s.Value
}

func (s *DescribeMonitorDataResponseBodyDataValue) SetTime(v string) *DescribeMonitorDataResponseBodyDataValue {
	s.Time = &v
	return s
}

func (s *DescribeMonitorDataResponseBodyDataValue) SetValue(v float64) *DescribeMonitorDataResponseBodyDataValue {
	s.Value = &v
	return s
}

func (s *DescribeMonitorDataResponseBodyDataValue) Validate() error {
	return dara.Validate(s)
}
