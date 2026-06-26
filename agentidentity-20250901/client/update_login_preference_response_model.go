// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLoginPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLoginPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLoginPreferenceResponseBody) *UpdateLoginPreferenceResponse
	GetBody() *UpdateLoginPreferenceResponseBody
}

type UpdateLoginPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoginPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoginPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginPreferenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoginPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLoginPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLoginPreferenceResponse) GetBody() *UpdateLoginPreferenceResponseBody {
	return s.Body
}

func (s *UpdateLoginPreferenceResponse) SetHeaders(v map[string]*string) *UpdateLoginPreferenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoginPreferenceResponse) SetStatusCode(v int32) *UpdateLoginPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoginPreferenceResponse) SetBody(v *UpdateLoginPreferenceResponseBody) *UpdateLoginPreferenceResponse {
	s.Body = v
	return s
}

func (s *UpdateLoginPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
