// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *DeleteNisInspectionReportRequest
	GetInspectionReportId() *string
}

type DeleteNisInspectionReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// nir-ffd1af****196d0
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
}

func (s DeleteNisInspectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionReportRequest) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionReportRequest) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DeleteNisInspectionReportRequest) SetInspectionReportId(v string) *DeleteNisInspectionReportRequest {
	s.InspectionReportId = &v
	return s
}

func (s *DeleteNisInspectionReportRequest) Validate() error {
	return dara.Validate(s)
}
