// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterOutputListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExporterOutputListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExporterOutputListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExporterOutputListResponseBody) *DescribeExporterOutputListResponse
	GetBody() *DescribeExporterOutputListResponseBody
}

type DescribeExporterOutputListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExporterOutputListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExporterOutputListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExporterOutputListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExporterOutputListResponse) GetBody() *DescribeExporterOutputListResponseBody {
	return s.Body
}

func (s *DescribeExporterOutputListResponse) SetHeaders(v map[string]*string) *DescribeExporterOutputListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExporterOutputListResponse) SetStatusCode(v int32) *DescribeExporterOutputListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExporterOutputListResponse) SetBody(v *DescribeExporterOutputListResponseBody) *DescribeExporterOutputListResponse {
	s.Body = v
	return s
}

func (s *DescribeExporterOutputListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
