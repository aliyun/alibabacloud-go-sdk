// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretReferencesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecretReferencesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecretReferencesResponse
	GetStatusCode() *int32
	SetBody(v *ListSecretReferencesResponseBody) *ListSecretReferencesResponse
	GetBody() *ListSecretReferencesResponseBody
}

type ListSecretReferencesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretReferencesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretReferencesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesResponse) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecretReferencesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecretReferencesResponse) GetBody() *ListSecretReferencesResponseBody {
	return s.Body
}

func (s *ListSecretReferencesResponse) SetHeaders(v map[string]*string) *ListSecretReferencesResponse {
	s.Headers = v
	return s
}

func (s *ListSecretReferencesResponse) SetStatusCode(v int32) *ListSecretReferencesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretReferencesResponse) SetBody(v *ListSecretReferencesResponseBody) *ListSecretReferencesResponse {
	s.Body = v
	return s
}

func (s *ListSecretReferencesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
