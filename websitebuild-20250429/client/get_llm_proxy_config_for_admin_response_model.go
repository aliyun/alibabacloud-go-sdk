// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmProxyConfigForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLlmProxyConfigForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLlmProxyConfigForAdminResponse
	GetStatusCode() *int32
	SetBody(v *GetLlmProxyConfigForAdminResponseBody) *GetLlmProxyConfigForAdminResponse
	GetBody() *GetLlmProxyConfigForAdminResponseBody
}

type GetLlmProxyConfigForAdminResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLlmProxyConfigForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLlmProxyConfigForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLlmProxyConfigForAdminResponse) GoString() string {
	return s.String()
}

func (s *GetLlmProxyConfigForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLlmProxyConfigForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLlmProxyConfigForAdminResponse) GetBody() *GetLlmProxyConfigForAdminResponseBody {
	return s.Body
}

func (s *GetLlmProxyConfigForAdminResponse) SetHeaders(v map[string]*string) *GetLlmProxyConfigForAdminResponse {
	s.Headers = v
	return s
}

func (s *GetLlmProxyConfigForAdminResponse) SetStatusCode(v int32) *GetLlmProxyConfigForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponse) SetBody(v *GetLlmProxyConfigForAdminResponseBody) *GetLlmProxyConfigForAdminResponse {
	s.Body = v
	return s
}

func (s *GetLlmProxyConfigForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
