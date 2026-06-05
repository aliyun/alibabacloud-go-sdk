// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSupabaseAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSupabaseAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSupabaseAuthConfigResponseBody) *GetAppSupabaseAuthConfigResponse
	GetBody() *GetAppSupabaseAuthConfigResponseBody
}

type GetAppSupabaseAuthConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSupabaseAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSupabaseAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSupabaseAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSupabaseAuthConfigResponse) GetBody() *GetAppSupabaseAuthConfigResponseBody {
	return s.Body
}

func (s *GetAppSupabaseAuthConfigResponse) SetHeaders(v map[string]*string) *GetAppSupabaseAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppSupabaseAuthConfigResponse) SetStatusCode(v int32) *GetAppSupabaseAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponse) SetBody(v *GetAppSupabaseAuthConfigResponseBody) *GetAppSupabaseAuthConfigResponse {
	s.Body = v
	return s
}

func (s *GetAppSupabaseAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
