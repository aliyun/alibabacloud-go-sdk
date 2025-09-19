// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHcExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *DescribeHcExportInfoRequest
	GetExportId() *int64
}

type DescribeHcExportInfoRequest struct {
	// The ID of the export task.
	//
	// >  You can call the [ExportWarning](~~ExportWarning~~) operation to query the IDs of export tasks.
	//
	// example:
	//
	// 443285
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeHcExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHcExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHcExportInfoRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeHcExportInfoRequest) SetExportId(v int64) *DescribeHcExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeHcExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
