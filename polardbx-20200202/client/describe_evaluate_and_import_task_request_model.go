// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeEvaluateAndImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *DescribeEvaluateAndImportTaskRequest
	GetSlinkTaskId() *string
}

type DescribeEvaluateAndImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s DescribeEvaluateAndImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEvaluateAndImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeEvaluateAndImportTaskRequest) SetRegionId(v string) *DescribeEvaluateAndImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskRequest) SetSlinkTaskId(v string) *DescribeEvaluateAndImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
