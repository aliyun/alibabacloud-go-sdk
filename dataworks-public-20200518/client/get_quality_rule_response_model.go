// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleResponseBody) *GetQualityRuleResponse
	GetBody() *GetQualityRuleResponseBody
}

type GetQualityRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleResponse) GetBody() *GetQualityRuleResponseBody {
	return s.Body
}

func (s *GetQualityRuleResponse) SetHeaders(v map[string]*string) *GetQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleResponse) SetStatusCode(v int32) *GetQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleResponse) SetBody(v *GetQualityRuleResponseBody) *GetQualityRuleResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
