// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertRulesResponseBody) *ListAlertRulesResponse
	GetBody() *ListAlertRulesResponseBody
}

type ListAlertRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertRulesResponse) GetBody() *ListAlertRulesResponseBody {
	return s.Body
}

func (s *ListAlertRulesResponse) SetHeaders(v map[string]*string) *ListAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertRulesResponse) SetStatusCode(v int32) *ListAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertRulesResponse) SetBody(v *ListAlertRulesResponseBody) *ListAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ListAlertRulesResponse) Validate() error {
	return dara.Validate(s)
}
