// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateElasticRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateElasticRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateElasticRuleResponseBody) *CreateElasticRuleResponse
	GetBody() *CreateElasticRuleResponseBody
}

type CreateElasticRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateElasticRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateElasticRuleResponse) GetBody() *CreateElasticRuleResponseBody {
	return s.Body
}

func (s *CreateElasticRuleResponse) SetHeaders(v map[string]*string) *CreateElasticRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticRuleResponse) SetStatusCode(v int32) *CreateElasticRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticRuleResponse) SetBody(v *CreateElasticRuleResponseBody) *CreateElasticRuleResponse {
	s.Body = v
	return s
}

func (s *CreateElasticRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
