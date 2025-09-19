// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLevelCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventLevelCountResponseBody
	GetCode() *string
	SetEventLevels(v *DescribeEventLevelCountResponseBodyEventLevels) *DescribeEventLevelCountResponseBody
	GetEventLevels() *DescribeEventLevelCountResponseBodyEventLevels
	SetMessage(v string) *DescribeEventLevelCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventLevelCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEventLevelCountResponseBody
	GetSuccess() *bool
}

type DescribeEventLevelCountResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The statistics of alerts by risk level.
	EventLevels *DescribeEventLevelCountResponseBodyEventLevels `json:"EventLevels,omitempty" xml:"EventLevels,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CE8CA5EA-24EF-5D41-B735-53ACE7XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventLevelCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLevelCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventLevelCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventLevelCountResponseBody) GetEventLevels() *DescribeEventLevelCountResponseBodyEventLevels {
	return s.EventLevels
}

func (s *DescribeEventLevelCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventLevelCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventLevelCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventLevelCountResponseBody) SetCode(v string) *DescribeEventLevelCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventLevelCountResponseBody) SetEventLevels(v *DescribeEventLevelCountResponseBodyEventLevels) *DescribeEventLevelCountResponseBody {
	s.EventLevels = v
	return s
}

func (s *DescribeEventLevelCountResponseBody) SetMessage(v string) *DescribeEventLevelCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventLevelCountResponseBody) SetRequestId(v string) *DescribeEventLevelCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventLevelCountResponseBody) SetSuccess(v bool) *DescribeEventLevelCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventLevelCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventLevelCountResponseBodyEventLevels struct {
	// The number of alerts whose Emergency level is Reminder.
	//
	// example:
	//
	// 2
	Remind *int32 `json:"Remind,omitempty" xml:"Remind,omitempty"`
	// The number of alerts whose Emergency level is Urgent.
	//
	// example:
	//
	// 0
	Serious *int32 `json:"Serious,omitempty" xml:"Serious,omitempty"`
	// The number of alerts whose Emergency level is Suspicious.
	//
	// example:
	//
	// 1
	Suspicious *int32 `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
}

func (s DescribeEventLevelCountResponseBodyEventLevels) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLevelCountResponseBodyEventLevels) GoString() string {
	return s.String()
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) GetRemind() *int32 {
	return s.Remind
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) GetSerious() *int32 {
	return s.Serious
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) GetSuspicious() *int32 {
	return s.Suspicious
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) SetRemind(v int32) *DescribeEventLevelCountResponseBodyEventLevels {
	s.Remind = &v
	return s
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) SetSerious(v int32) *DescribeEventLevelCountResponseBodyEventLevels {
	s.Serious = &v
	return s
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) SetSuspicious(v int32) *DescribeEventLevelCountResponseBodyEventLevels {
	s.Suspicious = &v
	return s
}

func (s *DescribeEventLevelCountResponseBodyEventLevels) Validate() error {
	return dara.Validate(s)
}
