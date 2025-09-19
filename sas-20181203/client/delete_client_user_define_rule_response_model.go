// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserDefineRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientUserDefineRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientUserDefineRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientUserDefineRuleResponseBody) *DeleteClientUserDefineRuleResponse
	GetBody() *DeleteClientUserDefineRuleResponseBody
}

type DeleteClientUserDefineRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientUserDefineRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientUserDefineRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserDefineRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientUserDefineRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientUserDefineRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientUserDefineRuleResponse) GetBody() *DeleteClientUserDefineRuleResponseBody {
	return s.Body
}

func (s *DeleteClientUserDefineRuleResponse) SetHeaders(v map[string]*string) *DeleteClientUserDefineRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientUserDefineRuleResponse) SetStatusCode(v int32) *DeleteClientUserDefineRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientUserDefineRuleResponse) SetBody(v *DeleteClientUserDefineRuleResponseBody) *DeleteClientUserDefineRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteClientUserDefineRuleResponse) Validate() error {
	return dara.Validate(s)
}
