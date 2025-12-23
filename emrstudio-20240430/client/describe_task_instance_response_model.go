// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskInstanceResponseBody) *DescribeTaskInstanceResponse
	GetBody() *DescribeTaskInstanceResponseBody
}

type DescribeTaskInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskInstanceResponse) GetBody() *DescribeTaskInstanceResponseBody {
	return s.Body
}

func (s *DescribeTaskInstanceResponse) SetHeaders(v map[string]*string) *DescribeTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInstanceResponse) SetStatusCode(v int32) *DescribeTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskInstanceResponse) SetBody(v *DescribeTaskInstanceResponseBody) *DescribeTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
