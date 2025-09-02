// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIAlarmRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIAlarmRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIAlarmRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDIAlarmRulesResponseBody) *ListDIAlarmRulesResponse
	GetBody() *ListDIAlarmRulesResponseBody
}

type ListDIAlarmRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIAlarmRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIAlarmRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIAlarmRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIAlarmRulesResponse) GetBody() *ListDIAlarmRulesResponseBody {
	return s.Body
}

func (s *ListDIAlarmRulesResponse) SetHeaders(v map[string]*string) *ListDIAlarmRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDIAlarmRulesResponse) SetStatusCode(v int32) *ListDIAlarmRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIAlarmRulesResponse) SetBody(v *ListDIAlarmRulesResponseBody) *ListDIAlarmRulesResponse {
	s.Body = v
	return s
}

func (s *ListDIAlarmRulesResponse) Validate() error {
	return dara.Validate(s)
}
