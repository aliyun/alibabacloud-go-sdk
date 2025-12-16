// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusForSerialNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientCertificateStatusForSerialNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientCertificateStatusForSerialNumberResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientCertificateStatusForSerialNumberResponseBody) *DescribeClientCertificateStatusForSerialNumberResponse
	GetBody() *DescribeClientCertificateStatusForSerialNumberResponseBody
}

type DescribeClientCertificateStatusForSerialNumberResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateStatusForSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) GetBody() *DescribeClientCertificateStatusForSerialNumberResponseBody {
	return s.Body
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateStatusForSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) SetStatusCode(v int32) *DescribeClientCertificateStatusForSerialNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) SetBody(v *DescribeClientCertificateStatusForSerialNumberResponseBody) *DescribeClientCertificateStatusForSerialNumberResponse {
	s.Body = v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
