// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceRuleResponseBody) *DeleteResourceRuleResponse
	GetBody() *DeleteResourceRuleResponseBody
}

type DeleteResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceRuleResponse) GetBody() *DeleteResourceRuleResponseBody {
	return s.Body
}

func (s *DeleteResourceRuleResponse) SetHeaders(v map[string]*string) *DeleteResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceRuleResponse) SetStatusCode(v int32) *DeleteResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceRuleResponse) SetBody(v *DeleteResourceRuleResponseBody) *DeleteResourceRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
