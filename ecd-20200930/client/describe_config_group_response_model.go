// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigGroupResponseBody) *DescribeConfigGroupResponse
	GetBody() *DescribeConfigGroupResponseBody
}

type DescribeConfigGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigGroupResponse) GetBody() *DescribeConfigGroupResponseBody {
	return s.Body
}

func (s *DescribeConfigGroupResponse) SetHeaders(v map[string]*string) *DescribeConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigGroupResponse) SetStatusCode(v int32) *DescribeConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigGroupResponse) SetBody(v *DescribeConfigGroupResponseBody) *DescribeConfigGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
