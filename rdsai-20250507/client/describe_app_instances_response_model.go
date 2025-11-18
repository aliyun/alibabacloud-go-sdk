// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppInstancesResponseBody) *DescribeAppInstancesResponse
	GetBody() *DescribeAppInstancesResponseBody
}

type DescribeAppInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppInstancesResponse) GetBody() *DescribeAppInstancesResponseBody {
	return s.Body
}

func (s *DescribeAppInstancesResponse) SetHeaders(v map[string]*string) *DescribeAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppInstancesResponse) SetStatusCode(v int32) *DescribeAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppInstancesResponse) SetBody(v *DescribeAppInstancesResponseBody) *DescribeAppInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
