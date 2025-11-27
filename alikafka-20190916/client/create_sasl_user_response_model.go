// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateSaslUserResponseBody) *CreateSaslUserResponse
	GetBody() *CreateSaslUserResponseBody
}

type CreateSaslUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSaslUserResponse) GoString() string {
	return s.String()
}

func (s *CreateSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSaslUserResponse) GetBody() *CreateSaslUserResponseBody {
	return s.Body
}

func (s *CreateSaslUserResponse) SetHeaders(v map[string]*string) *CreateSaslUserResponse {
	s.Headers = v
	return s
}

func (s *CreateSaslUserResponse) SetStatusCode(v int32) *CreateSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSaslUserResponse) SetBody(v *CreateSaslUserResponseBody) *CreateSaslUserResponse {
	s.Body = v
	return s
}

func (s *CreateSaslUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
