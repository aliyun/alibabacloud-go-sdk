// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceInspectionReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceInspectionReportResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceInspectionReportResponseBody) *GetInstanceInspectionReportResponse
	GetBody() *GetInstanceInspectionReportResponseBody
}

type GetInstanceInspectionReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceInspectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceInspectionReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceInspectionReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceInspectionReportResponse) GetBody() *GetInstanceInspectionReportResponseBody {
	return s.Body
}

func (s *GetInstanceInspectionReportResponse) SetHeaders(v map[string]*string) *GetInstanceInspectionReportResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceInspectionReportResponse) SetStatusCode(v int32) *GetInstanceInspectionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceInspectionReportResponse) SetBody(v *GetInstanceInspectionReportResponseBody) *GetInstanceInspectionReportResponse {
	s.Body = v
	return s
}

func (s *GetInstanceInspectionReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
