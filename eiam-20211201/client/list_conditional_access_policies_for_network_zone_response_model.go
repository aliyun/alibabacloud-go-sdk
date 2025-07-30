// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConditionalAccessPoliciesForNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *ListConditionalAccessPoliciesForNetworkZoneResponseBody) *ListConditionalAccessPoliciesForNetworkZoneResponse
	GetBody() *ListConditionalAccessPoliciesForNetworkZoneResponseBody
}

type ListConditionalAccessPoliciesForNetworkZoneResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConditionalAccessPoliciesForNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) GetBody() *ListConditionalAccessPoliciesForNetworkZoneResponseBody {
	return s.Body
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) SetHeaders(v map[string]*string) *ListConditionalAccessPoliciesForNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) SetStatusCode(v int32) *ListConditionalAccessPoliciesForNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) SetBody(v *ListConditionalAccessPoliciesForNetworkZoneResponseBody) *ListConditionalAccessPoliciesForNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponse) Validate() error {
	return dara.Validate(s)
}
