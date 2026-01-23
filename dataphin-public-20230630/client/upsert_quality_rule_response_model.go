// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpsertQualityRuleResponseBody) *UpsertQualityRuleResponse
	GetBody() *UpsertQualityRuleResponseBody
}

type UpsertQualityRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertQualityRuleResponse) GetBody() *UpsertQualityRuleResponseBody {
	return s.Body
}

func (s *UpsertQualityRuleResponse) SetHeaders(v map[string]*string) *UpsertQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *UpsertQualityRuleResponse) SetStatusCode(v int32) *UpsertQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertQualityRuleResponse) SetBody(v *UpsertQualityRuleResponseBody) *UpsertQualityRuleResponse {
	s.Body = v
	return s
}

func (s *UpsertQualityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
