// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRuleHitCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetRuleHitCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetRuleHitCountResponse
	GetStatusCode() *int32
	SetBody(v *ResetRuleHitCountResponseBody) *ResetRuleHitCountResponse
	GetBody() *ResetRuleHitCountResponseBody
}

type ResetRuleHitCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetRuleHitCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetRuleHitCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetRuleHitCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetRuleHitCountResponse) GetBody() *ResetRuleHitCountResponseBody {
	return s.Body
}

func (s *ResetRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetRuleHitCountResponse) SetStatusCode(v int32) *ResetRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetRuleHitCountResponse) SetBody(v *ResetRuleHitCountResponseBody) *ResetRuleHitCountResponse {
	s.Body = v
	return s
}

func (s *ResetRuleHitCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
