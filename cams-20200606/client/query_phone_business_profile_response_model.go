// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneBusinessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPhoneBusinessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPhoneBusinessProfileResponse
	GetStatusCode() *int32
	SetBody(v *QueryPhoneBusinessProfileResponseBody) *QueryPhoneBusinessProfileResponse
	GetBody() *QueryPhoneBusinessProfileResponseBody
}

type QueryPhoneBusinessProfileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPhoneBusinessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPhoneBusinessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneBusinessProfileResponse) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPhoneBusinessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPhoneBusinessProfileResponse) GetBody() *QueryPhoneBusinessProfileResponseBody {
	return s.Body
}

func (s *QueryPhoneBusinessProfileResponse) SetHeaders(v map[string]*string) *QueryPhoneBusinessProfileResponse {
	s.Headers = v
	return s
}

func (s *QueryPhoneBusinessProfileResponse) SetStatusCode(v int32) *QueryPhoneBusinessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhoneBusinessProfileResponse) SetBody(v *QueryPhoneBusinessProfileResponseBody) *QueryPhoneBusinessProfileResponse {
	s.Body = v
	return s
}

func (s *QueryPhoneBusinessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
