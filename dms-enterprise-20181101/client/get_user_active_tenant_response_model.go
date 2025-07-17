// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserActiveTenantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserActiveTenantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserActiveTenantResponse
	GetStatusCode() *int32
	SetBody(v *GetUserActiveTenantResponseBody) *GetUserActiveTenantResponse
	GetBody() *GetUserActiveTenantResponseBody
}

type GetUserActiveTenantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserActiveTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserActiveTenantResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserActiveTenantResponse) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserActiveTenantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserActiveTenantResponse) GetBody() *GetUserActiveTenantResponseBody {
	return s.Body
}

func (s *GetUserActiveTenantResponse) SetHeaders(v map[string]*string) *GetUserActiveTenantResponse {
	s.Headers = v
	return s
}

func (s *GetUserActiveTenantResponse) SetStatusCode(v int32) *GetUserActiveTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserActiveTenantResponse) SetBody(v *GetUserActiveTenantResponseBody) *GetUserActiveTenantResponse {
	s.Body = v
	return s
}

func (s *GetUserActiveTenantResponse) Validate() error {
	return dara.Validate(s)
}
