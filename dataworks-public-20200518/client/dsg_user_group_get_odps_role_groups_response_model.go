// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupGetOdpsRoleGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgUserGroupGetOdpsRoleGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgUserGroupGetOdpsRoleGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DsgUserGroupGetOdpsRoleGroupsResponseBody) *DsgUserGroupGetOdpsRoleGroupsResponse
	GetBody() *DsgUserGroupGetOdpsRoleGroupsResponseBody
}

type DsgUserGroupGetOdpsRoleGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgUserGroupGetOdpsRoleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgUserGroupGetOdpsRoleGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupGetOdpsRoleGroupsResponse) GoString() string {
	return s.String()
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) GetBody() *DsgUserGroupGetOdpsRoleGroupsResponseBody {
	return s.Body
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) SetHeaders(v map[string]*string) *DsgUserGroupGetOdpsRoleGroupsResponse {
	s.Headers = v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) SetStatusCode(v int32) *DsgUserGroupGetOdpsRoleGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) SetBody(v *DsgUserGroupGetOdpsRoleGroupsResponseBody) *DsgUserGroupGetOdpsRoleGroupsResponse {
	s.Body = v
	return s
}

func (s *DsgUserGroupGetOdpsRoleGroupsResponse) Validate() error {
	return dara.Validate(s)
}
