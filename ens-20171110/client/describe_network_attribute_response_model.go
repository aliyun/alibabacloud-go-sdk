// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkAttributeResponseBody) *DescribeNetworkAttributeResponse
	GetBody() *DescribeNetworkAttributeResponseBody
}

type DescribeNetworkAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkAttributeResponse) GetBody() *DescribeNetworkAttributeResponseBody {
	return s.Body
}

func (s *DescribeNetworkAttributeResponse) SetHeaders(v map[string]*string) *DescribeNetworkAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkAttributeResponse) SetStatusCode(v int32) *DescribeNetworkAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkAttributeResponse) SetBody(v *DescribeNetworkAttributeResponseBody) *DescribeNetworkAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
