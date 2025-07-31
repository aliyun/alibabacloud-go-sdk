// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEncryptionKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceEncryptionKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceEncryptionKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceEncryptionKeyResponseBody) *DescribeDBInstanceEncryptionKeyResponse
	GetBody() *DescribeDBInstanceEncryptionKeyResponseBody
}

type DescribeDBInstanceEncryptionKeyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceEncryptionKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceEncryptionKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceEncryptionKeyResponse) GetBody() *DescribeDBInstanceEncryptionKeyResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceEncryptionKeyResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceEncryptionKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponse) SetStatusCode(v int32) *DescribeDBInstanceEncryptionKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponse) SetBody(v *DescribeDBInstanceEncryptionKeyResponseBody) *DescribeDBInstanceEncryptionKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponse) Validate() error {
	return dara.Validate(s)
}
