// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNeedPop(v int32) *DescribeActiveOperationTaskCountResponseBody
	GetNeedPop() *int32
	SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody
	GetRequestId() *string
	SetTaskCount(v int32) *DescribeActiveOperationTaskCountResponseBody
	GetTaskCount() *int32
}

type DescribeActiveOperationTaskCountResponseBody struct {
	// Indicates whether any O\\&M tasks need pop-up windows to notify users actions. Valid values:
	//
	// 	- **0**: No O\\&M tasks need pop-up windows to notify users actions.
	//
	// 	- **1**: Some O\\&M tasks need pop-up windows to notify users actions.
	//
	// example:
	//
	// 0
	NeedPop *int32 `json:"NeedPop,omitempty" xml:"NeedPop,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 2D9F3768-EDA9-4811-943E-42C8006E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of pending O\\&M tasks.
	//
	// example:
	//
	// 1
	TaskCount *int32 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s DescribeActiveOperationTaskCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetNeedPop() *int32 {
	return s.NeedPop
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTaskCountResponseBody) GetTaskCount() *int32 {
	return s.TaskCount
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetNeedPop(v int32) *DescribeActiveOperationTaskCountResponseBody {
	s.NeedPop = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) SetTaskCount(v int32) *DescribeActiveOperationTaskCountResponseBody {
	s.TaskCount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountResponseBody) Validate() error {
	return dara.Validate(s)
}
