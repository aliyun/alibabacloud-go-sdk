// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackendReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackendReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackendReportResponseBody) *CreateBackendReportResponse
	GetBody() *CreateBackendReportResponseBody
}

type CreateBackendReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackendReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackendReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportResponse) GoString() string {
	return s.String()
}

func (s *CreateBackendReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackendReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackendReportResponse) GetBody() *CreateBackendReportResponseBody {
	return s.Body
}

func (s *CreateBackendReportResponse) SetHeaders(v map[string]*string) *CreateBackendReportResponse {
	s.Headers = v
	return s
}

func (s *CreateBackendReportResponse) SetStatusCode(v int32) *CreateBackendReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackendReportResponse) SetBody(v *CreateBackendReportResponseBody) *CreateBackendReportResponse {
	s.Body = v
	return s
}

func (s *CreateBackendReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
