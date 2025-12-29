// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsForMongoDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcsForMongoDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcsForMongoDBResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcsForMongoDBResponseBody) *DescribeVpcsForMongoDBResponse
	GetBody() *DescribeVpcsForMongoDBResponseBody
}

type DescribeVpcsForMongoDBResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcsForMongoDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcsForMongoDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsForMongoDBResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcsForMongoDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcsForMongoDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcsForMongoDBResponse) GetBody() *DescribeVpcsForMongoDBResponseBody {
	return s.Body
}

func (s *DescribeVpcsForMongoDBResponse) SetHeaders(v map[string]*string) *DescribeVpcsForMongoDBResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcsForMongoDBResponse) SetStatusCode(v int32) *DescribeVpcsForMongoDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponse) SetBody(v *DescribeVpcsForMongoDBResponseBody) *DescribeVpcsForMongoDBResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcsForMongoDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
