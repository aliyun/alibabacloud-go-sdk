// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportId(v string) *StartNisInspectionTaskResponseBody
	GetInspectionReportId() *string
	SetRequestId(v string) *StartNisInspectionTaskResponseBody
	GetRequestId() *string
}

type StartNisInspectionTaskResponseBody struct {
	// example:
	//
	// nir-9923e262c106461d86ad
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartNisInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartNisInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartNisInspectionTaskResponseBody) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *StartNisInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartNisInspectionTaskResponseBody) SetInspectionReportId(v string) *StartNisInspectionTaskResponseBody {
	s.InspectionReportId = &v
	return s
}

func (s *StartNisInspectionTaskResponseBody) SetRequestId(v string) *StartNisInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartNisInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
