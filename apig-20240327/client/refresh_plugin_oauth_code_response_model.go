// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshPluginOAuthCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshPluginOAuthCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshPluginOAuthCodeResponse
	GetStatusCode() *int32
	SetBody(v *RefreshPluginOAuthCodeResponseBody) *RefreshPluginOAuthCodeResponse
	GetBody() *RefreshPluginOAuthCodeResponseBody
}

type RefreshPluginOAuthCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshPluginOAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshPluginOAuthCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshPluginOAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *RefreshPluginOAuthCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshPluginOAuthCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshPluginOAuthCodeResponse) GetBody() *RefreshPluginOAuthCodeResponseBody {
	return s.Body
}

func (s *RefreshPluginOAuthCodeResponse) SetHeaders(v map[string]*string) *RefreshPluginOAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *RefreshPluginOAuthCodeResponse) SetStatusCode(v int32) *RefreshPluginOAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshPluginOAuthCodeResponse) SetBody(v *RefreshPluginOAuthCodeResponseBody) *RefreshPluginOAuthCodeResponse {
	s.Body = v
	return s
}

func (s *RefreshPluginOAuthCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
