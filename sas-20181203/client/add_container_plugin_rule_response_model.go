// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerPluginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContainerPluginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContainerPluginRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddContainerPluginRuleResponseBody) *AddContainerPluginRuleResponse
	GetBody() *AddContainerPluginRuleResponseBody
}

type AddContainerPluginRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContainerPluginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContainerPluginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContainerPluginRuleResponse) GoString() string {
	return s.String()
}

func (s *AddContainerPluginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContainerPluginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContainerPluginRuleResponse) GetBody() *AddContainerPluginRuleResponseBody {
	return s.Body
}

func (s *AddContainerPluginRuleResponse) SetHeaders(v map[string]*string) *AddContainerPluginRuleResponse {
	s.Headers = v
	return s
}

func (s *AddContainerPluginRuleResponse) SetStatusCode(v int32) *AddContainerPluginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContainerPluginRuleResponse) SetBody(v *AddContainerPluginRuleResponseBody) *AddContainerPluginRuleResponse {
	s.Body = v
	return s
}

func (s *AddContainerPluginRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
