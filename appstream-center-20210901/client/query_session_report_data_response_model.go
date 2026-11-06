// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionReportDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySessionReportDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySessionReportDataResponse
	GetStatusCode() *int32
	SetBody(v *QuerySessionReportDataResponseBody) *QuerySessionReportDataResponse
	GetBody() *QuerySessionReportDataResponseBody
}

type QuerySessionReportDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionReportDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySessionReportDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionReportDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionReportDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySessionReportDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySessionReportDataResponse) GetBody() *QuerySessionReportDataResponseBody {
	return s.Body
}

func (s *QuerySessionReportDataResponse) SetHeaders(v map[string]*string) *QuerySessionReportDataResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionReportDataResponse) SetStatusCode(v int32) *QuerySessionReportDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionReportDataResponse) SetBody(v *QuerySessionReportDataResponseBody) *QuerySessionReportDataResponse {
	s.Body = v
	return s
}

func (s *QuerySessionReportDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
