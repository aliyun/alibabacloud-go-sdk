// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneTwiceTelVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPhoneTwiceTelVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPhoneTwiceTelVerifyResponse
	GetStatusCode() *int32
	SetBody(v *QueryPhoneTwiceTelVerifyResponseBody) *QueryPhoneTwiceTelVerifyResponse
	GetBody() *QueryPhoneTwiceTelVerifyResponseBody
}

type QueryPhoneTwiceTelVerifyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneTwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneTwiceTelVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneTwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneTwiceTelVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPhoneTwiceTelVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPhoneTwiceTelVerifyResponse) GetBody() *QueryPhoneTwiceTelVerifyResponseBody {
	return s.Body
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetHeaders(v map[string]*string) *QueryPhoneTwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetStatusCode(v int32) *QueryPhoneTwiceTelVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponse) SetBody(v *QueryPhoneTwiceTelVerifyResponseBody) *QueryPhoneTwiceTelVerifyResponse {
	s.Body = v
	return s
}

func (s *QueryPhoneTwiceTelVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
