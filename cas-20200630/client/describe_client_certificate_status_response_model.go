// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientCertificateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientCertificateStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientCertificateStatusResponseBody) *DescribeClientCertificateStatusResponse
	GetBody() *DescribeClientCertificateStatusResponseBody
}

type DescribeClientCertificateStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientCertificateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientCertificateStatusResponse) GetBody() *DescribeClientCertificateStatusResponseBody {
	return s.Body
}

func (s *DescribeClientCertificateStatusResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateStatusResponse) SetStatusCode(v int32) *DescribeClientCertificateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateStatusResponse) SetBody(v *DescribeClientCertificateStatusResponseBody) *DescribeClientCertificateStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeClientCertificateStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
