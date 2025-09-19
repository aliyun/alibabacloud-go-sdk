// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAegisContainerPluginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAegisContainerPluginRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAegisContainerPluginRuleResponseBody) *GetAegisContainerPluginRuleResponse
	GetBody() *GetAegisContainerPluginRuleResponseBody
}

type GetAegisContainerPluginRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAegisContainerPluginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAegisContainerPluginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAegisContainerPluginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAegisContainerPluginRuleResponse) GetBody() *GetAegisContainerPluginRuleResponseBody {
	return s.Body
}

func (s *GetAegisContainerPluginRuleResponse) SetHeaders(v map[string]*string) *GetAegisContainerPluginRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAegisContainerPluginRuleResponse) SetStatusCode(v int32) *GetAegisContainerPluginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponse) SetBody(v *GetAegisContainerPluginRuleResponseBody) *GetAegisContainerPluginRuleResponse {
	s.Body = v
	return s
}

func (s *GetAegisContainerPluginRuleResponse) Validate() error {
	return dara.Validate(s)
}
