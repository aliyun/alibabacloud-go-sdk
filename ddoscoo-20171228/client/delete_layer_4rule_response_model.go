// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayer4RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLayer4RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLayer4RuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLayer4RuleResponseBody) *DeleteLayer4RuleResponse
	GetBody() *DeleteLayer4RuleResponseBody
}

type DeleteLayer4RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLayer4RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLayer4RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLayer4RuleResponse) GetBody() *DeleteLayer4RuleResponseBody {
	return s.Body
}

func (s *DeleteLayer4RuleResponse) SetHeaders(v map[string]*string) *DeleteLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer4RuleResponse) SetStatusCode(v int32) *DeleteLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer4RuleResponse) SetBody(v *DeleteLayer4RuleResponseBody) *DeleteLayer4RuleResponse {
	s.Body = v
	return s
}

func (s *DeleteLayer4RuleResponse) Validate() error {
	return dara.Validate(s)
}
