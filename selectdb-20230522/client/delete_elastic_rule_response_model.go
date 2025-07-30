// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteElasticRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteElasticRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteElasticRuleResponseBody) *DeleteElasticRuleResponse
	GetBody() *DeleteElasticRuleResponseBody
}

type DeleteElasticRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteElasticRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteElasticRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteElasticRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteElasticRuleResponse) GetBody() *DeleteElasticRuleResponseBody {
	return s.Body
}

func (s *DeleteElasticRuleResponse) SetHeaders(v map[string]*string) *DeleteElasticRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticRuleResponse) SetStatusCode(v int32) *DeleteElasticRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticRuleResponse) SetBody(v *DeleteElasticRuleResponseBody) *DeleteElasticRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteElasticRuleResponse) Validate() error {
	return dara.Validate(s)
}
