// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelAccountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelAccountResponseBody) *DescribeChannelAccountResponse
	GetBody() *DescribeChannelAccountResponseBody
}

type DescribeChannelAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAccountResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelAccountResponse) GetBody() *DescribeChannelAccountResponseBody {
	return s.Body
}

func (s *DescribeChannelAccountResponse) SetHeaders(v map[string]*string) *DescribeChannelAccountResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAccountResponse) SetStatusCode(v int32) *DescribeChannelAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAccountResponse) SetBody(v *DescribeChannelAccountResponseBody) *DescribeChannelAccountResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
