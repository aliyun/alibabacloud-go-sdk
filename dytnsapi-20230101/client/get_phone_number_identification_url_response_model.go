// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneNumberIdentificationUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneNumberIdentificationUrlResponseBody) *GetPhoneNumberIdentificationUrlResponse
	GetBody() *GetPhoneNumberIdentificationUrlResponseBody
}

type GetPhoneNumberIdentificationUrlResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberIdentificationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneNumberIdentificationUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneNumberIdentificationUrlResponse) GetBody() *GetPhoneNumberIdentificationUrlResponseBody {
	return s.Body
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetStatusCode(v int32) *GetPhoneNumberIdentificationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetBody(v *GetPhoneNumberIdentificationUrlResponseBody) *GetPhoneNumberIdentificationUrlResponse {
	s.Body = v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
