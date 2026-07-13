// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAtiCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAtiCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAtiCertificateResponseBody) *DescribeAtiCertificateResponse
	GetBody() *DescribeAtiCertificateResponseBody
}

type DescribeAtiCertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAtiCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAtiCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeAtiCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAtiCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAtiCertificateResponse) GetBody() *DescribeAtiCertificateResponseBody {
	return s.Body
}

func (s *DescribeAtiCertificateResponse) SetHeaders(v map[string]*string) *DescribeAtiCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeAtiCertificateResponse) SetStatusCode(v int32) *DescribeAtiCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAtiCertificateResponse) SetBody(v *DescribeAtiCertificateResponseBody) *DescribeAtiCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeAtiCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
