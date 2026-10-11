// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEgressProxyRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEgressProxyRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEgressProxyRulesResponseBody) *DeleteEgressProxyRulesResponse
	GetBody() *DeleteEgressProxyRulesResponseBody
}

type DeleteEgressProxyRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEgressProxyRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEgressProxyRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEgressProxyRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEgressProxyRulesResponse) GetBody() *DeleteEgressProxyRulesResponseBody {
	return s.Body
}

func (s *DeleteEgressProxyRulesResponse) SetHeaders(v map[string]*string) *DeleteEgressProxyRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteEgressProxyRulesResponse) SetStatusCode(v int32) *DeleteEgressProxyRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEgressProxyRulesResponse) SetBody(v *DeleteEgressProxyRulesResponseBody) *DeleteEgressProxyRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteEgressProxyRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
