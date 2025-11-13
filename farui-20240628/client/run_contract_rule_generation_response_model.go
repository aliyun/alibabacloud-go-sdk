// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractRuleGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunContractRuleGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunContractRuleGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunContractRuleGenerationResponseBody) *RunContractRuleGenerationResponse
	GetBody() *RunContractRuleGenerationResponseBody
}

type RunContractRuleGenerationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContractRuleGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContractRuleGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunContractRuleGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunContractRuleGenerationResponse) GetBody() *RunContractRuleGenerationResponseBody {
	return s.Body
}

func (s *RunContractRuleGenerationResponse) SetHeaders(v map[string]*string) *RunContractRuleGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunContractRuleGenerationResponse) SetStatusCode(v int32) *RunContractRuleGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContractRuleGenerationResponse) SetBody(v *RunContractRuleGenerationResponseBody) *RunContractRuleGenerationResponse {
	s.Body = v
	return s
}

func (s *RunContractRuleGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
