// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePersonalICodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePersonalICodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePersonalICodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePersonalICodeResponseBody) *UpdatePersonalICodeResponse
	GetBody() *UpdatePersonalICodeResponseBody
}

type UpdatePersonalICodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePersonalICodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePersonalICodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePersonalICodeResponse) GoString() string {
	return s.String()
}

func (s *UpdatePersonalICodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePersonalICodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePersonalICodeResponse) GetBody() *UpdatePersonalICodeResponseBody {
	return s.Body
}

func (s *UpdatePersonalICodeResponse) SetHeaders(v map[string]*string) *UpdatePersonalICodeResponse {
	s.Headers = v
	return s
}

func (s *UpdatePersonalICodeResponse) SetStatusCode(v int32) *UpdatePersonalICodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePersonalICodeResponse) SetBody(v *UpdatePersonalICodeResponseBody) *UpdatePersonalICodeResponse {
	s.Body = v
	return s
}

func (s *UpdatePersonalICodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
