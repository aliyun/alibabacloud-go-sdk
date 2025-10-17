// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEncryptionKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterEncryptionKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterEncryptionKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterEncryptionKeyResponseBody) *DescribeDBClusterEncryptionKeyResponse
	GetBody() *DescribeDBClusterEncryptionKeyResponseBody
}

type DescribeDBClusterEncryptionKeyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterEncryptionKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterEncryptionKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEncryptionKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEncryptionKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterEncryptionKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterEncryptionKeyResponse) GetBody() *DescribeDBClusterEncryptionKeyResponseBody {
	return s.Body
}

func (s *DescribeDBClusterEncryptionKeyResponse) SetHeaders(v map[string]*string) *DescribeDBClusterEncryptionKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponse) SetStatusCode(v int32) *DescribeDBClusterEncryptionKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponse) SetBody(v *DescribeDBClusterEncryptionKeyResponseBody) *DescribeDBClusterEncryptionKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
