// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *GetUserWafRulesetResponseBody) *GetUserWafRulesetResponse
	GetBody() *GetUserWafRulesetResponseBody
}

type GetUserWafRulesetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *GetUserWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserWafRulesetResponse) GetBody() *GetUserWafRulesetResponseBody {
	return s.Body
}

func (s *GetUserWafRulesetResponse) SetHeaders(v map[string]*string) *GetUserWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *GetUserWafRulesetResponse) SetStatusCode(v int32) *GetUserWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserWafRulesetResponse) SetBody(v *GetUserWafRulesetResponseBody) *GetUserWafRulesetResponse {
	s.Body = v
	return s
}

func (s *GetUserWafRulesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
