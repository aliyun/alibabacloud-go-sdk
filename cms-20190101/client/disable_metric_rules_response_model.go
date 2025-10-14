// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableMetricRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableMetricRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableMetricRulesResponse
	GetStatusCode() *int32
	SetBody(v *DisableMetricRulesResponseBody) *DisableMetricRulesResponse
	GetBody() *DisableMetricRulesResponseBody
}

type DisableMetricRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableMetricRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableMetricRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableMetricRulesResponse) GetBody() *DisableMetricRulesResponseBody {
	return s.Body
}

func (s *DisableMetricRulesResponse) SetHeaders(v map[string]*string) *DisableMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableMetricRulesResponse) SetStatusCode(v int32) *DisableMetricRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMetricRulesResponse) SetBody(v *DisableMetricRulesResponseBody) *DisableMetricRulesResponse {
	s.Body = v
	return s
}

func (s *DisableMetricRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
