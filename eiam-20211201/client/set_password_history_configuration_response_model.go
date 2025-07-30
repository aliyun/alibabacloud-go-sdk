// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordHistoryConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPasswordHistoryConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPasswordHistoryConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetPasswordHistoryConfigurationResponseBody) *SetPasswordHistoryConfigurationResponse
	GetBody() *SetPasswordHistoryConfigurationResponseBody
}

type SetPasswordHistoryConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordHistoryConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordHistoryConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordHistoryConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPasswordHistoryConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPasswordHistoryConfigurationResponse) GetBody() *SetPasswordHistoryConfigurationResponseBody {
	return s.Body
}

func (s *SetPasswordHistoryConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordHistoryConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordHistoryConfigurationResponse) SetStatusCode(v int32) *SetPasswordHistoryConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordHistoryConfigurationResponse) SetBody(v *SetPasswordHistoryConfigurationResponseBody) *SetPasswordHistoryConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetPasswordHistoryConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
