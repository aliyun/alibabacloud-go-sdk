// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *SetUserPasswordResponseBody) *SetUserPasswordResponse
	GetBody() *SetUserPasswordResponseBody
}

type SetUserPasswordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *SetUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserPasswordResponse) GetBody() *SetUserPasswordResponseBody {
	return s.Body
}

func (s *SetUserPasswordResponse) SetHeaders(v map[string]*string) *SetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *SetUserPasswordResponse) SetStatusCode(v int32) *SetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPasswordResponse) SetBody(v *SetUserPasswordResponseBody) *SetUserPasswordResponse {
	s.Body = v
	return s
}

func (s *SetUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
