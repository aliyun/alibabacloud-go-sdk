// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHttpDDoSAttackRuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHttpDDoSAttackRuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetHttpDDoSAttackRuleStatusResponseBody) *SetHttpDDoSAttackRuleStatusResponse
	GetBody() *SetHttpDDoSAttackRuleStatusResponseBody
}

type SetHttpDDoSAttackRuleStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackRuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHttpDDoSAttackRuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHttpDDoSAttackRuleStatusResponse) GetBody() *SetHttpDDoSAttackRuleStatusResponseBody {
	return s.Body
}

func (s *SetHttpDDoSAttackRuleStatusResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusResponse) SetStatusCode(v int32) *SetHttpDDoSAttackRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusResponse) SetBody(v *SetHttpDDoSAttackRuleStatusResponseBody) *SetHttpDDoSAttackRuleStatusResponse {
	s.Body = v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
