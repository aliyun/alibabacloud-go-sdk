// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactiveConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactiveConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeactiveConfigRulesResponseBody) *DeactiveConfigRulesResponse
	GetBody() *DeactiveConfigRulesResponseBody
}

type DeactiveConfigRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactiveConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactiveConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactiveConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactiveConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactiveConfigRulesResponse) GetBody() *DeactiveConfigRulesResponseBody {
	return s.Body
}

func (s *DeactiveConfigRulesResponse) SetHeaders(v map[string]*string) *DeactiveConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeactiveConfigRulesResponse) SetStatusCode(v int32) *DeactiveConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactiveConfigRulesResponse) SetBody(v *DeactiveConfigRulesResponseBody) *DeactiveConfigRulesResponse {
	s.Body = v
	return s
}

func (s *DeactiveConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
