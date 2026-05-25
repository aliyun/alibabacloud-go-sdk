// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberIdentificationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneNumberIdentificationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneNumberIdentificationResultResponseBody) *GetPhoneNumberIdentificationResultResponse
	GetBody() *GetPhoneNumberIdentificationResultResponseBody
}

type GetPhoneNumberIdentificationResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberIdentificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberIdentificationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneNumberIdentificationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneNumberIdentificationResultResponse) GetBody() *GetPhoneNumberIdentificationResultResponseBody {
	return s.Body
}

func (s *GetPhoneNumberIdentificationResultResponse) SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponse) SetStatusCode(v int32) *GetPhoneNumberIdentificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponse) SetBody(v *GetPhoneNumberIdentificationResultResponseBody) *GetPhoneNumberIdentificationResultResponse {
	s.Body = v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
