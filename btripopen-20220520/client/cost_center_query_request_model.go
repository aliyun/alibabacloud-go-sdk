// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisable(v int64) *CostCenterQueryRequest
	GetDisable() *int64
	SetNeedOrgEntity(v bool) *CostCenterQueryRequest
	GetNeedOrgEntity() *bool
	SetThirdpartId(v string) *CostCenterQueryRequest
	GetThirdpartId() *string
	SetTitle(v string) *CostCenterQueryRequest
	GetTitle() *string
	SetUserId(v string) *CostCenterQueryRequest
	GetUserId() *string
}

type CostCenterQueryRequest struct {
	Disable *int64 `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// false
	NeedOrgEntity *bool `json:"need_org_entity,omitempty" xml:"need_org_entity,omitempty"`
	// example:
	//
	// cost1
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// user1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CostCenterQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryRequest) GoString() string {
	return s.String()
}

func (s *CostCenterQueryRequest) GetDisable() *int64 {
	return s.Disable
}

func (s *CostCenterQueryRequest) GetNeedOrgEntity() *bool {
	return s.NeedOrgEntity
}

func (s *CostCenterQueryRequest) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CostCenterQueryRequest) GetTitle() *string {
	return s.Title
}

func (s *CostCenterQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *CostCenterQueryRequest) SetDisable(v int64) *CostCenterQueryRequest {
	s.Disable = &v
	return s
}

func (s *CostCenterQueryRequest) SetNeedOrgEntity(v bool) *CostCenterQueryRequest {
	s.NeedOrgEntity = &v
	return s
}

func (s *CostCenterQueryRequest) SetThirdpartId(v string) *CostCenterQueryRequest {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterQueryRequest) SetTitle(v string) *CostCenterQueryRequest {
	s.Title = &v
	return s
}

func (s *CostCenterQueryRequest) SetUserId(v string) *CostCenterQueryRequest {
	s.UserId = &v
	return s
}

func (s *CostCenterQueryRequest) Validate() error {
	return dara.Validate(s)
}
