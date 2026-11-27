// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingOutStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStreamingOutStatusResponseBody
	GetRequestId() *string
	SetStatus(v int32) *DescribeStreamingOutStatusResponseBody
	GetStatus() *int32
}

type DescribeStreamingOutStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// For abnormal tasks, RTC proactively ends the tasks and notifies the user server through callbacks. For more information, see [Callbacks](https://help.aliyun.com/document_detail/2709168.html?spm=a2c4g.11186623.help-menu-61399.d_2_0_0_0.49be1761tge4EX#6211464494udk).
	//
	// - 0: The task is started and waiting for users to join.
	//
	// - 1: Streaming in progress.
	//
	// - 2: Streaming exception, live streaming stopped.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeStreamingOutStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingOutStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamingOutStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamingOutStatusResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeStreamingOutStatusResponseBody) SetRequestId(v string) *DescribeStreamingOutStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamingOutStatusResponseBody) SetStatus(v int32) *DescribeStreamingOutStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStreamingOutStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
