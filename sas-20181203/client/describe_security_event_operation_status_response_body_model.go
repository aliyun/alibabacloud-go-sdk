// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityEventOperationStatusResponseBody
	GetRequestId() *string
	SetSecurityEventOperationStatusResponse(v *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) *DescribeSecurityEventOperationStatusResponseBody
	GetSecurityEventOperationStatusResponse() *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse
}

type DescribeSecurityEventOperationStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1683940A-E4AE-4473-8C40-F4075434B76B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the task that handles the alert events.
	SecurityEventOperationStatusResponse *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse `json:"SecurityEventOperationStatusResponse,omitempty" xml:"SecurityEventOperationStatusResponse,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventOperationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventOperationStatusResponseBody) GetSecurityEventOperationStatusResponse() *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse {
	return s.SecurityEventOperationStatusResponse
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetSecurityEventOperationStatusResponse(v *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) *DescribeSecurityEventOperationStatusResponseBody {
	s.SecurityEventOperationStatusResponse = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse struct {
	// An array consisting of the status of the alert events handled by the task.
	SecurityEventOperationStatuses []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses `json:"SecurityEventOperationStatuses,omitempty" xml:"SecurityEventOperationStatuses,omitempty" type:"Repeated"`
	// The status of the task that handles the alert events. Valid values:
	//
	// 	- **Processing**: The task is running.
	//
	// 	- **Success**: The task is successful.
	//
	// 	- **Failure**: The task failed.
	//
	// 	- **Pending**: The task is pending.
	//
	// example:
	//
	// Success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) GetSecurityEventOperationStatuses() []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	return s.SecurityEventOperationStatuses
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) SetSecurityEventOperationStatuses(v []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse {
	s.SecurityEventOperationStatuses = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) SetTaskStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses struct {
	// The code that indicates the handling result of the alert event.
	//
	// example:
	//
	// ignore.Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the alert event.
	//
	// example:
	//
	// 12321
	SecurityEventId *string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The handling status of the alert event. Valid values:
	//
	// 	- **Processing**: The alert event is being handled.
	//
	// 	- **Success**: The alert event is handled.
	//
	// 	- **Failed**: The alert event failed to be handled.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) GetSecurityEventId() *string {
	return s.SecurityEventId
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) GetStatus() *string {
	return s.Status
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetErrorCode(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetSecurityEventId(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.Status = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) Validate() error {
	return dara.Validate(s)
}
