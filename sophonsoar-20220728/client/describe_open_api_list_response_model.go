// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenApiListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenApiListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenApiListResponseBody) *DescribeOpenApiListResponse
	GetBody() *DescribeOpenApiListResponseBody
}

type DescribeOpenApiListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenApiListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenApiListResponse) GetBody() *DescribeOpenApiListResponseBody {
	return s.Body
}

func (s *DescribeOpenApiListResponse) SetHeaders(v map[string]*string) *DescribeOpenApiListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiListResponse) SetStatusCode(v int32) *DescribeOpenApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiListResponse) SetBody(v *DescribeOpenApiListResponseBody) *DescribeOpenApiListResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenApiListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
