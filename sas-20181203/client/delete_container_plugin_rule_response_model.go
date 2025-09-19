// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerPluginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContainerPluginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContainerPluginRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContainerPluginRuleResponseBody) *DeleteContainerPluginRuleResponse
	GetBody() *DeleteContainerPluginRuleResponseBody
}

type DeleteContainerPluginRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContainerPluginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContainerPluginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerPluginRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerPluginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContainerPluginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContainerPluginRuleResponse) GetBody() *DeleteContainerPluginRuleResponseBody {
	return s.Body
}

func (s *DeleteContainerPluginRuleResponse) SetHeaders(v map[string]*string) *DeleteContainerPluginRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerPluginRuleResponse) SetStatusCode(v int32) *DeleteContainerPluginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerPluginRuleResponse) SetBody(v *DeleteContainerPluginRuleResponseBody) *DeleteContainerPluginRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteContainerPluginRuleResponse) Validate() error {
	return dara.Validate(s)
}
