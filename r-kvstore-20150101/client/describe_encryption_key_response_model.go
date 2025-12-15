// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEncryptionKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEncryptionKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEncryptionKeyResponseBody) *DescribeEncryptionKeyResponse
	GetBody() *DescribeEncryptionKeyResponseBody
}

type DescribeEncryptionKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEncryptionKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEncryptionKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEncryptionKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEncryptionKeyResponse) GetBody() *DescribeEncryptionKeyResponseBody {
	return s.Body
}

func (s *DescribeEncryptionKeyResponse) SetHeaders(v map[string]*string) *DescribeEncryptionKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeEncryptionKeyResponse) SetStatusCode(v int32) *DescribeEncryptionKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEncryptionKeyResponse) SetBody(v *DescribeEncryptionKeyResponseBody) *DescribeEncryptionKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeEncryptionKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
