// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeFunctionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeFunctionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeFunctionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeFunctionsResponseBody) *DescribeEdgeFunctionsResponse
	GetBody() *DescribeEdgeFunctionsResponseBody
}

type DescribeEdgeFunctionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeFunctionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeFunctionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeFunctionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeFunctionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeFunctionsResponse) GetBody() *DescribeEdgeFunctionsResponseBody {
	return s.Body
}

func (s *DescribeEdgeFunctionsResponse) SetHeaders(v map[string]*string) *DescribeEdgeFunctionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeFunctionsResponse) SetStatusCode(v int32) *DescribeEdgeFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeFunctionsResponse) SetBody(v *DescribeEdgeFunctionsResponseBody) *DescribeEdgeFunctionsResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeFunctionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
