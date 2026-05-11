// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *DescribeExportProgressRequest
	GetExportTaskId() *string
	SetInstanceId(v string) *DescribeExportProgressRequest
	GetInstanceId() *string
}

type DescribeExportProgressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0de8e5ccc2b645039ae6fbda443da73f
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 868eef14-7515-4856-8a50-5c9a22abdbcc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeExportProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressRequest) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *DescribeExportProgressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExportProgressRequest) SetExportTaskId(v string) *DescribeExportProgressRequest {
	s.ExportTaskId = &v
	return s
}

func (s *DescribeExportProgressRequest) SetInstanceId(v string) *DescribeExportProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExportProgressRequest) Validate() error {
	return dara.Validate(s)
}
