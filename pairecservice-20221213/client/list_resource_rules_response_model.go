// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceRulesResponseBody) *ListResourceRulesResponse
	GetBody() *ListResourceRulesResponseBody
}

type ListResourceRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceRulesResponse) GetBody() *ListResourceRulesResponseBody {
	return s.Body
}

func (s *ListResourceRulesResponse) SetHeaders(v map[string]*string) *ListResourceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRulesResponse) SetStatusCode(v int32) *ListResourceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRulesResponse) SetBody(v *ListResourceRulesResponseBody) *ListResourceRulesResponse {
	s.Body = v
	return s
}

func (s *ListResourceRulesResponse) Validate() error {
	return dara.Validate(s)
}
