// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSensitiveDefineRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSensitiveDefineRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSensitiveDefineRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetSensitiveDefineRuleConfigResponseBody) *SetSensitiveDefineRuleConfigResponse
	GetBody() *SetSensitiveDefineRuleConfigResponseBody
}

type SetSensitiveDefineRuleConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSensitiveDefineRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSensitiveDefineRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSensitiveDefineRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetSensitiveDefineRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSensitiveDefineRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSensitiveDefineRuleConfigResponse) GetBody() *SetSensitiveDefineRuleConfigResponseBody {
	return s.Body
}

func (s *SetSensitiveDefineRuleConfigResponse) SetHeaders(v map[string]*string) *SetSensitiveDefineRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponse) SetStatusCode(v int32) *SetSensitiveDefineRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponse) SetBody(v *SetSensitiveDefineRuleConfigResponseBody) *SetSensitiveDefineRuleConfigResponse {
	s.Body = v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
