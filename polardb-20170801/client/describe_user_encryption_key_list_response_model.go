// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserEncryptionKeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserEncryptionKeyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse
	GetBody() *DescribeUserEncryptionKeyListResponseBody
}

type DescribeUserEncryptionKeyListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserEncryptionKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserEncryptionKeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserEncryptionKeyListResponse) GetBody() *DescribeUserEncryptionKeyListResponseBody {
	return s.Body
}

func (s *DescribeUserEncryptionKeyListResponse) SetHeaders(v map[string]*string) *DescribeUserEncryptionKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetStatusCode(v int32) *DescribeUserEncryptionKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetBody(v *DescribeUserEncryptionKeyListResponseBody) *DescribeUserEncryptionKeyListResponse {
	s.Body = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) Validate() error {
	return dara.Validate(s)
}
