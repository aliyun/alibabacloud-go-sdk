// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventCenterRuleNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventCenterRuleNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventCenterRuleNameResponse
	GetStatusCode() *int32
	SetBody(v *ListEventCenterRuleNameResponseBody) *ListEventCenterRuleNameResponse
	GetBody() *ListEventCenterRuleNameResponseBody
}

type ListEventCenterRuleNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventCenterRuleNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventCenterRuleNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventCenterRuleNameResponse) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventCenterRuleNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventCenterRuleNameResponse) GetBody() *ListEventCenterRuleNameResponseBody {
	return s.Body
}

func (s *ListEventCenterRuleNameResponse) SetHeaders(v map[string]*string) *ListEventCenterRuleNameResponse {
	s.Headers = v
	return s
}

func (s *ListEventCenterRuleNameResponse) SetStatusCode(v int32) *ListEventCenterRuleNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventCenterRuleNameResponse) SetBody(v *ListEventCenterRuleNameResponseBody) *ListEventCenterRuleNameResponse {
	s.Body = v
	return s
}

func (s *ListEventCenterRuleNameResponse) Validate() error {
	return dara.Validate(s)
}
