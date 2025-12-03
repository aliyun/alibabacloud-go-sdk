// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeletedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeletedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeletedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeletedInstancesResponseBody) *DescribeDeletedInstancesResponse
	GetBody() *DescribeDeletedInstancesResponseBody
}

type DescribeDeletedInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeletedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeletedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeletedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeletedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeletedInstancesResponse) GetBody() *DescribeDeletedInstancesResponseBody {
	return s.Body
}

func (s *DescribeDeletedInstancesResponse) SetHeaders(v map[string]*string) *DescribeDeletedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeletedInstancesResponse) SetStatusCode(v int32) *DescribeDeletedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeletedInstancesResponse) SetBody(v *DescribeDeletedInstancesResponseBody) *DescribeDeletedInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDeletedInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
