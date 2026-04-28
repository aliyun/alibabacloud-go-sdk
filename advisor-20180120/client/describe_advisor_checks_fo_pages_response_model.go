// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksFoPagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvisorChecksFoPagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvisorChecksFoPagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvisorChecksFoPagesResponseBody) *DescribeAdvisorChecksFoPagesResponse
	GetBody() *DescribeAdvisorChecksFoPagesResponseBody
}

type DescribeAdvisorChecksFoPagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorChecksFoPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvisorChecksFoPagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvisorChecksFoPagesResponse) GetBody() *DescribeAdvisorChecksFoPagesResponseBody {
	return s.Body
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetHeaders(v map[string]*string) *DescribeAdvisorChecksFoPagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetStatusCode(v int32) *DescribeAdvisorChecksFoPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponse) SetBody(v *DescribeAdvisorChecksFoPagesResponseBody) *DescribeAdvisorChecksFoPagesResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
