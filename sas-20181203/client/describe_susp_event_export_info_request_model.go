// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int32) *DescribeSuspEventExportInfoRequest
	GetExportId() *int32
}

type DescribeSuspEventExportInfoRequest struct {
	// The ID of the export task.
	//
	// > You can call the [ExportSuspEvents](~~ExportSuspEvents~~) operation to query the ID.
	//
	// example:
	//
	// 123
	ExportId *int32 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeSuspEventExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventExportInfoRequest) GetExportId() *int32 {
	return s.ExportId
}

func (s *DescribeSuspEventExportInfoRequest) SetExportId(v int32) *DescribeSuspEventExportInfoRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeSuspEventExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
