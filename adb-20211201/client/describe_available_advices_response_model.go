// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableAdvicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableAdvicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableAdvicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableAdvicesResponseBody) *DescribeAvailableAdvicesResponse
	GetBody() *DescribeAvailableAdvicesResponseBody
}

type DescribeAvailableAdvicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableAdvicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableAdvicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableAdvicesResponse) GetBody() *DescribeAvailableAdvicesResponseBody {
	return s.Body
}

func (s *DescribeAvailableAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAvailableAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableAdvicesResponse) SetStatusCode(v int32) *DescribeAvailableAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableAdvicesResponse) SetBody(v *DescribeAvailableAdvicesResponseBody) *DescribeAvailableAdvicesResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableAdvicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
