// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupedTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupedTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupedTagsResponseBody) *DescribeGroupedTagsResponse
	GetBody() *DescribeGroupedTagsResponseBody
}

type DescribeGroupedTagsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupedTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupedTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupedTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupedTagsResponse) GetBody() *DescribeGroupedTagsResponseBody {
	return s.Body
}

func (s *DescribeGroupedTagsResponse) SetHeaders(v map[string]*string) *DescribeGroupedTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedTagsResponse) SetStatusCode(v int32) *DescribeGroupedTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupedTagsResponse) SetBody(v *DescribeGroupedTagsResponseBody) *DescribeGroupedTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupedTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
