// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRuleStrategyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientRuleStrategyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientRuleStrategyDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetClientRuleStrategyDetailResponseBody) *GetClientRuleStrategyDetailResponse
	GetBody() *GetClientRuleStrategyDetailResponseBody
}

type GetClientRuleStrategyDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientRuleStrategyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientRuleStrategyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyDetailResponse) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientRuleStrategyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientRuleStrategyDetailResponse) GetBody() *GetClientRuleStrategyDetailResponseBody {
	return s.Body
}

func (s *GetClientRuleStrategyDetailResponse) SetHeaders(v map[string]*string) *GetClientRuleStrategyDetailResponse {
	s.Headers = v
	return s
}

func (s *GetClientRuleStrategyDetailResponse) SetStatusCode(v int32) *GetClientRuleStrategyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientRuleStrategyDetailResponse) SetBody(v *GetClientRuleStrategyDetailResponseBody) *GetClientRuleStrategyDetailResponse {
	s.Body = v
	return s
}

func (s *GetClientRuleStrategyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
