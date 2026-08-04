// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListDeptMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterListDeptMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterListDeptMembersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterListDeptMembersResponseBody) *ModelRouterListDeptMembersResponse
	GetBody() *ModelRouterListDeptMembersResponseBody
}

type ModelRouterListDeptMembersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterListDeptMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterListDeptMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListDeptMembersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterListDeptMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterListDeptMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterListDeptMembersResponse) GetBody() *ModelRouterListDeptMembersResponseBody {
	return s.Body
}

func (s *ModelRouterListDeptMembersResponse) SetHeaders(v map[string]*string) *ModelRouterListDeptMembersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterListDeptMembersResponse) SetStatusCode(v int32) *ModelRouterListDeptMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterListDeptMembersResponse) SetBody(v *ModelRouterListDeptMembersResponseBody) *ModelRouterListDeptMembersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterListDeptMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
