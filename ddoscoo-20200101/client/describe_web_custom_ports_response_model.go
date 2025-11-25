// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCustomPortsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCustomPortsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCustomPortsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebCustomPortsResponseBody) *DescribeWebCustomPortsResponse
	GetBody() *DescribeWebCustomPortsResponseBody
}

type DescribeWebCustomPortsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebCustomPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCustomPortsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomPortsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCustomPortsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCustomPortsResponse) GetBody() *DescribeWebCustomPortsResponseBody {
	return s.Body
}

func (s *DescribeWebCustomPortsResponse) SetHeaders(v map[string]*string) *DescribeWebCustomPortsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCustomPortsResponse) SetStatusCode(v int32) *DescribeWebCustomPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCustomPortsResponse) SetBody(v *DescribeWebCustomPortsResponseBody) *DescribeWebCustomPortsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebCustomPortsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
