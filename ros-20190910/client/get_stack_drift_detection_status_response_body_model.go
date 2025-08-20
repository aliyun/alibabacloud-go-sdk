// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDriftDetectionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriftDetectionId(v string) *GetStackDriftDetectionStatusResponseBody
	GetDriftDetectionId() *string
	SetDriftDetectionStatus(v string) *GetStackDriftDetectionStatusResponseBody
	GetDriftDetectionStatus() *string
	SetDriftDetectionStatusReason(v string) *GetStackDriftDetectionStatusResponseBody
	GetDriftDetectionStatusReason() *string
	SetDriftDetectionTime(v string) *GetStackDriftDetectionStatusResponseBody
	GetDriftDetectionTime() *string
	SetDriftedStackResourceCount(v int32) *GetStackDriftDetectionStatusResponseBody
	GetDriftedStackResourceCount() *int32
	SetRequestId(v string) *GetStackDriftDetectionStatusResponseBody
	GetRequestId() *string
	SetStackDriftStatus(v string) *GetStackDriftDetectionStatusResponseBody
	GetStackDriftStatus() *string
	SetStackId(v string) *GetStackDriftDetectionStatusResponseBody
	GetStackId() *string
}

type GetStackDriftDetectionStatusResponseBody struct {
	// The ID of the drift detection operation.
	//
	// example:
	//
	// a7044f0d-6f2e-4128-a307-4524ef88****
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The drift detection status. Valid values:
	//
	// 	- DETECTION_COMPLETE: The drift detection operation has been completed for all resources that support drift detection in the stack.
	//
	// 	- DETECTION_FAILED: The stack drift detection operation has failed for at least one resource in the stack.
	//
	// 	- DETECTION_IN_PROGRESS: The stack drift detection operation is in progress.
	//
	// example:
	//
	// DETECTION_COMPLETE
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The reason why the stack drift detection operation has its current status.
	//
	// example:
	//
	// Detect stack drift successfully
	DriftDetectionStatusReason *string `json:"DriftDetectionStatusReason,omitempty" xml:"DriftDetectionStatusReason,omitempty"`
	// The time when the stack drift detection operation was initiated.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The total number of stack resources that have drifted.
	//
	// example:
	//
	// 1
	DriftedStackResourceCount *int32 `json:"DriftedStackResourceCount,omitempty" xml:"DriftedStackResourceCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The drift status of the stack. Valid values:
	//
	// 	- DRIFTED: The actual configuration of the stack differs, or has drifted, from its expected template configuration. A stack is considered to have drifted if one or more of its resources have drifted.
	//
	// 	- NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	//
	// 	- IN_SYNC: The current configuration of each supported resource matches its expected template configuration. A stack with no resources that support drift detection also has a status of IN_SYNC.
	//
	// example:
	//
	// DRIFTED
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackDriftDetectionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponseBody) GetDriftDetectionId() *string {
	return s.DriftDetectionId
}

func (s *GetStackDriftDetectionStatusResponseBody) GetDriftDetectionStatus() *string {
	return s.DriftDetectionStatus
}

func (s *GetStackDriftDetectionStatusResponseBody) GetDriftDetectionStatusReason() *string {
	return s.DriftDetectionStatusReason
}

func (s *GetStackDriftDetectionStatusResponseBody) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackDriftDetectionStatusResponseBody) GetDriftedStackResourceCount() *int32 {
	return s.DriftedStackResourceCount
}

func (s *GetStackDriftDetectionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackDriftDetectionStatusResponseBody) GetStackDriftStatus() *string {
	return s.StackDriftStatus
}

func (s *GetStackDriftDetectionStatusResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatusReason(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatusReason = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionTime(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftedStackResourceCount(v int32) *GetStackDriftDetectionStatusResponseBody {
	s.DriftedStackResourceCount = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetRequestId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackDriftStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
