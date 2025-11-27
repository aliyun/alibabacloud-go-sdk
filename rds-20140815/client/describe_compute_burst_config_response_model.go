// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeBurstConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComputeBurstConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComputeBurstConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComputeBurstConfigResponseBody) *DescribeComputeBurstConfigResponse
	GetBody() *DescribeComputeBurstConfigResponseBody
}

type DescribeComputeBurstConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComputeBurstConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComputeBurstConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeBurstConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeBurstConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComputeBurstConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComputeBurstConfigResponse) GetBody() *DescribeComputeBurstConfigResponseBody {
	return s.Body
}

func (s *DescribeComputeBurstConfigResponse) SetHeaders(v map[string]*string) *DescribeComputeBurstConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeBurstConfigResponse) SetStatusCode(v int32) *DescribeComputeBurstConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComputeBurstConfigResponse) SetBody(v *DescribeComputeBurstConfigResponseBody) *DescribeComputeBurstConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeComputeBurstConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
