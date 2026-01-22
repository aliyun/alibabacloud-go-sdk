// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePasskeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePasskeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePasskeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePasskeyResponseBody) *UpdatePasskeyResponse
	GetBody() *UpdatePasskeyResponseBody
}

type UpdatePasskeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePasskeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePasskeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePasskeyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePasskeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePasskeyResponse) GetBody() *UpdatePasskeyResponseBody {
	return s.Body
}

func (s *UpdatePasskeyResponse) SetHeaders(v map[string]*string) *UpdatePasskeyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePasskeyResponse) SetStatusCode(v int32) *UpdatePasskeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePasskeyResponse) SetBody(v *UpdatePasskeyResponseBody) *UpdatePasskeyResponse {
	s.Body = v
	return s
}

func (s *UpdatePasskeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
