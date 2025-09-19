// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *DescribeWarningExportInfoRequest
	GetExportId() *int64
}

type DescribeWarningExportInfoRequest struct {
	// The ID of the export task.
	//
	// >  You can can call the [ExportWarning](~~ExportWarning~~) operation to query the IDs of export tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14356
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeWarningExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarningExportInfoRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeWarningExportInfoRequest) SetExportId(v int64) *DescribeWarningExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeWarningExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
