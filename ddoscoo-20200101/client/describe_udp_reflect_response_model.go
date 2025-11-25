// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdpReflectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUdpReflectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUdpReflectResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUdpReflectResponseBody) *DescribeUdpReflectResponse
	GetBody() *DescribeUdpReflectResponseBody
}

type DescribeUdpReflectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUdpReflectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUdpReflectResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdpReflectResponse) GoString() string {
	return s.String()
}

func (s *DescribeUdpReflectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUdpReflectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUdpReflectResponse) GetBody() *DescribeUdpReflectResponseBody {
	return s.Body
}

func (s *DescribeUdpReflectResponse) SetHeaders(v map[string]*string) *DescribeUdpReflectResponse {
	s.Headers = v
	return s
}

func (s *DescribeUdpReflectResponse) SetStatusCode(v int32) *DescribeUdpReflectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUdpReflectResponse) SetBody(v *DescribeUdpReflectResponseBody) *DescribeUdpReflectResponse {
	s.Body = v
	return s
}

func (s *DescribeUdpReflectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
