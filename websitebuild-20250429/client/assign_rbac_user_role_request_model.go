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
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The role parameters.
	//
	// example:
	//
	// {"user_id": "c1721c87-5401-40b3-b836-9b2653507feb",
	//
	//     "role_id": "804a6b2a-82ac-49f2-960a-a6c2ed9e73dd",
	//
	//     "org_id": "f0534aa7-94ec-4368-a34a-310390f66253"}
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
