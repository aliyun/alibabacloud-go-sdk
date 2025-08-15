// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *DescribeExportInfoRequest
	GetExportId() *int64
}

type DescribeExportInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 121****
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeExportInfoRequest) SetExportId(v int64) *DescribeExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
