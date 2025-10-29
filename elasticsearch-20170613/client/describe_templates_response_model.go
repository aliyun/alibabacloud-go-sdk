// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse
	GetBody() *DescribeTemplatesResponseBody
}

type DescribeTemplatesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplatesResponse) GetBody() *DescribeTemplatesResponseBody {
	return s.Body
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesResponse) SetStatusCode(v int32) *DescribeTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplatesResponse) SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
