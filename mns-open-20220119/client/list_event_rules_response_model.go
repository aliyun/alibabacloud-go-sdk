// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListEventRulesResponseBody) *ListEventRulesResponse
	GetBody() *ListEventRulesResponseBody
}

type ListEventRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventRulesResponse) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventRulesResponse) GetBody() *ListEventRulesResponseBody {
	return s.Body
}

func (s *ListEventRulesResponse) SetHeaders(v map[string]*string) *ListEventRulesResponse {
	s.Headers = v
	return s
}

func (s *ListEventRulesResponse) SetStatusCode(v int32) *ListEventRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventRulesResponse) SetBody(v *ListEventRulesResponseBody) *ListEventRulesResponse {
	s.Body = v
	return s
}

func (s *ListEventRulesResponse) Validate() error {
	return dara.Validate(s)
}
