// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSupabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSupabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSupabaseInstanceResponseBody) *GetAppSupabaseInstanceResponse
	GetBody() *GetAppSupabaseInstanceResponseBody
}

type GetAppSupabaseInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSupabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSupabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSupabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSupabaseInstanceResponse) GetBody() *GetAppSupabaseInstanceResponseBody {
	return s.Body
}

func (s *GetAppSupabaseInstanceResponse) SetHeaders(v map[string]*string) *GetAppSupabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetAppSupabaseInstanceResponse) SetStatusCode(v int32) *GetAppSupabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSupabaseInstanceResponse) SetBody(v *GetAppSupabaseInstanceResponseBody) *GetAppSupabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *GetAppSupabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
