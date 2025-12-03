// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsByApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrafficControlsByApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrafficControlsByApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrafficControlsByApiResponseBody) *DescribeTrafficControlsByApiResponse
	GetBody() *DescribeTrafficControlsByApiResponseBody
}

type DescribeTrafficControlsByApiResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficControlsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficControlsByApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrafficControlsByApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrafficControlsByApiResponse) GetBody() *DescribeTrafficControlsByApiResponseBody {
	return s.Body
}

func (s *DescribeTrafficControlsByApiResponse) SetHeaders(v map[string]*string) *DescribeTrafficControlsByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficControlsByApiResponse) SetStatusCode(v int32) *DescribeTrafficControlsByApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponse) SetBody(v *DescribeTrafficControlsByApiResponseBody) *DescribeTrafficControlsByApiResponse {
	s.Body = v
	return s
}

func (s *DescribeTrafficControlsByApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
