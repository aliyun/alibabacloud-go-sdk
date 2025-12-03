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
	SetResourceDirectoryAccountId(v int64) *DescribeExportInfoRequest
	GetResourceDirectoryAccountId() *int64
}

type DescribeExportInfoRequest struct {
	// The ID of the export task.
	//
	// > You can call the [ExportRecord](~~ExportRecord~~) operation to query the IDs of export tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	ExportId                   *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
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

func (s *DescribeExportInfoRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeExportInfoRequest) SetExportId(v int64) *DescribeExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeExportInfoRequest) SetResourceDirectoryAccountId(v int64) *DescribeExportInfoRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
