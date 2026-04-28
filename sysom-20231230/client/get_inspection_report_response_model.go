// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInspectionReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInspectionReportResponse
	GetStatusCode() *int32
	SetBody(v *GetInspectionReportResponseBody) *GetInspectionReportResponse
	GetBody() *GetInspectionReportResponseBody
}

type GetInspectionReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInspectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInspectionReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportResponse) GoString() string {
	return s.String()
}

func (s *GetInspectionReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInspectionReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInspectionReportResponse) GetBody() *GetInspectionReportResponseBody {
	return s.Body
}

func (s *GetInspectionReportResponse) SetHeaders(v map[string]*string) *GetInspectionReportResponse {
	s.Headers = v
	return s
}

func (s *GetInspectionReportResponse) SetStatusCode(v int32) *GetInspectionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInspectionReportResponse) SetBody(v *GetInspectionReportResponseBody) *GetInspectionReportResponse {
	s.Body = v
	return s
}

func (s *GetInspectionReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
