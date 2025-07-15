// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisplayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisplayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisplayConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisplayConfigResponseBody) *DescribeDisplayConfigResponse
	GetBody() *DescribeDisplayConfigResponseBody
}

type DescribeDisplayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisplayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisplayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisplayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisplayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisplayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisplayConfigResponse) GetBody() *DescribeDisplayConfigResponseBody {
	return s.Body
}

func (s *DescribeDisplayConfigResponse) SetHeaders(v map[string]*string) *DescribeDisplayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisplayConfigResponse) SetStatusCode(v int32) *DescribeDisplayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisplayConfigResponse) SetBody(v *DescribeDisplayConfigResponseBody) *DescribeDisplayConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDisplayConfigResponse) Validate() error {
	return dara.Validate(s)
}
