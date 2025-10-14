// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStructureImportTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeStructureImportTaskInfoRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *DescribeStructureImportTaskInfoRequest
	GetSlinkTaskId() *string
}

type DescribeStructureImportTaskInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s DescribeStructureImportTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStructureImportTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeStructureImportTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStructureImportTaskInfoRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeStructureImportTaskInfoRequest) SetRegionId(v string) *DescribeStructureImportTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStructureImportTaskInfoRequest) SetSlinkTaskId(v string) *DescribeStructureImportTaskInfoRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeStructureImportTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
