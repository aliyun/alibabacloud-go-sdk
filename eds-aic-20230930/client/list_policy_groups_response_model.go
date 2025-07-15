// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyGroupsResponseBody) *ListPolicyGroupsResponse
	GetBody() *ListPolicyGroupsResponseBody
}

type ListPolicyGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyGroupsResponse) GetBody() *ListPolicyGroupsResponseBody {
	return s.Body
}

func (s *ListPolicyGroupsResponse) SetHeaders(v map[string]*string) *ListPolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyGroupsResponse) SetStatusCode(v int32) *ListPolicyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyGroupsResponse) SetBody(v *ListPolicyGroupsResponseBody) *ListPolicyGroupsResponse {
	s.Body = v
	return s
}

func (s *ListPolicyGroupsResponse) Validate() error {
	return dara.Validate(s)
}
