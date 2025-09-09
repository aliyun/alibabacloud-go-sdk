// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostAccountsForHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostAccountsForHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *ListHostAccountsForHostShareKeyResponseBody) *ListHostAccountsForHostShareKeyResponse
	GetBody() *ListHostAccountsForHostShareKeyResponseBody
}

type ListHostAccountsForHostShareKeyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostAccountsForHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostAccountsForHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostAccountsForHostShareKeyResponse) GetBody() *ListHostAccountsForHostShareKeyResponseBody {
	return s.Body
}

func (s *ListHostAccountsForHostShareKeyResponse) SetHeaders(v map[string]*string) *ListHostAccountsForHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponse) SetStatusCode(v int32) *ListHostAccountsForHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponse) SetBody(v *ListHostAccountsForHostShareKeyResponseBody) *ListHostAccountsForHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}
