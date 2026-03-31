// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSecurityPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSecurityPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *SetSecurityPreferenceResponseBody) *SetSecurityPreferenceResponse
	GetBody() *SetSecurityPreferenceResponseBody
}

type SetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecurityPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSecurityPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSecurityPreferenceResponse) GetBody() *SetSecurityPreferenceResponseBody {
	return s.Body
}

func (s *SetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *SetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetSecurityPreferenceResponse) SetStatusCode(v int32) *SetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecurityPreferenceResponse) SetBody(v *SetSecurityPreferenceResponseBody) *SetSecurityPreferenceResponse {
	s.Body = v
	return s
}

func (s *SetSecurityPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
