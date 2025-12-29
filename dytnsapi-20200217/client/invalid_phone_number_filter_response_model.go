// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidPhoneNumberFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvalidPhoneNumberFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvalidPhoneNumberFilterResponse
	GetStatusCode() *int32
	SetBody(v *InvalidPhoneNumberFilterResponseBody) *InvalidPhoneNumberFilterResponse
	GetBody() *InvalidPhoneNumberFilterResponseBody
}

type InvalidPhoneNumberFilterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvalidPhoneNumberFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvalidPhoneNumberFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s InvalidPhoneNumberFilterResponse) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvalidPhoneNumberFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvalidPhoneNumberFilterResponse) GetBody() *InvalidPhoneNumberFilterResponseBody {
	return s.Body
}

func (s *InvalidPhoneNumberFilterResponse) SetHeaders(v map[string]*string) *InvalidPhoneNumberFilterResponse {
	s.Headers = v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetStatusCode(v int32) *InvalidPhoneNumberFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) SetBody(v *InvalidPhoneNumberFilterResponseBody) *InvalidPhoneNumberFilterResponse {
	s.Body = v
	return s
}

func (s *InvalidPhoneNumberFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
