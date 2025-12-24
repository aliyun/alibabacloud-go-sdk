// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsJsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentsJsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentsJsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentsJsResponseBody) *DescribeComponentsJsResponse
	GetBody() *DescribeComponentsJsResponseBody
}

type DescribeComponentsJsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentsJsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentsJsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsJsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentsJsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentsJsResponse) GetBody() *DescribeComponentsJsResponseBody {
	return s.Body
}

func (s *DescribeComponentsJsResponse) SetHeaders(v map[string]*string) *DescribeComponentsJsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentsJsResponse) SetStatusCode(v int32) *DescribeComponentsJsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentsJsResponse) SetBody(v *DescribeComponentsJsResponseBody) *DescribeComponentsJsResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentsJsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
