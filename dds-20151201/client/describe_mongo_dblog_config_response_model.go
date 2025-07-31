// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMongoDBLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMongoDBLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMongoDBLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMongoDBLogConfigResponseBody) *DescribeMongoDBLogConfigResponse
	GetBody() *DescribeMongoDBLogConfigResponseBody
}

type DescribeMongoDBLogConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMongoDBLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMongoDBLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMongoDBLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMongoDBLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMongoDBLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMongoDBLogConfigResponse) GetBody() *DescribeMongoDBLogConfigResponseBody {
	return s.Body
}

func (s *DescribeMongoDBLogConfigResponse) SetHeaders(v map[string]*string) *DescribeMongoDBLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMongoDBLogConfigResponse) SetStatusCode(v int32) *DescribeMongoDBLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMongoDBLogConfigResponse) SetBody(v *DescribeMongoDBLogConfigResponseBody) *DescribeMongoDBLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeMongoDBLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
