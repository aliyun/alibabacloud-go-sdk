// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterBaseLineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaClusterBaseLineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaClusterBaseLineListResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaClusterBaseLineListResponseBody) *GetOpaClusterBaseLineListResponse
	GetBody() *GetOpaClusterBaseLineListResponseBody
}

type GetOpaClusterBaseLineListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaClusterBaseLineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaClusterBaseLineListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterBaseLineListResponse) GoString() string {
	return s.String()
}

func (s *GetOpaClusterBaseLineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaClusterBaseLineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaClusterBaseLineListResponse) GetBody() *GetOpaClusterBaseLineListResponseBody {
	return s.Body
}

func (s *GetOpaClusterBaseLineListResponse) SetHeaders(v map[string]*string) *GetOpaClusterBaseLineListResponse {
	s.Headers = v
	return s
}

func (s *GetOpaClusterBaseLineListResponse) SetStatusCode(v int32) *GetOpaClusterBaseLineListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponse) SetBody(v *GetOpaClusterBaseLineListResponseBody) *GetOpaClusterBaseLineListResponse {
	s.Body = v
	return s
}

func (s *GetOpaClusterBaseLineListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
