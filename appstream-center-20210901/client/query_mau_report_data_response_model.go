// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMauReportDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMauReportDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMauReportDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryMauReportDataResponseBody) *QueryMauReportDataResponse
	GetBody() *QueryMauReportDataResponseBody
}

type QueryMauReportDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMauReportDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMauReportDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMauReportDataResponse) GoString() string {
	return s.String()
}

func (s *QueryMauReportDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMauReportDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMauReportDataResponse) GetBody() *QueryMauReportDataResponseBody {
	return s.Body
}

func (s *QueryMauReportDataResponse) SetHeaders(v map[string]*string) *QueryMauReportDataResponse {
	s.Headers = v
	return s
}

func (s *QueryMauReportDataResponse) SetStatusCode(v int32) *QueryMauReportDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMauReportDataResponse) SetBody(v *QueryMauReportDataResponseBody) *QueryMauReportDataResponse {
	s.Body = v
	return s
}

func (s *QueryMauReportDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
