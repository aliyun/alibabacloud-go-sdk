// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserLoginPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserLoginPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserLoginPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserLoginPoliciesResponseBody) *RemoveUserLoginPoliciesResponse
	GetBody() *RemoveUserLoginPoliciesResponseBody
}

type RemoveUserLoginPoliciesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserLoginPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserLoginPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserLoginPoliciesResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserLoginPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserLoginPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserLoginPoliciesResponse) GetBody() *RemoveUserLoginPoliciesResponseBody {
	return s.Body
}

func (s *RemoveUserLoginPoliciesResponse) SetHeaders(v map[string]*string) *RemoveUserLoginPoliciesResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserLoginPoliciesResponse) SetStatusCode(v int32) *RemoveUserLoginPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserLoginPoliciesResponse) SetBody(v *RemoveUserLoginPoliciesResponseBody) *RemoveUserLoginPoliciesResponse {
	s.Body = v
	return s
}

func (s *RemoveUserLoginPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
