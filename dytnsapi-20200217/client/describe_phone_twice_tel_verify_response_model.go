// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneTwiceTelVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneTwiceTelVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneTwiceTelVerifyResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneTwiceTelVerifyResponseBody) *DescribePhoneTwiceTelVerifyResponse
	GetBody() *DescribePhoneTwiceTelVerifyResponseBody
}

type DescribePhoneTwiceTelVerifyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneTwiceTelVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneTwiceTelVerifyResponse) GetBody() *DescribePhoneTwiceTelVerifyResponseBody {
	return s.Body
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetHeaders(v map[string]*string) *DescribePhoneTwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetStatusCode(v int32) *DescribePhoneTwiceTelVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) SetBody(v *DescribePhoneTwiceTelVerifyResponseBody) *DescribePhoneTwiceTelVerifyResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
