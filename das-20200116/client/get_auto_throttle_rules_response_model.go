// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoThrottleRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoThrottleRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoThrottleRulesResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoThrottleRulesResponseBody) *GetAutoThrottleRulesResponse
	GetBody() *GetAutoThrottleRulesResponseBody
}

type GetAutoThrottleRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoThrottleRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoThrottleRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesResponse) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoThrottleRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoThrottleRulesResponse) GetBody() *GetAutoThrottleRulesResponseBody {
	return s.Body
}

func (s *GetAutoThrottleRulesResponse) SetHeaders(v map[string]*string) *GetAutoThrottleRulesResponse {
	s.Headers = v
	return s
}

func (s *GetAutoThrottleRulesResponse) SetStatusCode(v int32) *GetAutoThrottleRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoThrottleRulesResponse) SetBody(v *GetAutoThrottleRulesResponseBody) *GetAutoThrottleRulesResponse {
	s.Body = v
	return s
}

func (s *GetAutoThrottleRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
