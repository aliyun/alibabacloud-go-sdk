// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetReportResponse(v *GetReportResp) *GetReportResponseBody
	GetGetReportResponse() *GetReportResp
}

type GetReportResponseBody struct {
	// The details for obtaining the migration report.
	GetReportResponse *GetReportResp `json:"GetReportResponse,omitempty" xml:"GetReportResponse,omitempty"`
}

func (s GetReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportResponseBody) GetGetReportResponse() *GetReportResp {
	return s.GetReportResponse
}

func (s *GetReportResponseBody) SetGetReportResponse(v *GetReportResp) *GetReportResponseBody {
	s.GetReportResponse = v
	return s
}

func (s *GetReportResponseBody) Validate() error {
	if s.GetReportResponse != nil {
		if err := s.GetReportResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}
