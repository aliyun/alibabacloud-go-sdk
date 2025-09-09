// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceLDAPAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceLDAPAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceLDAPAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceLDAPAuthServerResponseBody) *ModifyInstanceLDAPAuthServerResponse
	GetBody() *ModifyInstanceLDAPAuthServerResponseBody
}

type ModifyInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceLDAPAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceLDAPAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceLDAPAuthServerResponse) GetBody() *ModifyInstanceLDAPAuthServerResponseBody {
	return s.Body
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *ModifyInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *ModifyInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetBody(v *ModifyInstanceLDAPAuthServerResponseBody) *ModifyInstanceLDAPAuthServerResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponse) Validate() error {
	return dara.Validate(s)
}
