// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionByCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPermissionByCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPermissionByCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetPermissionByCodeResponseBody) *GetPermissionByCodeResponse
	GetBody() *GetPermissionByCodeResponseBody
}

type GetPermissionByCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionByCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionByCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPermissionByCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPermissionByCodeResponse) GetBody() *GetPermissionByCodeResponseBody {
	return s.Body
}

func (s *GetPermissionByCodeResponse) SetHeaders(v map[string]*string) *GetPermissionByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionByCodeResponse) SetStatusCode(v int32) *GetPermissionByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionByCodeResponse) SetBody(v *GetPermissionByCodeResponseBody) *GetPermissionByCodeResponse {
	s.Body = v
	return s
}

func (s *GetPermissionByCodeResponse) Validate() error {
	return dara.Validate(s)
}
