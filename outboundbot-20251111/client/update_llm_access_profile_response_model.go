// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLlmAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLlmAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLlmAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLlmAccessProfileResponseBody) *UpdateLlmAccessProfileResponse
	GetBody() *UpdateLlmAccessProfileResponseBody
}

type UpdateLlmAccessProfileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLlmAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLlmAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLlmAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLlmAccessProfileResponse) GetBody() *UpdateLlmAccessProfileResponseBody {
	return s.Body
}

func (s *UpdateLlmAccessProfileResponse) SetHeaders(v map[string]*string) *UpdateLlmAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateLlmAccessProfileResponse) SetStatusCode(v int32) *UpdateLlmAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLlmAccessProfileResponse) SetBody(v *UpdateLlmAccessProfileResponseBody) *UpdateLlmAccessProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateLlmAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
