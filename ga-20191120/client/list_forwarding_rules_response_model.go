// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListForwardingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListForwardingRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListForwardingRulesResponseBody) *ListForwardingRulesResponse
	GetBody() *ListForwardingRulesResponseBody
}

type ListForwardingRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListForwardingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListForwardingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListForwardingRulesResponse) GetBody() *ListForwardingRulesResponseBody {
	return s.Body
}

func (s *ListForwardingRulesResponse) SetHeaders(v map[string]*string) *ListForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListForwardingRulesResponse) SetStatusCode(v int32) *ListForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListForwardingRulesResponse) SetBody(v *ListForwardingRulesResponseBody) *ListForwardingRulesResponse {
	s.Body = v
	return s
}

func (s *ListForwardingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
