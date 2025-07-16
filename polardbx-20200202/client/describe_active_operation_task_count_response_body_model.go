// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNeedPop(v int64) *DescribeActiveOperationTaskCountResponseBody
	GetNeedPop() *int64
	SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody
	GetRequestId() *string
	SetTaskCount(v int64) *DescribeActiveOperationTaskCountResponseBody
	GetTaskCount() *int64
}

type DescribeActiveOperationTaskCountResponseBody struct {
	// example:
	//
	// 1
	NeedPop *int64 `json:"NeedPop,omitempty" xml:"NeedPop,omitempty"`
	// example:
	//
	// EC7E27FC-58F8-4722-89BB-D1B6D0971956
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s DescribeActiveOperationTaskCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetNeedPop() *int64 {
	return s.NeedPop
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetNeedPop(v int64) *DescribeActiveOperationTaskCountResponseBody {
	s.NeedPop = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetTaskCount(v int64) *DescribeActiveOperationTaskCountResponseBody {
	s.TaskCount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) Validate() error {
	return dara.Validate(s)
}
