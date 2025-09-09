// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableListByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableListByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableListByTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableListByTypeResponseBody) *DescribeTableListByTypeResponse
	GetBody() *DescribeTableListByTypeResponseBody
}

type DescribeTableListByTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableListByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableListByTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableListByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableListByTypeResponse) GetBody() *DescribeTableListByTypeResponseBody {
	return s.Body
}

func (s *DescribeTableListByTypeResponse) SetHeaders(v map[string]*string) *DescribeTableListByTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableListByTypeResponse) SetStatusCode(v int32) *DescribeTableListByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableListByTypeResponse) SetBody(v *DescribeTableListByTypeResponseBody) *DescribeTableListByTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeTableListByTypeResponse) Validate() error {
	return dara.Validate(s)
}
