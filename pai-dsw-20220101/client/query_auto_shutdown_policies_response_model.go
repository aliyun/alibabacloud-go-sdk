// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAutoShutdownPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAutoShutdownPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAutoShutdownPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *QueryAutoShutdownPoliciesResponseBody) *QueryAutoShutdownPoliciesResponse
	GetBody() *QueryAutoShutdownPoliciesResponseBody
}

type QueryAutoShutdownPoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAutoShutdownPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAutoShutdownPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAutoShutdownPoliciesResponse) GoString() string {
	return s.String()
}

func (s *QueryAutoShutdownPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAutoShutdownPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAutoShutdownPoliciesResponse) GetBody() *QueryAutoShutdownPoliciesResponseBody {
	return s.Body
}

func (s *QueryAutoShutdownPoliciesResponse) SetHeaders(v map[string]*string) *QueryAutoShutdownPoliciesResponse {
	s.Headers = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponse) SetStatusCode(v int32) *QueryAutoShutdownPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAutoShutdownPoliciesResponse) SetBody(v *QueryAutoShutdownPoliciesResponseBody) *QueryAutoShutdownPoliciesResponse {
	s.Body = v
	return s
}

func (s *QueryAutoShutdownPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
