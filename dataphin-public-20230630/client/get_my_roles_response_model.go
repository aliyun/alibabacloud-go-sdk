// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMyRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMyRolesResponse
	GetStatusCode() *int32
	SetBody(v *GetMyRolesResponseBody) *GetMyRolesResponse
	GetBody() *GetMyRolesResponseBody
}

type GetMyRolesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMyRolesResponse) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMyRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMyRolesResponse) GetBody() *GetMyRolesResponseBody {
	return s.Body
}

func (s *GetMyRolesResponse) SetHeaders(v map[string]*string) *GetMyRolesResponse {
	s.Headers = v
	return s
}

func (s *GetMyRolesResponse) SetStatusCode(v int32) *GetMyRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyRolesResponse) SetBody(v *GetMyRolesResponseBody) *GetMyRolesResponse {
	s.Body = v
	return s
}

func (s *GetMyRolesResponse) Validate() error {
	return dara.Validate(s)
}
