// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateRbacPermissionRequest
	GetBizId() *string
	SetPermissionData(v string) *CreateRbacPermissionRequest
	GetPermissionData() *string
}

type CreateRbacPermissionRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PermissionData *string `json:"PermissionData,omitempty" xml:"PermissionData,omitempty"`
}

func (s CreateRbacPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacPermissionRequest) GoString() string {
	return s.String()
}

func (s *CreateRbacPermissionRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateRbacPermissionRequest) GetPermissionData() *string {
	return s.PermissionData
}

func (s *CreateRbacPermissionRequest) SetBizId(v string) *CreateRbacPermissionRequest {
	s.BizId = &v
	return s
}

func (s *CreateRbacPermissionRequest) SetPermissionData(v string) *CreateRbacPermissionRequest {
	s.PermissionData = &v
	return s
}

func (s *CreateRbacPermissionRequest) Validate() error {
	return dara.Validate(s)
}
