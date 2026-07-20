// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacOrgUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateRbacOrgUnitRequest
	GetBizId() *string
	SetOrgUnitData(v string) *UpdateRbacOrgUnitRequest
	GetOrgUnitData() *string
	SetOrgUnitId(v string) *UpdateRbacOrgUnitRequest
	GetOrgUnitId() *string
}

type UpdateRbacOrgUnitRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrgUnitData *string `json:"OrgUnitData,omitempty" xml:"OrgUnitData,omitempty"`
	OrgUnitId   *string `json:"OrgUnitId,omitempty" xml:"OrgUnitId,omitempty"`
}

func (s UpdateRbacOrgUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacOrgUnitRequest) GoString() string {
	return s.String()
}

func (s *UpdateRbacOrgUnitRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateRbacOrgUnitRequest) GetOrgUnitData() *string {
	return s.OrgUnitData
}

func (s *UpdateRbacOrgUnitRequest) GetOrgUnitId() *string {
	return s.OrgUnitId
}

func (s *UpdateRbacOrgUnitRequest) SetBizId(v string) *UpdateRbacOrgUnitRequest {
	s.BizId = &v
	return s
}

func (s *UpdateRbacOrgUnitRequest) SetOrgUnitData(v string) *UpdateRbacOrgUnitRequest {
	s.OrgUnitData = &v
	return s
}

func (s *UpdateRbacOrgUnitRequest) SetOrgUnitId(v string) *UpdateRbacOrgUnitRequest {
	s.OrgUnitId = &v
	return s
}

func (s *UpdateRbacOrgUnitRequest) Validate() error {
	return dara.Validate(s)
}
