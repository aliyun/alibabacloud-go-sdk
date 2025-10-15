// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationFederatedCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationFederatedCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationFederatedCredentialResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationFederatedCredentialResponseBody) *DeleteApplicationFederatedCredentialResponse
	GetBody() *DeleteApplicationFederatedCredentialResponseBody
}

type DeleteApplicationFederatedCredentialResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationFederatedCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationFederatedCredentialResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationFederatedCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationFederatedCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationFederatedCredentialResponse) GetBody() *DeleteApplicationFederatedCredentialResponseBody {
	return s.Body
}

func (s *DeleteApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *DeleteApplicationFederatedCredentialResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationFederatedCredentialResponse) SetStatusCode(v int32) *DeleteApplicationFederatedCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationFederatedCredentialResponse) SetBody(v *DeleteApplicationFederatedCredentialResponseBody) *DeleteApplicationFederatedCredentialResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationFederatedCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
