// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpRequestHeaderModificationRuleResponseBody) *DeleteHttpRequestHeaderModificationRuleResponse
	GetBody() *DeleteHttpRequestHeaderModificationRuleResponseBody
}

type DeleteHttpRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) GetBody() *DeleteHttpRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *DeleteHttpRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *DeleteHttpRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) SetBody(v *DeleteHttpRequestHeaderModificationRuleResponseBody) *DeleteHttpRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleResponse) Validate() error {
	return dara.Validate(s)
}
