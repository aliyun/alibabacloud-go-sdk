// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSCIMServerCredentialStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSCIMServerCredentialStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSCIMServerCredentialStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSCIMServerCredentialStatusResponseBody) *UpdateSCIMServerCredentialStatusResponse
	GetBody() *UpdateSCIMServerCredentialStatusResponseBody
}

type UpdateSCIMServerCredentialStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSCIMServerCredentialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSCIMServerCredentialStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSCIMServerCredentialStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSCIMServerCredentialStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSCIMServerCredentialStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSCIMServerCredentialStatusResponse) GetBody() *UpdateSCIMServerCredentialStatusResponseBody {
	return s.Body
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetHeaders(v map[string]*string) *UpdateSCIMServerCredentialStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetStatusCode(v int32) *UpdateSCIMServerCredentialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponse) SetBody(v *UpdateSCIMServerCredentialStatusResponseBody) *UpdateSCIMServerCredentialStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateSCIMServerCredentialStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
