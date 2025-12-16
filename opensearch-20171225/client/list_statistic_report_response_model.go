// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStatisticReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStatisticReportResponse
	GetStatusCode() *int32
	SetBody(v *ListStatisticReportResponseBody) *ListStatisticReportResponse
	GetBody() *ListStatisticReportResponseBody
}

type ListStatisticReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStatisticReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStatisticReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticReportResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStatisticReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStatisticReportResponse) GetBody() *ListStatisticReportResponseBody {
	return s.Body
}

func (s *ListStatisticReportResponse) SetHeaders(v map[string]*string) *ListStatisticReportResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticReportResponse) SetStatusCode(v int32) *ListStatisticReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStatisticReportResponse) SetBody(v *ListStatisticReportResponseBody) *ListStatisticReportResponse {
	s.Body = v
	return s
}

func (s *ListStatisticReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
