// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkRuleResponseBody) *DeleteNetworkRuleResponse
	GetBody() *DeleteNetworkRuleResponseBody
}

type DeleteNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkRuleResponse) GetBody() *DeleteNetworkRuleResponseBody {
	return s.Body
}

func (s *DeleteNetworkRuleResponse) SetHeaders(v map[string]*string) *DeleteNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkRuleResponse) SetStatusCode(v int32) *DeleteNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkRuleResponse) SetBody(v *DeleteNetworkRuleResponseBody) *DeleteNetworkRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkRuleResponse) Validate() error {
	return dara.Validate(s)
}
