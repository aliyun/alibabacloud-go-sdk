// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse
	GetBody() *DeleteRuleResponseBody
}

type DeleteRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRuleResponse) GetBody() *DeleteRuleResponseBody {
	return s.Body
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRuleResponse) Validate() error {
	return dara.Validate(s)
}
