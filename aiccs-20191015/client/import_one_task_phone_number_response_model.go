// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneTaskPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportOneTaskPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportOneTaskPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *ImportOneTaskPhoneNumberResponseBody) *ImportOneTaskPhoneNumberResponse
	GetBody() *ImportOneTaskPhoneNumberResponseBody
}

type ImportOneTaskPhoneNumberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportOneTaskPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportOneTaskPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportOneTaskPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ImportOneTaskPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportOneTaskPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportOneTaskPhoneNumberResponse) GetBody() *ImportOneTaskPhoneNumberResponseBody {
	return s.Body
}

func (s *ImportOneTaskPhoneNumberResponse) SetHeaders(v map[string]*string) *ImportOneTaskPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ImportOneTaskPhoneNumberResponse) SetStatusCode(v int32) *ImportOneTaskPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportOneTaskPhoneNumberResponse) SetBody(v *ImportOneTaskPhoneNumberResponseBody) *ImportOneTaskPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *ImportOneTaskPhoneNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
