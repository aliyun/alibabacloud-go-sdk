// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *DescribeTaskDetailResponseBody
	GetBeginTime() *string
	SetRequestId(v string) *DescribeTaskDetailResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeTaskDetailResponseBody
	GetStatus() *string
}

type DescribeTaskDetailResponseBody struct {
	// example:
	//
	// 2026-05-12T07:18:57Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailResponseBody) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskDetailResponseBody) SetBeginTime(v string) *DescribeTaskDetailResponseBody {
	s.BeginTime = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetRequestId(v string) *DescribeTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetStatus(v string) *DescribeTaskDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
