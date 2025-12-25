// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserWafRulesetResponseBody) *UpdateUserWafRulesetResponse
	GetBody() *UpdateUserWafRulesetResponseBody
}

type UpdateUserWafRulesetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserWafRulesetResponse) GetBody() *UpdateUserWafRulesetResponseBody {
	return s.Body
}

func (s *UpdateUserWafRulesetResponse) SetHeaders(v map[string]*string) *UpdateUserWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserWafRulesetResponse) SetStatusCode(v int32) *UpdateUserWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserWafRulesetResponse) SetBody(v *UpdateUserWafRulesetResponseBody) *UpdateUserWafRulesetResponse {
	s.Body = v
	return s
}

func (s *UpdateUserWafRulesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
