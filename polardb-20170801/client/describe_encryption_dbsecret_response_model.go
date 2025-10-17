// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEncryptionDBSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEncryptionDBSecretResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEncryptionDBSecretResponseBody) *DescribeEncryptionDBSecretResponse
	GetBody() *DescribeEncryptionDBSecretResponseBody
}

type DescribeEncryptionDBSecretResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEncryptionDBSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEncryptionDBSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBSecretResponse) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEncryptionDBSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEncryptionDBSecretResponse) GetBody() *DescribeEncryptionDBSecretResponseBody {
	return s.Body
}

func (s *DescribeEncryptionDBSecretResponse) SetHeaders(v map[string]*string) *DescribeEncryptionDBSecretResponse {
	s.Headers = v
	return s
}

func (s *DescribeEncryptionDBSecretResponse) SetStatusCode(v int32) *DescribeEncryptionDBSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponse) SetBody(v *DescribeEncryptionDBSecretResponseBody) *DescribeEncryptionDBSecretResponse {
	s.Body = v
	return s
}

func (s *DescribeEncryptionDBSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
