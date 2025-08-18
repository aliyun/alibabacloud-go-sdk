// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManagedRulesGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManagedRulesGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListManagedRulesGroupsResponseBody) *ListManagedRulesGroupsResponse
	GetBody() *ListManagedRulesGroupsResponseBody
}

type ListManagedRulesGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManagedRulesGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManagedRulesGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListManagedRulesGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManagedRulesGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManagedRulesGroupsResponse) GetBody() *ListManagedRulesGroupsResponseBody {
	return s.Body
}

func (s *ListManagedRulesGroupsResponse) SetHeaders(v map[string]*string) *ListManagedRulesGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListManagedRulesGroupsResponse) SetStatusCode(v int32) *ListManagedRulesGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedRulesGroupsResponse) SetBody(v *ListManagedRulesGroupsResponseBody) *ListManagedRulesGroupsResponse {
	s.Body = v
	return s
}

func (s *ListManagedRulesGroupsResponse) Validate() error {
	return dara.Validate(s)
}
