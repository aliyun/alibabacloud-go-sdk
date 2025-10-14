// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWafRulesetResponseBody) *DeleteWafRulesetResponse
	GetBody() *DeleteWafRulesetResponseBody
}

type DeleteWafRulesetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWafRulesetResponse) GetBody() *DeleteWafRulesetResponseBody {
	return s.Body
}

func (s *DeleteWafRulesetResponse) SetHeaders(v map[string]*string) *DeleteWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteWafRulesetResponse) SetStatusCode(v int32) *DeleteWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWafRulesetResponse) SetBody(v *DeleteWafRulesetResponseBody) *DeleteWafRulesetResponse {
	s.Body = v
	return s
}

func (s *DeleteWafRulesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
