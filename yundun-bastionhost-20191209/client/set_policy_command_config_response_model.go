// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyCommandConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyCommandConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyCommandConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyCommandConfigResponseBody) *SetPolicyCommandConfigResponse
	GetBody() *SetPolicyCommandConfigResponseBody
}

type SetPolicyCommandConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyCommandConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyCommandConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyCommandConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyCommandConfigResponse) GetBody() *SetPolicyCommandConfigResponseBody {
	return s.Body
}

func (s *SetPolicyCommandConfigResponse) SetHeaders(v map[string]*string) *SetPolicyCommandConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyCommandConfigResponse) SetStatusCode(v int32) *SetPolicyCommandConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyCommandConfigResponse) SetBody(v *SetPolicyCommandConfigResponseBody) *SetPolicyCommandConfigResponse {
	s.Body = v
	return s
}

func (s *SetPolicyCommandConfigResponse) Validate() error {
	return dara.Validate(s)
}
