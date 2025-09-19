// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRebootStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRebootStatuses(v []*DescribeInstanceRebootStatusResponseBodyRebootStatuses) *DescribeInstanceRebootStatusResponseBody
	GetRebootStatuses() []*DescribeInstanceRebootStatusResponseBodyRebootStatuses
	SetRequestId(v string) *DescribeInstanceRebootStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceRebootStatusResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceRebootStatusResponseBody struct {
	// An array that consists of the status information about the servers that you restart.
	RebootStatuses []*DescribeInstanceRebootStatusResponseBodyRebootStatuses `json:"RebootStatuses,omitempty" xml:"RebootStatuses,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5BD95679-D63A-4151-97D0-188432F4A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceRebootStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRebootStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRebootStatusResponseBody) GetRebootStatuses() []*DescribeInstanceRebootStatusResponseBodyRebootStatuses {
	return s.RebootStatuses
}

func (s *DescribeInstanceRebootStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRebootStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceRebootStatusResponseBody) SetRebootStatuses(v []*DescribeInstanceRebootStatusResponseBodyRebootStatuses) *DescribeInstanceRebootStatusResponseBody {
	s.RebootStatuses = v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBody) SetRequestId(v string) *DescribeInstanceRebootStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBody) SetTotalCount(v int32) *DescribeInstanceRebootStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceRebootStatusResponseBodyRebootStatuses struct {
	// The error code that is returned when the server failed to be restarted. Valid values:
	//
	// 	- **10001**: The restart command failed to be sent.
	//
	// 	- **10002**: The restart operation failed.
	//
	// 	- **10003**: A timeout error occurs.
	//
	// example:
	//
	// 10001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned when the server failed to be restarted.
	//
	// example:
	//
	// push failed
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The status of the server. Valid values:
	//
	// 	- **0**: The server is being restarted.
	//
	// 	- **1**: The server is restarted.
	//
	// 	- **2**: The server failed to be restarted.
	//
	// example:
	//
	// 2
	RebootStatus *int32 `json:"RebootStatus,omitempty" xml:"RebootStatus,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 9b59c2d6-0967-46e3-ad7b-152227c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceRebootStatusResponseBodyRebootStatuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRebootStatusResponseBodyRebootStatuses) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) GetMsg() *string {
	return s.Msg
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) GetRebootStatus() *int32 {
	return s.RebootStatus
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) SetCode(v string) *DescribeInstanceRebootStatusResponseBodyRebootStatuses {
	s.Code = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) SetMsg(v string) *DescribeInstanceRebootStatusResponseBodyRebootStatuses {
	s.Msg = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) SetRebootStatus(v int32) *DescribeInstanceRebootStatusResponseBodyRebootStatuses {
	s.RebootStatus = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) SetUuid(v string) *DescribeInstanceRebootStatusResponseBodyRebootStatuses {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponseBodyRebootStatuses) Validate() error {
	return dara.Validate(s)
}
