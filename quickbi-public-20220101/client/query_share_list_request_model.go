// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShareListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *QueryShareListRequest
	GetReportId() *string
}

type QueryShareListRequest struct {
	// The ID of the work. The works include data portal, dashboard, spreadsheet, self-service data retrieval, ad-hoc analysis, data entry, and data screen.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s QueryShareListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryShareListRequest) GoString() string {
	return s.String()
}

func (s *QueryShareListRequest) GetReportId() *string {
	return s.ReportId
}

func (s *QueryShareListRequest) SetReportId(v string) *QueryShareListRequest {
	s.ReportId = &v
	return s
}

func (s *QueryShareListRequest) Validate() error {
	return dara.Validate(s)
}
