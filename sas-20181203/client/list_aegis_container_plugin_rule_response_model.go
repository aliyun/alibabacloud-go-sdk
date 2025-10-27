// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisContainerPluginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAegisContainerPluginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAegisContainerPluginRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListAegisContainerPluginRuleResponseBody) *ListAegisContainerPluginRuleResponse
	GetBody() *ListAegisContainerPluginRuleResponseBody
}

type ListAegisContainerPluginRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAegisContainerPluginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAegisContainerPluginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAegisContainerPluginRuleResponse) GoString() string {
	return s.String()
}

func (s *ListAegisContainerPluginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAegisContainerPluginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAegisContainerPluginRuleResponse) GetBody() *ListAegisContainerPluginRuleResponseBody {
	return s.Body
}

func (s *ListAegisContainerPluginRuleResponse) SetHeaders(v map[string]*string) *ListAegisContainerPluginRuleResponse {
	s.Headers = v
	return s
}

func (s *ListAegisContainerPluginRuleResponse) SetStatusCode(v int32) *ListAegisContainerPluginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAegisContainerPluginRuleResponse) SetBody(v *ListAegisContainerPluginRuleResponseBody) *ListAegisContainerPluginRuleResponse {
	s.Body = v
	return s
}

func (s *ListAegisContainerPluginRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
