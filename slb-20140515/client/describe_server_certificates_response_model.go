// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerCertificatesResponseBody) *DescribeServerCertificatesResponse
	GetBody() *DescribeServerCertificatesResponseBody
}

type DescribeServerCertificatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerCertificatesResponse) GetBody() *DescribeServerCertificatesResponseBody {
	return s.Body
}

func (s *DescribeServerCertificatesResponse) SetHeaders(v map[string]*string) *DescribeServerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerCertificatesResponse) SetStatusCode(v int32) *DescribeServerCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerCertificatesResponse) SetBody(v *DescribeServerCertificatesResponseBody) *DescribeServerCertificatesResponse {
	s.Body = v
	return s
}

func (s *DescribeServerCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
