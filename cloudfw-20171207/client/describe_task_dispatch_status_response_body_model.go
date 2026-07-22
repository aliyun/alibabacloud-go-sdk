// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTaskDispatchStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeTaskDispatchStatusResponseBody
	GetStatus() *string
}

type DescribeTaskDispatchStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 71209DFE-XXX-XXX-52B4A4E9DA3B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution status of the task. Valid values:
	//
	// - PENDING: pending execution.
	//
	// - RUNNING: running.
	//
	// - SUCCESS: succeeded.
	//
	// - FAILED: failed.
	//
	// - CANCELLED: cancelled.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTaskDispatchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskDispatchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskDispatchStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskDispatchStatusResponseBody) SetRequestId(v string) *DescribeTaskDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskDispatchStatusResponseBody) SetStatus(v string) *DescribeTaskDispatchStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTaskDispatchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
