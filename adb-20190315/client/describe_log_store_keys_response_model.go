// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogStoreKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogStoreKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogStoreKeysResponseBody) *DescribeLogStoreKeysResponse
	GetBody() *DescribeLogStoreKeysResponseBody
}

type DescribeLogStoreKeysResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogStoreKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogStoreKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogStoreKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogStoreKeysResponse) GetBody() *DescribeLogStoreKeysResponseBody {
	return s.Body
}

func (s *DescribeLogStoreKeysResponse) SetHeaders(v map[string]*string) *DescribeLogStoreKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreKeysResponse) SetStatusCode(v int32) *DescribeLogStoreKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreKeysResponse) SetBody(v *DescribeLogStoreKeysResponseBody) *DescribeLogStoreKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeLogStoreKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
