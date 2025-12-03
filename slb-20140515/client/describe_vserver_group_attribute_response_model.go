// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVServerGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVServerGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVServerGroupAttributeResponseBody) *DescribeVServerGroupAttributeResponse
	GetBody() *DescribeVServerGroupAttributeResponseBody
}

type DescribeVServerGroupAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVServerGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVServerGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVServerGroupAttributeResponse) GetBody() *DescribeVServerGroupAttributeResponseBody {
	return s.Body
}

func (s *DescribeVServerGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeVServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVServerGroupAttributeResponse) SetStatusCode(v int32) *DescribeVServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVServerGroupAttributeResponse) SetBody(v *DescribeVServerGroupAttributeResponseBody) *DescribeVServerGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeVServerGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
