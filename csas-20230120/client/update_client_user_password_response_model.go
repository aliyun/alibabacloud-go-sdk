// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientUserPasswordResponseBody) *UpdateClientUserPasswordResponse
	GetBody() *UpdateClientUserPasswordResponseBody
}

type UpdateClientUserPasswordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientUserPasswordResponse) GetBody() *UpdateClientUserPasswordResponseBody {
	return s.Body
}

func (s *UpdateClientUserPasswordResponse) SetHeaders(v map[string]*string) *UpdateClientUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserPasswordResponse) SetStatusCode(v int32) *UpdateClientUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserPasswordResponse) SetBody(v *UpdateClientUserPasswordResponseBody) *UpdateClientUserPasswordResponse {
	s.Body = v
	return s
}

func (s *UpdateClientUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
