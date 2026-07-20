// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacOrgUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteRbacOrgUnitRequest
	GetBizId() *string
	SetOrgUnitId(v string) *DeleteRbacOrgUnitRequest
	GetOrgUnitId() *string
}

type DeleteRbacOrgUnitRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrgUnitId *string `json:"OrgUnitId,omitempty" xml:"OrgUnitId,omitempty"`
}

func (s DeleteRbacOrgUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacOrgUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteRbacOrgUnitRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteRbacOrgUnitRequest) GetOrgUnitId() *string {
	return s.OrgUnitId
}

func (s *DeleteRbacOrgUnitRequest) SetBizId(v string) *DeleteRbacOrgUnitRequest {
	s.BizId = &v
	return s
}

func (s *DeleteRbacOrgUnitRequest) SetOrgUnitId(v string) *DeleteRbacOrgUnitRequest {
	s.OrgUnitId = &v
	return s
}

func (s *DeleteRbacOrgUnitRequest) Validate() error {
	return dara.Validate(s)
}
