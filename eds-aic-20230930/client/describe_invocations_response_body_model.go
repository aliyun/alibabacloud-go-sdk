// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeInvocationsResponseBodyData) *DescribeInvocationsResponseBody
	GetData() []*DescribeInvocationsResponseBodyData
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeInvocationsResponseBody
	GetTotalCount() *string
}

type DescribeInvocationsResponseBody struct {
	// The objects that are returned.
	Data []*DescribeInvocationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 440D7342-5E7C-B2DB-D0B4EAC2BDF1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) GetData() []*DescribeInvocationsResponseBodyData {
	return s.Data
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeInvocationsResponseBody) SetData(v []*DescribeInvocationsResponseBodyData) *DescribeInvocationsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v string) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInvocationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyData struct {
	// The end time of the command execution.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the cloud phone instance on which the command is executed.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// t-15775dc8****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// The execution state of the command.
	//
	// Valid values:
	//
	// 	- Failed: The execution of the command failed.
	//
	// 	- Timeout: The execution of the command timed out.
	//
	// 	- Running: The command is being executed.
	//
	// 	- Success: The execution of the command is successful.
	//
	// 	- Pending: The command is waiting to be executed.
	//
	// example:
	//
	// RUNNING
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The output of the command execution.
	//
	// example:
	//
	// success
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The start time of the command execution.
	//
	// example:
	//
	// 2022-08-11 17:45:03
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInvocationsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationsResponseBodyData) GetInvocationId() *string {
	return s.InvocationId
}

func (s *DescribeInvocationsResponseBodyData) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationsResponseBodyData) SetFinishTime(v string) *DescribeInvocationsResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInstanceId(v string) *DescribeInvocationsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInvocationId(v string) *DescribeInvocationsResponseBodyData {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyData {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetOutput(v string) *DescribeInvocationsResponseBodyData {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) SetStartTime(v string) *DescribeInvocationsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
