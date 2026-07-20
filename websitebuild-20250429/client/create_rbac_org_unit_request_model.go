// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacOrgUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateRbacOrgUnitRequest
	GetBizId() *string
	SetOrgUnitData(v string) *CreateRbacOrgUnitRequest
	GetOrgUnitData() *string
}

type CreateRbacOrgUnitRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrgUnitData *string `json:"OrgUnitData,omitempty" xml:"OrgUnitData,omitempty"`
}

func (s CreateRbacOrgUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacOrgUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateRbacOrgUnitRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateRbacOrgUnitRequest) GetOrgUnitData() *string {
	return s.OrgUnitData
}

func (s *CreateRbacOrgUnitRequest) SetBizId(v string) *CreateRbacOrgUnitRequest {
	s.BizId = &v
	return s
}

func (s *CreateRbacOrgUnitRequest) SetOrgUnitData(v string) *CreateRbacOrgUnitRequest {
	s.OrgUnitData = &v
	return s
}

func (s *CreateRbacOrgUnitRequest) Validate() error {
	return dara.Validate(s)
}
