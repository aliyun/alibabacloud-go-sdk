// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserUsageReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserUsageReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserUsageReportResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserUsageReportResponseBody) *QueryUserUsageReportResponse
	GetBody() *QueryUserUsageReportResponseBody
}

type QueryUserUsageReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserUsageReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserUsageReportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserUsageReportResponse) GoString() string {
	return s.String()
}

func (s *QueryUserUsageReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserUsageReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserUsageReportResponse) GetBody() *QueryUserUsageReportResponseBody {
	return s.Body
}

func (s *QueryUserUsageReportResponse) SetHeaders(v map[string]*string) *QueryUserUsageReportResponse {
	s.Headers = v
	return s
}

func (s *QueryUserUsageReportResponse) SetStatusCode(v int32) *QueryUserUsageReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserUsageReportResponse) SetBody(v *QueryUserUsageReportResponseBody) *QueryUserUsageReportResponse {
	s.Body = v
	return s
}

func (s *QueryUserUsageReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
