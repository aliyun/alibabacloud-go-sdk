// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteRbacPermissionRequest
	GetBizId() *string
	SetPermissionId(v string) *DeleteRbacPermissionRequest
	GetPermissionId() *string
}

type DeleteRbacPermissionRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PermissionId *string `json:"PermissionId,omitempty" xml:"PermissionId,omitempty"`
}

func (s DeleteRbacPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacPermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRbacPermissionRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteRbacPermissionRequest) GetPermissionId() *string {
	return s.PermissionId
}

func (s *DeleteRbacPermissionRequest) SetBizId(v string) *DeleteRbacPermissionRequest {
	s.BizId = &v
	return s
}

func (s *DeleteRbacPermissionRequest) SetPermissionId(v string) *DeleteRbacPermissionRequest {
	s.PermissionId = &v
	return s
}

func (s *DeleteRbacPermissionRequest) Validate() error {
	return dara.Validate(s)
}
