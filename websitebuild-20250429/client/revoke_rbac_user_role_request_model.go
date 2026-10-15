// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRbacUserRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RevokeRbacUserRoleRequest
	GetBizId() *string
	SetUserRoleData(v string) *RevokeRbacUserRoleRequest
	GetUserRoleData() *string
}

type RevokeRbacUserRoleRequest struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The user role information.
	//
	// example:
	//
	// {
	//
	//         "org_id": "f0534aa7-94ec-4368-a34a-310390f66253",
	//
	//         "role_id": "804a6b2a-82ac-49f2-960a-a6c2ed9e73dd",
	//
	//         "user_id": "c1721c87-5401-40b3-b836-9b2653507feb"
	//
	//       }
	UserRoleData *string `json:"UserRoleData,omitempty" xml:"UserRoleData,omitempty"`
}

func (s RevokeRbacUserRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRbacUserRoleRequest) GoString() string {
	return s.String()
}

func (s *RevokeRbacUserRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *RevokeRbacUserRoleRequest) GetUserRoleData() *string {
	return s.UserRoleData
}

func (s *RevokeRbacUserRoleRequest) SetBizId(v string) *RevokeRbacUserRoleRequest {
	s.BizId = &v
	return s
}

func (s *RevokeRbacUserRoleRequest) SetUserRoleData(v string) *RevokeRbacUserRoleRequest {
	s.UserRoleData = &v
	return s
}

func (s *RevokeRbacUserRoleRequest) Validate() error {
	return dara.Validate(s)
}
