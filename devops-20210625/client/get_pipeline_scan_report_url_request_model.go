// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineScanReportUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportPath(v string) *GetPipelineScanReportUrlRequest
	GetReportPath() *string
}

type GetPipelineScanReportUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /test/test/test.html
	ReportPath *string `json:"reportPath,omitempty" xml:"reportPath,omitempty"`
}

func (s GetPipelineScanReportUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineScanReportUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlRequest) GetReportPath() *string {
	return s.ReportPath
}

func (s *GetPipelineScanReportUrlRequest) SetReportPath(v string) *GetPipelineScanReportUrlRequest {
	s.ReportPath = &v
	return s
}

func (s *GetPipelineScanReportUrlRequest) Validate() error {
	return dara.Validate(s)
}
