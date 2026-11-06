// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConnectionReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConnectionReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConnectionReportResponse
	GetStatusCode() *int32
	SetBody(v *QueryConnectionReportResponseBody) *QueryConnectionReportResponse
	GetBody() *QueryConnectionReportResponseBody
}

type QueryConnectionReportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConnectionReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConnectionReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConnectionReportResponse) GoString() string {
	return s.String()
}

func (s *QueryConnectionReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConnectionReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConnectionReportResponse) GetBody() *QueryConnectionReportResponseBody {
	return s.Body
}

func (s *QueryConnectionReportResponse) SetHeaders(v map[string]*string) *QueryConnectionReportResponse {
	s.Headers = v
	return s
}

func (s *QueryConnectionReportResponse) SetStatusCode(v int32) *QueryConnectionReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConnectionReportResponse) SetBody(v *QueryConnectionReportResponseBody) *QueryConnectionReportResponse {
	s.Body = v
	return s
}

func (s *QueryConnectionReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
