// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMiniAppBindingForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMiniAppBindingForAdminResponse
	GetStatusCode() *int32
	SetBody(v *GetMiniAppBindingForAdminResponseBody) *GetMiniAppBindingForAdminResponse
	GetBody() *GetMiniAppBindingForAdminResponseBody
}

type GetMiniAppBindingForAdminResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMiniAppBindingForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMiniAppBindingForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingForAdminResponse) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMiniAppBindingForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMiniAppBindingForAdminResponse) GetBody() *GetMiniAppBindingForAdminResponseBody {
	return s.Body
}

func (s *GetMiniAppBindingForAdminResponse) SetHeaders(v map[string]*string) *GetMiniAppBindingForAdminResponse {
	s.Headers = v
	return s
}

func (s *GetMiniAppBindingForAdminResponse) SetStatusCode(v int32) *GetMiniAppBindingForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMiniAppBindingForAdminResponse) SetBody(v *GetMiniAppBindingForAdminResponseBody) *GetMiniAppBindingForAdminResponse {
	s.Body = v
	return s
}

func (s *GetMiniAppBindingForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
