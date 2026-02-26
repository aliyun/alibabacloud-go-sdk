// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCredentialDescriptionResponseBody) *UpdateCredentialDescriptionResponse
	GetBody() *UpdateCredentialDescriptionResponseBody
}

type UpdateCredentialDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialDescriptionResponse) GetBody() *UpdateCredentialDescriptionResponseBody {
	return s.Body
}

func (s *UpdateCredentialDescriptionResponse) SetHeaders(v map[string]*string) *UpdateCredentialDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialDescriptionResponse) SetStatusCode(v int32) *UpdateCredentialDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialDescriptionResponse) SetBody(v *UpdateCredentialDescriptionResponseBody) *UpdateCredentialDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
