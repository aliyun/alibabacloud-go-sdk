// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserWafRulesetResponseBody) *DeleteUserWafRulesetResponse
	GetBody() *DeleteUserWafRulesetResponseBody
}

type DeleteUserWafRulesetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserWafRulesetResponse) GetBody() *DeleteUserWafRulesetResponseBody {
	return s.Body
}

func (s *DeleteUserWafRulesetResponse) SetHeaders(v map[string]*string) *DeleteUserWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserWafRulesetResponse) SetStatusCode(v int32) *DeleteUserWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserWafRulesetResponse) SetBody(v *DeleteUserWafRulesetResponseBody) *DeleteUserWafRulesetResponse {
	s.Body = v
	return s
}

func (s *DeleteUserWafRulesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
