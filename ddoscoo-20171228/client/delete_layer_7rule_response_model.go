// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer7RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLayer7RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLayer7RuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLayer7RuleResponseBody) *DeleteLayer7RuleResponse
	GetBody() *DeleteLayer7RuleResponseBody
}

type DeleteLayer7RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLayer7RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLayer7RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLayer7RuleResponse) GetBody() *DeleteLayer7RuleResponseBody {
	return s.Body
}

func (s *DeleteLayer7RuleResponse) SetHeaders(v map[string]*string) *DeleteLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer7RuleResponse) SetStatusCode(v int32) *DeleteLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer7RuleResponse) SetBody(v *DeleteLayer7RuleResponseBody) *DeleteLayer7RuleResponse {
	s.Body = v
	return s
}

func (s *DeleteLayer7RuleResponse) Validate() error {
	return dara.Validate(s)
}
