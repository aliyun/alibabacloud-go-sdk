// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppViewTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppViewTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppViewTemplatesResponseBody) *DescribeAppViewTemplatesResponse
	GetBody() *DescribeAppViewTemplatesResponseBody
}

type DescribeAppViewTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppViewTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppViewTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppViewTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppViewTemplatesResponse) GetBody() *DescribeAppViewTemplatesResponseBody {
	return s.Body
}

func (s *DescribeAppViewTemplatesResponse) SetHeaders(v map[string]*string) *DescribeAppViewTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppViewTemplatesResponse) SetStatusCode(v int32) *DescribeAppViewTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppViewTemplatesResponse) SetBody(v *DescribeAppViewTemplatesResponseBody) *DescribeAppViewTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppViewTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
