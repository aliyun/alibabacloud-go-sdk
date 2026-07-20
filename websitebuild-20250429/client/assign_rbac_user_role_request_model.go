// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignRbacUserRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AssignRbacUserRoleRequest
	GetBizId() *string
	SetUserRoleData(v string) *AssignRbacUserRoleRequest
	GetUserRoleData() *string
}

type AssignRbacUserRoleRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UserRoleData *string `json:"UserRoleData,omitempty" xml:"UserRoleData,omitempty"`
}

func (s AssignRbacUserRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignRbacUserRoleRequest) GoString() string {
	return s.String()
}

func (s *AssignRbacUserRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *AssignRbacUserRoleRequest) GetUserRoleData() *string {
	return s.UserRoleData
}

func (s *AssignRbacUserRoleRequest) SetBizId(v string) *AssignRbacUserRoleRequest {
	s.BizId = &v
	return s
}

func (s *AssignRbacUserRoleRequest) SetUserRoleData(v string) *AssignRbacUserRoleRequest {
	s.UserRoleData = &v
	return s
}

func (s *AssignRbacUserRoleRequest) Validate() error {
	return dara.Validate(s)
}
