// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResultCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResultCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResultCountResponseBody) *DescribeResultCountResponse
	GetBody() *DescribeResultCountResponseBody
}

type DescribeResultCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResultCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResultCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeResultCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResultCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResultCountResponse) GetBody() *DescribeResultCountResponseBody {
	return s.Body
}

func (s *DescribeResultCountResponse) SetHeaders(v map[string]*string) *DescribeResultCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeResultCountResponse) SetStatusCode(v int32) *DescribeResultCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResultCountResponse) SetBody(v *DescribeResultCountResponseBody) *DescribeResultCountResponse {
	s.Body = v
	return s
}

func (s *DescribeResultCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
