// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterEntityList(v []*ModifyCostCenterRequestCostCenterEntityList) *ModifyCostCenterRequest
	GetCostCenterEntityList() []*ModifyCostCenterRequestCostCenterEntityList
	SetNbid(v string) *ModifyCostCenterRequest
	GetNbid() *string
}

type ModifyCostCenterRequest struct {
	// The list of cost center entities.
	//
	// This parameter is required.
	CostCenterEntityList []*ModifyCostCenterRequestCostCenterEntityList `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty" type:"Repeated"`
	// The level-1 marketplace ID. If this parameter is left empty, the marketplace ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ModifyCostCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRequest) GetCostCenterEntityList() []*ModifyCostCenterRequestCostCenterEntityList {
	return s.CostCenterEntityList
}

func (s *ModifyCostCenterRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ModifyCostCenterRequest) SetCostCenterEntityList(v []*ModifyCostCenterRequestCostCenterEntityList) *ModifyCostCenterRequest {
	s.CostCenterEntityList = v
	return s
}

func (s *ModifyCostCenterRequest) SetNbid(v string) *ModifyCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *ModifyCostCenterRequest) Validate() error {
	if s.CostCenterEntityList != nil {
		for _, item := range s.CostCenterEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCostCenterRequestCostCenterEntityList struct {
	// The ID of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// The new name of the cost center. The name must be unique within the same account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 资源组
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// The ID of the owner to which the cost center belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterRequestCostCenterEntityList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRequestCostCenterEntityList) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRequestCostCenterEntityList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ModifyCostCenterRequestCostCenterEntityList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ModifyCostCenterRequestCostCenterEntityList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetCostCenterId(v int64) *ModifyCostCenterRequestCostCenterEntityList {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetCostCenterName(v string) *ModifyCostCenterRequestCostCenterEntityList {
	s.CostCenterName = &v
	return s
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetOwnerAccountId(v int64) *ModifyCostCenterRequestCostCenterEntityList {
	s.OwnerAccountId = &v
	return s
}

func (s *ModifyCostCenterRequestCostCenterEntityList) Validate() error {
	return dara.Validate(s)
}
