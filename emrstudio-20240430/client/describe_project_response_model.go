// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse
	GetBody() *DescribeProjectResponseBody
}

type DescribeProjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectResponse) GetBody() *DescribeProjectResponseBody {
	return s.Body
}

func (s *DescribeProjectResponse) SetHeaders(v map[string]*string) *DescribeProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectResponse) SetStatusCode(v int32) *DescribeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectResponse) SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
