// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityPreferenceResponseBody) *GetSecurityPreferenceResponse
	GetBody() *GetSecurityPreferenceResponseBody
}

type GetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityPreferenceResponse) GetBody() *GetSecurityPreferenceResponseBody {
	return s.Body
}

func (s *GetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *GetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityPreferenceResponse) SetStatusCode(v int32) *GetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityPreferenceResponse) SetBody(v *GetSecurityPreferenceResponseBody) *GetSecurityPreferenceResponse {
	s.Body = v
	return s
}

func (s *GetSecurityPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
