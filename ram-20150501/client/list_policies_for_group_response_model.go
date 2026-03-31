// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoliciesForGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoliciesForGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListPoliciesForGroupResponseBody) *ListPoliciesForGroupResponse
	GetBody() *ListPoliciesForGroupResponseBody
}

type ListPoliciesForGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoliciesForGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoliciesForGroupResponse) GetBody() *ListPoliciesForGroupResponseBody {
	return s.Body
}

func (s *ListPoliciesForGroupResponse) SetHeaders(v map[string]*string) *ListPoliciesForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForGroupResponse) SetStatusCode(v int32) *ListPoliciesForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForGroupResponse) SetBody(v *ListPoliciesForGroupResponseBody) *ListPoliciesForGroupResponse {
	s.Body = v
	return s
}

func (s *ListPoliciesForGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
