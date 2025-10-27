// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComputeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComputeResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComputeResourceResponseBody) *DescribeComputeResourceResponse
	GetBody() *DescribeComputeResourceResponseBody
}

type DescribeComputeResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComputeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComputeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComputeResourceResponse) GetBody() *DescribeComputeResourceResponseBody {
	return s.Body
}

func (s *DescribeComputeResourceResponse) SetHeaders(v map[string]*string) *DescribeComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeResourceResponse) SetStatusCode(v int32) *DescribeComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComputeResourceResponse) SetBody(v *DescribeComputeResourceResponseBody) *DescribeComputeResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeComputeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
