// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebPreciseAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebPreciseAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebPreciseAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebPreciseAccessRuleResponseBody) *DeleteWebPreciseAccessRuleResponse
	GetBody() *DeleteWebPreciseAccessRuleResponseBody
}

type DeleteWebPreciseAccessRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebPreciseAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebPreciseAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebPreciseAccessRuleResponse) GetBody() *DeleteWebPreciseAccessRuleResponseBody {
	return s.Body
}

func (s *DeleteWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebPreciseAccessRuleResponse) SetStatusCode(v int32) *DeleteWebPreciseAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleResponse) SetBody(v *DeleteWebPreciseAccessRuleResponseBody) *DeleteWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWebPreciseAccessRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
