// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNormalizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNormalizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNormalizationRuleResponseBody) *UpdateNormalizationRuleResponse
	GetBody() *UpdateNormalizationRuleResponseBody
}

type UpdateNormalizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNormalizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNormalizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNormalizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNormalizationRuleResponse) GetBody() *UpdateNormalizationRuleResponseBody {
	return s.Body
}

func (s *UpdateNormalizationRuleResponse) SetHeaders(v map[string]*string) *UpdateNormalizationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateNormalizationRuleResponse) SetStatusCode(v int32) *UpdateNormalizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNormalizationRuleResponse) SetBody(v *UpdateNormalizationRuleResponseBody) *UpdateNormalizationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateNormalizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
