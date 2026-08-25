// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSCIMServerCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSCIMServerCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSCIMServerCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateSCIMServerCredentialResponseBody) *CreateSCIMServerCredentialResponse
	GetBody() *CreateSCIMServerCredentialResponseBody
}

type CreateSCIMServerCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSCIMServerCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSCIMServerCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSCIMServerCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateSCIMServerCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSCIMServerCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSCIMServerCredentialResponse) GetBody() *CreateSCIMServerCredentialResponseBody {
	return s.Body
}

func (s *CreateSCIMServerCredentialResponse) SetHeaders(v map[string]*string) *CreateSCIMServerCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateSCIMServerCredentialResponse) SetStatusCode(v int32) *CreateSCIMServerCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSCIMServerCredentialResponse) SetBody(v *CreateSCIMServerCredentialResponseBody) *CreateSCIMServerCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateSCIMServerCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
