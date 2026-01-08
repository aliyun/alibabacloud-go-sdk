// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappPhoneNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryChatappPhoneNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryChatappPhoneNumbersResponse
	GetStatusCode() *int32
	SetBody(v *QueryChatappPhoneNumbersResponseBody) *QueryChatappPhoneNumbersResponse
	GetBody() *QueryChatappPhoneNumbersResponseBody
}

type QueryChatappPhoneNumbersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChatappPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChatappPhoneNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryChatappPhoneNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryChatappPhoneNumbersResponse) GetBody() *QueryChatappPhoneNumbersResponseBody {
	return s.Body
}

func (s *QueryChatappPhoneNumbersResponse) SetHeaders(v map[string]*string) *QueryChatappPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *QueryChatappPhoneNumbersResponse) SetStatusCode(v int32) *QueryChatappPhoneNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponse) SetBody(v *QueryChatappPhoneNumbersResponseBody) *QueryChatappPhoneNumbersResponse {
	s.Body = v
	return s
}

func (s *QueryChatappPhoneNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
