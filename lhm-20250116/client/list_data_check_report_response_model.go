// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckReportResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckReportResponseBody) *ListDataCheckReportResponse
	GetBody() *ListDataCheckReportResponseBody
}

type ListDataCheckReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckReportResponse) GetBody() *ListDataCheckReportResponseBody {
	return s.Body
}

func (s *ListDataCheckReportResponse) SetHeaders(v map[string]*string) *ListDataCheckReportResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckReportResponse) SetStatusCode(v int32) *ListDataCheckReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckReportResponse) SetBody(v *ListDataCheckReportResponseBody) *ListDataCheckReportResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
