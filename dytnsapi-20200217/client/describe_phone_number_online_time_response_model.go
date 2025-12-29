// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOnlineTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberOnlineTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberOnlineTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberOnlineTimeResponseBody) *DescribePhoneNumberOnlineTimeResponse
	GetBody() *DescribePhoneNumberOnlineTimeResponseBody
}

type DescribePhoneNumberOnlineTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberOnlineTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberOnlineTimeResponse) GetBody() *DescribePhoneNumberOnlineTimeResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetStatusCode(v int32) *DescribePhoneNumberOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) SetBody(v *DescribePhoneNumberOnlineTimeResponseBody) *DescribePhoneNumberOnlineTimeResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
