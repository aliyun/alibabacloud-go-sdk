// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOriginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOriginRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOriginRuleResponseBody) *DeleteOriginRuleResponse
	GetBody() *DeleteOriginRuleResponseBody
}

type DeleteOriginRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOriginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOriginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteOriginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOriginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOriginRuleResponse) GetBody() *DeleteOriginRuleResponseBody {
	return s.Body
}

func (s *DeleteOriginRuleResponse) SetHeaders(v map[string]*string) *DeleteOriginRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteOriginRuleResponse) SetStatusCode(v int32) *DeleteOriginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOriginRuleResponse) SetBody(v *DeleteOriginRuleResponseBody) *DeleteOriginRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteOriginRuleResponse) Validate() error {
	return dara.Validate(s)
}
