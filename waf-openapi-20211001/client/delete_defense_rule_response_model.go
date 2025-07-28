// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDefenseRuleResponseBody) *DeleteDefenseRuleResponse
	GetBody() *DeleteDefenseRuleResponseBody
}

type DeleteDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDefenseRuleResponse) GetBody() *DeleteDefenseRuleResponseBody {
	return s.Body
}

func (s *DeleteDefenseRuleResponse) SetHeaders(v map[string]*string) *DeleteDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseRuleResponse) SetStatusCode(v int32) *DeleteDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseRuleResponse) SetBody(v *DeleteDefenseRuleResponseBody) *DeleteDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
