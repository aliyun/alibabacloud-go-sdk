// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScaleOutMigrateTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgress(v int32) *DescribeScaleOutMigrateTaskListResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeScaleOutMigrateTaskListResponseBody
	GetRequestId() *string
}

type DescribeScaleOutMigrateTaskListResponseBody struct {
	// example:
	//
	// 32
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) SetProgress(v int32) *DescribeScaleOutMigrateTaskListResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) SetRequestId(v string) *DescribeScaleOutMigrateTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}
