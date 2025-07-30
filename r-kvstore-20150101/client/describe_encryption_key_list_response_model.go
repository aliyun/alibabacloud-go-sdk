// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEncryptionKeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEncryptionKeyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEncryptionKeyListResponseBody) *DescribeEncryptionKeyListResponse
	GetBody() *DescribeEncryptionKeyListResponseBody
}

type DescribeEncryptionKeyListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEncryptionKeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEncryptionKeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEncryptionKeyListResponse) GetBody() *DescribeEncryptionKeyListResponseBody {
	return s.Body
}

func (s *DescribeEncryptionKeyListResponse) SetHeaders(v map[string]*string) *DescribeEncryptionKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEncryptionKeyListResponse) SetStatusCode(v int32) *DescribeEncryptionKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEncryptionKeyListResponse) SetBody(v *DescribeEncryptionKeyListResponseBody) *DescribeEncryptionKeyListResponse {
	s.Body = v
	return s
}

func (s *DescribeEncryptionKeyListResponse) Validate() error {
	return dara.Validate(s)
}
