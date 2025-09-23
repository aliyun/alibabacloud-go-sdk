// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer7CCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLayer7CCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLayer7CCRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLayer7CCRuleResponseBody) *DeleteLayer7CCRuleResponse
	GetBody() *DeleteLayer7CCRuleResponseBody
}

type DeleteLayer7CCRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLayer7CCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLayer7CCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLayer7CCRuleResponse) GetBody() *DeleteLayer7CCRuleResponseBody {
	return s.Body
}

func (s *DeleteLayer7CCRuleResponse) SetHeaders(v map[string]*string) *DeleteLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer7CCRuleResponse) SetStatusCode(v int32) *DeleteLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer7CCRuleResponse) SetBody(v *DeleteLayer7CCRuleResponseBody) *DeleteLayer7CCRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteLayer7CCRuleResponse) Validate() error {
	return dara.Validate(s)
}
