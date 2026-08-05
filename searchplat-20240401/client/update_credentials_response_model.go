// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCredentialsResponseBody) *UpdateCredentialsResponse
	GetBody() *UpdateCredentialsResponseBody
}

type UpdateCredentialsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialsResponse) GetBody() *UpdateCredentialsResponseBody {
	return s.Body
}

func (s *UpdateCredentialsResponse) SetHeaders(v map[string]*string) *UpdateCredentialsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialsResponse) SetStatusCode(v int32) *UpdateCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialsResponse) SetBody(v *UpdateCredentialsResponseBody) *UpdateCredentialsResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
