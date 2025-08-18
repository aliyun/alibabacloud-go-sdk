// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWafRulesResponseBody) *ListWafRulesResponse
	GetBody() *ListWafRulesResponseBody
}

type ListWafRulesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafRulesResponse) GetBody() *ListWafRulesResponseBody {
	return s.Body
}

func (s *ListWafRulesResponse) SetHeaders(v map[string]*string) *ListWafRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafRulesResponse) SetStatusCode(v int32) *ListWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafRulesResponse) SetBody(v *ListWafRulesResponseBody) *ListWafRulesResponse {
	s.Body = v
	return s
}

func (s *ListWafRulesResponse) Validate() error {
	return dara.Validate(s)
}
