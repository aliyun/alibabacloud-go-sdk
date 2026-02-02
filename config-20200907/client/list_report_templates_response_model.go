// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReportTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReportTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListReportTemplatesResponseBody) *ListReportTemplatesResponse
	GetBody() *ListReportTemplatesResponseBody
}

type ListReportTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReportTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReportTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReportTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListReportTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReportTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReportTemplatesResponse) GetBody() *ListReportTemplatesResponseBody {
	return s.Body
}

func (s *ListReportTemplatesResponse) SetHeaders(v map[string]*string) *ListReportTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListReportTemplatesResponse) SetStatusCode(v int32) *ListReportTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReportTemplatesResponse) SetBody(v *ListReportTemplatesResponseBody) *ListReportTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListReportTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
