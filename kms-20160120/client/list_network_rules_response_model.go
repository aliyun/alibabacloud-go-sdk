// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkRulesResponseBody) *ListNetworkRulesResponse
	GetBody() *ListNetworkRulesResponseBody
}

type ListNetworkRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkRulesResponse) GetBody() *ListNetworkRulesResponseBody {
	return s.Body
}

func (s *ListNetworkRulesResponse) SetHeaders(v map[string]*string) *ListNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkRulesResponse) SetStatusCode(v int32) *ListNetworkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkRulesResponse) SetBody(v *ListNetworkRulesResponseBody) *ListNetworkRulesResponse {
	s.Body = v
	return s
}

func (s *ListNetworkRulesResponse) Validate() error {
	return dara.Validate(s)
}
