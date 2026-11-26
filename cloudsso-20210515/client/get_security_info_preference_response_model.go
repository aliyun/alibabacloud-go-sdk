// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoPreferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityInfoPreferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityInfoPreferenceResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityInfoPreferenceResponseBody) *GetSecurityInfoPreferenceResponse
	GetBody() *GetSecurityInfoPreferenceResponseBody
}

type GetSecurityInfoPreferenceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityInfoPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityInfoPreferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoPreferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityInfoPreferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityInfoPreferenceResponse) GetBody() *GetSecurityInfoPreferenceResponseBody {
	return s.Body
}

func (s *GetSecurityInfoPreferenceResponse) SetHeaders(v map[string]*string) *GetSecurityInfoPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityInfoPreferenceResponse) SetStatusCode(v int32) *GetSecurityInfoPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityInfoPreferenceResponse) SetBody(v *GetSecurityInfoPreferenceResponseBody) *GetSecurityInfoPreferenceResponse {
	s.Body = v
	return s
}

func (s *GetSecurityInfoPreferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
