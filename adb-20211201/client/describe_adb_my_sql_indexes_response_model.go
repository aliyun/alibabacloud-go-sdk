// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlIndexesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdbMySqlIndexesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdbMySqlIndexesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdbMySqlIndexesResponseBody) *DescribeAdbMySqlIndexesResponse
	GetBody() *DescribeAdbMySqlIndexesResponseBody
}

type DescribeAdbMySqlIndexesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlIndexesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlIndexesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlIndexesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdbMySqlIndexesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdbMySqlIndexesResponse) GetBody() *DescribeAdbMySqlIndexesResponseBody {
	return s.Body
}

func (s *DescribeAdbMySqlIndexesResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlIndexesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlIndexesResponse) SetStatusCode(v int32) *DescribeAdbMySqlIndexesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponse) SetBody(v *DescribeAdbMySqlIndexesResponseBody) *DescribeAdbMySqlIndexesResponse {
	s.Body = v
	return s
}

func (s *DescribeAdbMySqlIndexesResponse) Validate() error {
	return dara.Validate(s)
}
