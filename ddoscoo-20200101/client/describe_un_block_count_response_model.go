// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlockCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnBlockCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnBlockCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnBlockCountResponseBody) *DescribeUnBlockCountResponse
	GetBody() *DescribeUnBlockCountResponseBody
}

type DescribeUnBlockCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnBlockCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnBlockCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlockCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnBlockCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnBlockCountResponse) GetBody() *DescribeUnBlockCountResponseBody {
	return s.Body
}

func (s *DescribeUnBlockCountResponse) SetHeaders(v map[string]*string) *DescribeUnBlockCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnBlockCountResponse) SetStatusCode(v int32) *DescribeUnBlockCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnBlockCountResponse) SetBody(v *DescribeUnBlockCountResponseBody) *DescribeUnBlockCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUnBlockCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
