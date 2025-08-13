// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoliciesForTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoliciesForTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListPoliciesForTargetResponseBody) *ListPoliciesForTargetResponse
	GetBody() *ListPoliciesForTargetResponseBody
}

type ListPoliciesForTargetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoliciesForTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoliciesForTargetResponse) GetBody() *ListPoliciesForTargetResponseBody {
	return s.Body
}

func (s *ListPoliciesForTargetResponse) SetHeaders(v map[string]*string) *ListPoliciesForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForTargetResponse) SetStatusCode(v int32) *ListPoliciesForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForTargetResponse) SetBody(v *ListPoliciesForTargetResponseBody) *ListPoliciesForTargetResponse {
	s.Body = v
	return s
}

func (s *ListPoliciesForTargetResponse) Validate() error {
	return dara.Validate(s)
}
