// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEgressProxyRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEgressProxyRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateEgressProxyRulesResponseBody) *CreateEgressProxyRulesResponse
	GetBody() *CreateEgressProxyRulesResponseBody
}

type CreateEgressProxyRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEgressProxyRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEgressProxyRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEgressProxyRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEgressProxyRulesResponse) GetBody() *CreateEgressProxyRulesResponseBody {
	return s.Body
}

func (s *CreateEgressProxyRulesResponse) SetHeaders(v map[string]*string) *CreateEgressProxyRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateEgressProxyRulesResponse) SetStatusCode(v int32) *CreateEgressProxyRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEgressProxyRulesResponse) SetBody(v *CreateEgressProxyRulesResponseBody) *CreateEgressProxyRulesResponse {
	s.Body = v
	return s
}

func (s *CreateEgressProxyRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
