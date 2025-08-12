// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricRuleCountResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMetricRuleCountResponseBody
	GetMessage() *string
	SetMetricRuleCount(v *DescribeMetricRuleCountResponseBodyMetricRuleCount) *DescribeMetricRuleCountResponseBody
	GetMetricRuleCount() *DescribeMetricRuleCountResponseBodyMetricRuleCount
	SetRequestId(v string) *DescribeMetricRuleCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleCountResponseBody
	GetSuccess() *bool
}

type DescribeMetricRuleCountResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of alert rules in each state.
	MetricRuleCount *DescribeMetricRuleCountResponseBodyMetricRuleCount `json:"MetricRuleCount,omitempty" xml:"MetricRuleCount,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FF38D33A-67C1-40EB-AB65-FAEE51EDB644
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

func (s DescribeMetricRuleCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricRuleCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleCountResponseBody) GetMetricRuleCount() *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	return s.MetricRuleCount
}

func (s *DescribeMetricRuleCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleCountResponseBody) SetCode(v string) *DescribeMetricRuleCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetMessage(v string) *DescribeMetricRuleCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetMetricRuleCount(v *DescribeMetricRuleCountResponseBodyMetricRuleCount) *DescribeMetricRuleCountResponseBody {
	s.MetricRuleCount = v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetRequestId(v string) *DescribeMetricRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetSuccess(v bool) *DescribeMetricRuleCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleCountResponseBodyMetricRuleCount struct {
	// The number of alert rules with active alerts.
	//
	// example:
	//
	// 5
	Alarm *int32 `json:"Alarm,omitempty" xml:"Alarm,omitempty"`
	// The number of disabled alert rules.
	//
	// example:
	//
	// 0
	Disable *int32 `json:"Disable,omitempty" xml:"Disable,omitempty"`
	// The number of alert rules without data.
	//
	// example:
	//
	// 0
	Nodata *int32 `json:"Nodata,omitempty" xml:"Nodata,omitempty"`
	// The number of alert rules without active alerts.
	//
	// example:
	//
	// 40
	Ok *int32 `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// The total number of alert rules.
	//
	// example:
	//
	// 45
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleCountResponseBodyMetricRuleCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleCountResponseBodyMetricRuleCount) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) GetAlarm() *int32 {
	return s.Alarm
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) GetDisable() *int32 {
	return s.Disable
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) GetNodata() *int32 {
	return s.Nodata
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) GetOk() *int32 {
	return s.Ok
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetAlarm(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Alarm = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetDisable(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Disable = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetNodata(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Nodata = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetOk(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Ok = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetTotal(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) Validate() error {
	return dara.Validate(s)
}
