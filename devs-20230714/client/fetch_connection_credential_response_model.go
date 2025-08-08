// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchConnectionCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchConnectionCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchConnectionCredentialResponse
	GetStatusCode() *int32
	SetBody(v *OAuthCredential) *FetchConnectionCredentialResponse
	GetBody() *OAuthCredential
}

type FetchConnectionCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OAuthCredential   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchConnectionCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchConnectionCredentialResponse) GoString() string {
	return s.String()
}

func (s *FetchConnectionCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchConnectionCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchConnectionCredentialResponse) GetBody() *OAuthCredential {
	return s.Body
}

func (s *FetchConnectionCredentialResponse) SetHeaders(v map[string]*string) *FetchConnectionCredentialResponse {
	s.Headers = v
	return s
}

func (s *FetchConnectionCredentialResponse) SetStatusCode(v int32) *FetchConnectionCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchConnectionCredentialResponse) SetBody(v *OAuthCredential) *FetchConnectionCredentialResponse {
	s.Body = v
	return s
}

func (s *FetchConnectionCredentialResponse) Validate() error {
	return dara.Validate(s)
}
