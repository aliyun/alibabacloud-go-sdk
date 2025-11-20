// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaInstancesResponseBody) *DescribeHanaInstancesResponse
	GetBody() *DescribeHanaInstancesResponseBody
}

type DescribeHanaInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaInstancesResponse) GetBody() *DescribeHanaInstancesResponseBody {
	return s.Body
}

func (s *DescribeHanaInstancesResponse) SetHeaders(v map[string]*string) *DescribeHanaInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaInstancesResponse) SetStatusCode(v int32) *DescribeHanaInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaInstancesResponse) SetBody(v *DescribeHanaInstancesResponseBody) *DescribeHanaInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
