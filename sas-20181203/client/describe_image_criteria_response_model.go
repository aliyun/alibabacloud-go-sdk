// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageCriteriaResponseBody) *DescribeImageCriteriaResponse
	GetBody() *DescribeImageCriteriaResponseBody
}

type DescribeImageCriteriaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageCriteriaResponse) GetBody() *DescribeImageCriteriaResponseBody {
	return s.Body
}

func (s *DescribeImageCriteriaResponse) SetHeaders(v map[string]*string) *DescribeImageCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageCriteriaResponse) SetStatusCode(v int32) *DescribeImageCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageCriteriaResponse) SetBody(v *DescribeImageCriteriaResponseBody) *DescribeImageCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeImageCriteriaResponse) Validate() error {
	return dara.Validate(s)
}
