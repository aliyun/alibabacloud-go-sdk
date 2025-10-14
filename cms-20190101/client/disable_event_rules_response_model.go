// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEventRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableEventRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableEventRulesResponse
	GetStatusCode() *int32
	SetBody(v *DisableEventRulesResponseBody) *DisableEventRulesResponse
	GetBody() *DisableEventRulesResponseBody
}

type DisableEventRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableEventRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableEventRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableEventRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableEventRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableEventRulesResponse) GetBody() *DisableEventRulesResponseBody {
	return s.Body
}

func (s *DisableEventRulesResponse) SetHeaders(v map[string]*string) *DisableEventRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableEventRulesResponse) SetStatusCode(v int32) *DisableEventRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableEventRulesResponse) SetBody(v *DisableEventRulesResponseBody) *DisableEventRulesResponse {
	s.Body = v
	return s
}

func (s *DisableEventRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
