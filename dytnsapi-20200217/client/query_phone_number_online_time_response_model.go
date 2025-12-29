// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNumberOnlineTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPhoneNumberOnlineTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPhoneNumberOnlineTimeResponse
	GetStatusCode() *int32
	SetBody(v *QueryPhoneNumberOnlineTimeResponseBody) *QueryPhoneNumberOnlineTimeResponse
	GetBody() *QueryPhoneNumberOnlineTimeResponseBody
}

type QueryPhoneNumberOnlineTimeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneNumberOnlineTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPhoneNumberOnlineTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPhoneNumberOnlineTimeResponse) GetBody() *QueryPhoneNumberOnlineTimeResponseBody {
	return s.Body
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetHeaders(v map[string]*string) *QueryPhoneNumberOnlineTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetStatusCode(v int32) *QueryPhoneNumberOnlineTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponse) SetBody(v *QueryPhoneNumberOnlineTimeResponseBody) *QueryPhoneNumberOnlineTimeResponse {
	s.Body = v
	return s
}

func (s *QueryPhoneNumberOnlineTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
