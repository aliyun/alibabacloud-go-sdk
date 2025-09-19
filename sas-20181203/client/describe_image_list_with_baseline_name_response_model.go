// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListWithBaselineNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageListWithBaselineNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageListWithBaselineNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageListWithBaselineNameResponseBody) *DescribeImageListWithBaselineNameResponse
	GetBody() *DescribeImageListWithBaselineNameResponseBody
}

type DescribeImageListWithBaselineNameResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageListWithBaselineNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageListWithBaselineNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListWithBaselineNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageListWithBaselineNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageListWithBaselineNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageListWithBaselineNameResponse) GetBody() *DescribeImageListWithBaselineNameResponseBody {
	return s.Body
}

func (s *DescribeImageListWithBaselineNameResponse) SetHeaders(v map[string]*string) *DescribeImageListWithBaselineNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageListWithBaselineNameResponse) SetStatusCode(v int32) *DescribeImageListWithBaselineNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageListWithBaselineNameResponse) SetBody(v *DescribeImageListWithBaselineNameResponseBody) *DescribeImageListWithBaselineNameResponse {
	s.Body = v
	return s
}

func (s *DescribeImageListWithBaselineNameResponse) Validate() error {
	return dara.Validate(s)
}
