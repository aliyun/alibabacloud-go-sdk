// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepoBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepoBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepoBuildRuleResponseBody) *DeleteRepoBuildRuleResponse
	GetBody() *DeleteRepoBuildRuleResponseBody
}

type DeleteRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepoBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepoBuildRuleResponse) GetBody() *DeleteRepoBuildRuleResponseBody {
	return s.Body
}

func (s *DeleteRepoBuildRuleResponse) SetHeaders(v map[string]*string) *DeleteRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoBuildRuleResponse) SetStatusCode(v int32) *DeleteRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoBuildRuleResponse) SetBody(v *DeleteRepoBuildRuleResponseBody) *DeleteRepoBuildRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRepoBuildRuleResponse) Validate() error {
	return dara.Validate(s)
}
