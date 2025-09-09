// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLDAPAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceLDAPAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceLDAPAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceLDAPAuthServerResponseBody) *GetInstanceLDAPAuthServerResponse
	GetBody() *GetInstanceLDAPAuthServerResponseBody
}

type GetInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceLDAPAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceLDAPAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceLDAPAuthServerResponse) GetBody() *GetInstanceLDAPAuthServerResponseBody {
	return s.Body
}

func (s *GetInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *GetInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *GetInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponse) SetBody(v *GetInstanceLDAPAuthServerResponseBody) *GetInstanceLDAPAuthServerResponse {
	s.Body = v
	return s
}

func (s *GetInstanceLDAPAuthServerResponse) Validate() error {
	return dara.Validate(s)
}
