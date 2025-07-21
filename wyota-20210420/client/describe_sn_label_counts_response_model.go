// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnLabelCountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnLabelCountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnLabelCountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnLabelCountsResponseBody) *DescribeSnLabelCountsResponse
	GetBody() *DescribeSnLabelCountsResponseBody
}

type DescribeSnLabelCountsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnLabelCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnLabelCountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnLabelCountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnLabelCountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnLabelCountsResponse) GetBody() *DescribeSnLabelCountsResponseBody {
	return s.Body
}

func (s *DescribeSnLabelCountsResponse) SetHeaders(v map[string]*string) *DescribeSnLabelCountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnLabelCountsResponse) SetStatusCode(v int32) *DescribeSnLabelCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnLabelCountsResponse) SetBody(v *DescribeSnLabelCountsResponseBody) *DescribeSnLabelCountsResponse {
	s.Body = v
	return s
}

func (s *DescribeSnLabelCountsResponse) Validate() error {
	return dara.Validate(s)
}
