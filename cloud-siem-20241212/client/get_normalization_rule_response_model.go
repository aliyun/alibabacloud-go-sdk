// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNormalizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNormalizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetNormalizationRuleResponseBody) *GetNormalizationRuleResponse
	GetBody() *GetNormalizationRuleResponseBody
}

type GetNormalizationRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNormalizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNormalizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNormalizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNormalizationRuleResponse) GetBody() *GetNormalizationRuleResponseBody {
	return s.Body
}

func (s *GetNormalizationRuleResponse) SetHeaders(v map[string]*string) *GetNormalizationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetNormalizationRuleResponse) SetStatusCode(v int32) *GetNormalizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNormalizationRuleResponse) SetBody(v *GetNormalizationRuleResponseBody) *GetNormalizationRuleResponse {
	s.Body = v
	return s
}

func (s *GetNormalizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
