// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeParamTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeParamTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeParamTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeParamTagsResponseBody) *DescribeNodeParamTagsResponse
	GetBody() *DescribeNodeParamTagsResponseBody
}

type DescribeNodeParamTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeParamTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeParamTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeParamTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeParamTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeParamTagsResponse) GetBody() *DescribeNodeParamTagsResponseBody {
	return s.Body
}

func (s *DescribeNodeParamTagsResponse) SetHeaders(v map[string]*string) *DescribeNodeParamTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetStatusCode(v int32) *DescribeNodeParamTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetBody(v *DescribeNodeParamTagsResponseBody) *DescribeNodeParamTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeParamTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
