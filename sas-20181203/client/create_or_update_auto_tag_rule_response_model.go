// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAutoTagRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateAutoTagRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateAutoTagRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateAutoTagRuleResponseBody) *CreateOrUpdateAutoTagRuleResponse
	GetBody() *CreateOrUpdateAutoTagRuleResponseBody
}

type CreateOrUpdateAutoTagRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateAutoTagRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateAutoTagRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAutoTagRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAutoTagRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateAutoTagRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateAutoTagRuleResponse) GetBody() *CreateOrUpdateAutoTagRuleResponseBody {
	return s.Body
}

func (s *CreateOrUpdateAutoTagRuleResponse) SetHeaders(v map[string]*string) *CreateOrUpdateAutoTagRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateAutoTagRuleResponse) SetStatusCode(v int32) *CreateOrUpdateAutoTagRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleResponse) SetBody(v *CreateOrUpdateAutoTagRuleResponseBody) *CreateOrUpdateAutoTagRuleResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateAutoTagRuleResponse) Validate() error {
	return dara.Validate(s)
}
