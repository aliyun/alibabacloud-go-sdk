// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPrivacyCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPrivacyCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPrivacyCallResponse
	GetStatusCode() *int32
	SetBody(v *StartPrivacyCallResponseBody) *StartPrivacyCallResponse
	GetBody() *StartPrivacyCallResponseBody
}

type StartPrivacyCallResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPrivacyCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPrivacyCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPrivacyCallResponse) GoString() string {
	return s.String()
}

func (s *StartPrivacyCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPrivacyCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPrivacyCallResponse) GetBody() *StartPrivacyCallResponseBody {
	return s.Body
}

func (s *StartPrivacyCallResponse) SetHeaders(v map[string]*string) *StartPrivacyCallResponse {
	s.Headers = v
	return s
}

func (s *StartPrivacyCallResponse) SetStatusCode(v int32) *StartPrivacyCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPrivacyCallResponse) SetBody(v *StartPrivacyCallResponseBody) *StartPrivacyCallResponse {
	s.Body = v
	return s
}

func (s *StartPrivacyCallResponse) Validate() error {
	return dara.Validate(s)
}
