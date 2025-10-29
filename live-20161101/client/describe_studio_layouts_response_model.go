// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStudioLayoutsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStudioLayoutsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStudioLayoutsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStudioLayoutsResponseBody) *DescribeStudioLayoutsResponse
	GetBody() *DescribeStudioLayoutsResponseBody
}

type DescribeStudioLayoutsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStudioLayoutsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStudioLayoutsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStudioLayoutsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStudioLayoutsResponse) GetBody() *DescribeStudioLayoutsResponseBody {
	return s.Body
}

func (s *DescribeStudioLayoutsResponse) SetHeaders(v map[string]*string) *DescribeStudioLayoutsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStudioLayoutsResponse) SetStatusCode(v int32) *DescribeStudioLayoutsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStudioLayoutsResponse) SetBody(v *DescribeStudioLayoutsResponseBody) *DescribeStudioLayoutsResponse {
	s.Body = v
	return s
}

func (s *DescribeStudioLayoutsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
