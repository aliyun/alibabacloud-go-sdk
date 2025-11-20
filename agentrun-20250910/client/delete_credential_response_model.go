// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CredentialResult) *DeleteCredentialResponse
	GetBody() *CredentialResult
}

type DeleteCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialResponse) GoString() string {
	return s.String()
}

func (s *DeleteCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCredentialResponse) GetBody() *CredentialResult {
	return s.Body
}

func (s *DeleteCredentialResponse) SetHeaders(v map[string]*string) *DeleteCredentialResponse {
	s.Headers = v
	return s
}

func (s *DeleteCredentialResponse) SetStatusCode(v int32) *DeleteCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCredentialResponse) SetBody(v *CredentialResult) *DeleteCredentialResponse {
	s.Body = v
	return s
}

func (s *DeleteCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
