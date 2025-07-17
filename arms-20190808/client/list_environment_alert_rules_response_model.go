// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentAlertRulesResponseBody) *ListEnvironmentAlertRulesResponse
	GetBody() *ListEnvironmentAlertRulesResponseBody
}

type ListEnvironmentAlertRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentAlertRulesResponse) GetBody() *ListEnvironmentAlertRulesResponseBody {
	return s.Body
}

func (s *ListEnvironmentAlertRulesResponse) SetHeaders(v map[string]*string) *ListEnvironmentAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentAlertRulesResponse) SetStatusCode(v int32) *ListEnvironmentAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponse) SetBody(v *ListEnvironmentAlertRulesResponseBody) *ListEnvironmentAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentAlertRulesResponse) Validate() error {
	return dara.Validate(s)
}
