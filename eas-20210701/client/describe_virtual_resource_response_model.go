// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVirtualResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVirtualResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVirtualResourceResponseBody) *DescribeVirtualResourceResponse
	GetBody() *DescribeVirtualResourceResponseBody
}

type DescribeVirtualResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVirtualResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVirtualResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVirtualResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVirtualResourceResponse) GetBody() *DescribeVirtualResourceResponseBody {
	return s.Body
}

func (s *DescribeVirtualResourceResponse) SetHeaders(v map[string]*string) *DescribeVirtualResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualResourceResponse) SetStatusCode(v int32) *DescribeVirtualResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualResourceResponse) SetBody(v *DescribeVirtualResourceResponseBody) *DescribeVirtualResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeVirtualResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
