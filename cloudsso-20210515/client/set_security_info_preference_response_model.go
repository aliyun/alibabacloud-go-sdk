// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSecurityInfoPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSecurityInfoPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *SetSecurityInfoPreferenceResponseBody) *SetSecurityInfoPreferenceResponse
	GetBody() *SetSecurityInfoPreferenceResponseBody
}

type SetSecurityInfoPreferenceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecurityInfoPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecurityInfoPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSecurityInfoPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSecurityInfoPreferenceResponse) GetBody() *SetSecurityInfoPreferenceResponseBody {
	return s.Body
}

func (s *SetSecurityInfoPreferenceResponse) SetHeaders(v map[string]*string) *SetSecurityInfoPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetSecurityInfoPreferenceResponse) SetStatusCode(v int32) *SetSecurityInfoPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecurityInfoPreferenceResponse) SetBody(v *SetSecurityInfoPreferenceResponseBody) *SetSecurityInfoPreferenceResponse {
	s.Body = v
	return s
}

func (s *SetSecurityInfoPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
