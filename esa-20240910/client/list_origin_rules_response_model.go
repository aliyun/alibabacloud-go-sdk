// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOriginRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOriginRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListOriginRulesResponseBody) *ListOriginRulesResponse
	GetBody() *ListOriginRulesResponseBody
}

type ListOriginRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOriginRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOriginRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOriginRulesResponse) GoString() string {
	return s.String()
}

func (s *ListOriginRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOriginRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOriginRulesResponse) GetBody() *ListOriginRulesResponseBody {
	return s.Body
}

func (s *ListOriginRulesResponse) SetHeaders(v map[string]*string) *ListOriginRulesResponse {
	s.Headers = v
	return s
}

func (s *ListOriginRulesResponse) SetStatusCode(v int32) *ListOriginRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOriginRulesResponse) SetBody(v *ListOriginRulesResponseBody) *ListOriginRulesResponse {
	s.Body = v
	return s
}

func (s *ListOriginRulesResponse) Validate() error {
	return dara.Validate(s)
}
