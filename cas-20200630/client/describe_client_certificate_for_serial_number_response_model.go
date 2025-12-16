// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateForSerialNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientCertificateForSerialNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientCertificateForSerialNumberResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientCertificateForSerialNumberResponseBody) *DescribeClientCertificateForSerialNumberResponse
	GetBody() *DescribeClientCertificateForSerialNumberResponseBody
}

type DescribeClientCertificateForSerialNumberResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateForSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientCertificateForSerialNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientCertificateForSerialNumberResponse) GetBody() *DescribeClientCertificateForSerialNumberResponseBody {
	return s.Body
}

func (s *DescribeClientCertificateForSerialNumberResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateForSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponse) SetStatusCode(v int32) *DescribeClientCertificateForSerialNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponse) SetBody(v *DescribeClientCertificateForSerialNumberResponseBody) *DescribeClientCertificateForSerialNumberResponse {
	s.Body = v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
