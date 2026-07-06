// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceCredentialResponseBody) *DeleteServiceCredentialResponse
	GetBody() *DeleteServiceCredentialResponseBody
}

type DeleteServiceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCredentialResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceCredentialResponse) GetBody() *DeleteServiceCredentialResponseBody {
	return s.Body
}

func (s *DeleteServiceCredentialResponse) SetHeaders(v map[string]*string) *DeleteServiceCredentialResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceCredentialResponse) SetStatusCode(v int32) *DeleteServiceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceCredentialResponse) SetBody(v *DeleteServiceCredentialResponseBody) *DeleteServiceCredentialResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
