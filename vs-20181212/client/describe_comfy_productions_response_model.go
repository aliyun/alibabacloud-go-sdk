// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyProductionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyProductionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyProductionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyProductionsResponseBody) *DescribeComfyProductionsResponse
	GetBody() *DescribeComfyProductionsResponseBody
}

type DescribeComfyProductionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyProductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyProductionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyProductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyProductionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyProductionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyProductionsResponse) GetBody() *DescribeComfyProductionsResponseBody {
	return s.Body
}

func (s *DescribeComfyProductionsResponse) SetHeaders(v map[string]*string) *DescribeComfyProductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyProductionsResponse) SetStatusCode(v int32) *DescribeComfyProductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyProductionsResponse) SetBody(v *DescribeComfyProductionsResponseBody) *DescribeComfyProductionsResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyProductionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
