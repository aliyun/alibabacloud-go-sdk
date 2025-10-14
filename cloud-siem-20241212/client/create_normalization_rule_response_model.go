// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNormalizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNormalizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateNormalizationRuleResponseBody) *CreateNormalizationRuleResponse
	GetBody() *CreateNormalizationRuleResponseBody
}

type CreateNormalizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNormalizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNormalizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateNormalizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNormalizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNormalizationRuleResponse) GetBody() *CreateNormalizationRuleResponseBody {
	return s.Body
}

func (s *CreateNormalizationRuleResponse) SetHeaders(v map[string]*string) *CreateNormalizationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateNormalizationRuleResponse) SetStatusCode(v int32) *CreateNormalizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNormalizationRuleResponse) SetBody(v *CreateNormalizationRuleResponseBody) *CreateNormalizationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateNormalizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
