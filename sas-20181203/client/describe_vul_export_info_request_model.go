// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *DescribeVulExportInfoRequest
	GetExportId() *int64
}

type DescribeVulExportInfoRequest struct {
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14356
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeVulExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulExportInfoRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeVulExportInfoRequest) SetExportId(v int64) *DescribeVulExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeVulExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
