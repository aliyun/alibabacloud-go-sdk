// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProhibitedPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProhibitedPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProhibitedPoliciesResponseBody) *DeleteProhibitedPoliciesResponse
	GetBody() *DeleteProhibitedPoliciesResponseBody
}

type DeleteProhibitedPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProhibitedPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProhibitedPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProhibitedPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProhibitedPoliciesResponse) GetBody() *DeleteProhibitedPoliciesResponseBody {
	return s.Body
}

func (s *DeleteProhibitedPoliciesResponse) SetHeaders(v map[string]*string) *DeleteProhibitedPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteProhibitedPoliciesResponse) SetStatusCode(v int32) *DeleteProhibitedPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProhibitedPoliciesResponse) SetBody(v *DeleteProhibitedPoliciesResponseBody) *DeleteProhibitedPoliciesResponse {
	s.Body = v
	return s
}

func (s *DeleteProhibitedPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
