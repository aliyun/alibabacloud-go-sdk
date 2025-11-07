// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNisInspectionReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNisInspectionReportResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNisInspectionReportResponseBody) *DeleteNisInspectionReportResponse
	GetBody() *DeleteNisInspectionReportResponseBody
}

type DeleteNisInspectionReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNisInspectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNisInspectionReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionReportResponse) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNisInspectionReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNisInspectionReportResponse) GetBody() *DeleteNisInspectionReportResponseBody {
	return s.Body
}

func (s *DeleteNisInspectionReportResponse) SetHeaders(v map[string]*string) *DeleteNisInspectionReportResponse {
	s.Headers = v
	return s
}

func (s *DeleteNisInspectionReportResponse) SetStatusCode(v int32) *DeleteNisInspectionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNisInspectionReportResponse) SetBody(v *DeleteNisInspectionReportResponseBody) *DeleteNisInspectionReportResponse {
	s.Body = v
	return s
}

func (s *DeleteNisInspectionReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
