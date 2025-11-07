// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNisInspectionReportResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNisInspectionReportResponseBody
	GetRequestId() *string
}

type DeleteNisInspectionReportResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNisInspectionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionReportResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNisInspectionReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNisInspectionReportResponseBody) SetData(v bool) *DeleteNisInspectionReportResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNisInspectionReportResponseBody) SetRequestId(v string) *DeleteNisInspectionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNisInspectionReportResponseBody) Validate() error {
	return dara.Validate(s)
}
