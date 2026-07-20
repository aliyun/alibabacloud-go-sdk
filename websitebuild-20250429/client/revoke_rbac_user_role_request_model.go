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
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
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
