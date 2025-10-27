// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKmsKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKmsKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKmsKeysResponseBody) *DescribeKmsKeysResponse
	GetBody() *DescribeKmsKeysResponseBody
}

type DescribeKmsKeysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKmsKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKmsKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKmsKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKmsKeysResponse) GetBody() *DescribeKmsKeysResponseBody {
	return s.Body
}

func (s *DescribeKmsKeysResponse) SetHeaders(v map[string]*string) *DescribeKmsKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeKmsKeysResponse) SetStatusCode(v int32) *DescribeKmsKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKmsKeysResponse) SetBody(v *DescribeKmsKeysResponseBody) *DescribeKmsKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeKmsKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
