// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakebaseTenantTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLakebaseTenantTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLakebaseTenantTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetLakebaseTenantTokenResponseBody) *GetLakebaseTenantTokenResponse
	GetBody() *GetLakebaseTenantTokenResponseBody
}

type GetLakebaseTenantTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLakebaseTenantTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLakebaseTenantTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLakebaseTenantTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLakebaseTenantTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLakebaseTenantTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLakebaseTenantTokenResponse) GetBody() *GetLakebaseTenantTokenResponseBody {
	return s.Body
}

func (s *GetLakebaseTenantTokenResponse) SetHeaders(v map[string]*string) *GetLakebaseTenantTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLakebaseTenantTokenResponse) SetStatusCode(v int32) *GetLakebaseTenantTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLakebaseTenantTokenResponse) SetBody(v *GetLakebaseTenantTokenResponseBody) *GetLakebaseTenantTokenResponse {
	s.Body = v
	return s
}

func (s *GetLakebaseTenantTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
