// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnReportResponseBody) *DescribeCdnReportResponse
	GetBody() *DescribeCdnReportResponseBody
}

type DescribeCdnReportResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnReportResponse) GetBody() *DescribeCdnReportResponseBody {
	return s.Body
}

func (s *DescribeCdnReportResponse) SetHeaders(v map[string]*string) *DescribeCdnReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnReportResponse) SetStatusCode(v int32) *DescribeCdnReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnReportResponse) SetBody(v *DescribeCdnReportResponseBody) *DescribeCdnReportResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
