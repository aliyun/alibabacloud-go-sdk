// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStorageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceStorageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceStorageConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceStorageConfigResponseBody) *DescribeInstanceStorageConfigResponse
	GetBody() *DescribeInstanceStorageConfigResponseBody
}

type DescribeInstanceStorageConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceStorageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceStorageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStorageConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceStorageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceStorageConfigResponse) GetBody() *DescribeInstanceStorageConfigResponseBody {
	return s.Body
}

func (s *DescribeInstanceStorageConfigResponse) SetHeaders(v map[string]*string) *DescribeInstanceStorageConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStorageConfigResponse) SetStatusCode(v int32) *DescribeInstanceStorageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponse) SetBody(v *DescribeInstanceStorageConfigResponseBody) *DescribeInstanceStorageConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceStorageConfigResponse) Validate() error {
	return dara.Validate(s)
}
