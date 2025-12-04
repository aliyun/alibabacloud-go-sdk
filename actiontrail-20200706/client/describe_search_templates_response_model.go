// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSearchTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSearchTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSearchTemplatesResponseBody) *DescribeSearchTemplatesResponse
	GetBody() *DescribeSearchTemplatesResponseBody
}

type DescribeSearchTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSearchTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSearchTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSearchTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSearchTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSearchTemplatesResponse) GetBody() *DescribeSearchTemplatesResponseBody {
	return s.Body
}

func (s *DescribeSearchTemplatesResponse) SetHeaders(v map[string]*string) *DescribeSearchTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSearchTemplatesResponse) SetStatusCode(v int32) *DescribeSearchTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSearchTemplatesResponse) SetBody(v *DescribeSearchTemplatesResponseBody) *DescribeSearchTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeSearchTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
