// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenApiInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenApiInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenApiInfoResponseBody) *DescribeOpenApiInfoResponse
	GetBody() *DescribeOpenApiInfoResponseBody
}

type DescribeOpenApiInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenApiInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenApiInfoResponse) GetBody() *DescribeOpenApiInfoResponseBody {
	return s.Body
}

func (s *DescribeOpenApiInfoResponse) SetHeaders(v map[string]*string) *DescribeOpenApiInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiInfoResponse) SetStatusCode(v int32) *DescribeOpenApiInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiInfoResponse) SetBody(v *DescribeOpenApiInfoResponseBody) *DescribeOpenApiInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenApiInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
