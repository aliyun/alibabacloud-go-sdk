// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConsumersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConsumersResponseBody) *DescribeConsumersResponse
	GetBody() *DescribeConsumersResponseBody
}

type DescribeConsumersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumersResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConsumersResponse) GetBody() *DescribeConsumersResponseBody {
	return s.Body
}

func (s *DescribeConsumersResponse) SetHeaders(v map[string]*string) *DescribeConsumersResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumersResponse) SetStatusCode(v int32) *DescribeConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConsumersResponse) SetBody(v *DescribeConsumersResponseBody) *DescribeConsumersResponse {
	s.Body = v
	return s
}

func (s *DescribeConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
