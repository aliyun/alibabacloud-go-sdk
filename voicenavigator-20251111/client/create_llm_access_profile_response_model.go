// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLlmAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLlmAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLlmAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateLlmAccessProfileResponseBody) *CreateLlmAccessProfileResponse
	GetBody() *CreateLlmAccessProfileResponseBody
}

type CreateLlmAccessProfileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLlmAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLlmAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLlmAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLlmAccessProfileResponse) GetBody() *CreateLlmAccessProfileResponseBody {
	return s.Body
}

func (s *CreateLlmAccessProfileResponse) SetHeaders(v map[string]*string) *CreateLlmAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateLlmAccessProfileResponse) SetStatusCode(v int32) *CreateLlmAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLlmAccessProfileResponse) SetBody(v *CreateLlmAccessProfileResponseBody) *CreateLlmAccessProfileResponse {
	s.Body = v
	return s
}

func (s *CreateLlmAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
