// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDauReportDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDauReportDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDauReportDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryDauReportDataResponseBody) *QueryDauReportDataResponse
	GetBody() *QueryDauReportDataResponseBody
}

type QueryDauReportDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDauReportDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDauReportDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDauReportDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDauReportDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDauReportDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDauReportDataResponse) GetBody() *QueryDauReportDataResponseBody {
	return s.Body
}

func (s *QueryDauReportDataResponse) SetHeaders(v map[string]*string) *QueryDauReportDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDauReportDataResponse) SetStatusCode(v int32) *QueryDauReportDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDauReportDataResponse) SetBody(v *QueryDauReportDataResponseBody) *QueryDauReportDataResponse {
	s.Body = v
	return s
}

func (s *QueryDauReportDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
