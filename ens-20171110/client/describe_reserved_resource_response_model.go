// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReservedResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReservedResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReservedResourceResponseBody) *DescribeReservedResourceResponse
	GetBody() *DescribeReservedResourceResponseBody
}

type DescribeReservedResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReservedResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReservedResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeReservedResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReservedResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReservedResourceResponse) GetBody() *DescribeReservedResourceResponseBody {
	return s.Body
}

func (s *DescribeReservedResourceResponse) SetHeaders(v map[string]*string) *DescribeReservedResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeReservedResourceResponse) SetStatusCode(v int32) *DescribeReservedResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReservedResourceResponse) SetBody(v *DescribeReservedResourceResponseBody) *DescribeReservedResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeReservedResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
