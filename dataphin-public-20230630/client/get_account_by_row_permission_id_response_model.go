// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountByRowPermissionIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountByRowPermissionIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountByRowPermissionIdResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountByRowPermissionIdResponseBody) *GetAccountByRowPermissionIdResponse
	GetBody() *GetAccountByRowPermissionIdResponseBody
}

type GetAccountByRowPermissionIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountByRowPermissionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountByRowPermissionIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdResponse) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountByRowPermissionIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountByRowPermissionIdResponse) GetBody() *GetAccountByRowPermissionIdResponseBody {
	return s.Body
}

func (s *GetAccountByRowPermissionIdResponse) SetHeaders(v map[string]*string) *GetAccountByRowPermissionIdResponse {
	s.Headers = v
	return s
}

func (s *GetAccountByRowPermissionIdResponse) SetStatusCode(v int32) *GetAccountByRowPermissionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponse) SetBody(v *GetAccountByRowPermissionIdResponseBody) *GetAccountByRowPermissionIdResponse {
	s.Body = v
	return s
}

func (s *GetAccountByRowPermissionIdResponse) Validate() error {
	return dara.Validate(s)
}
