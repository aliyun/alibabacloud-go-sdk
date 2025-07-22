// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterEntityList(v []*CreateCostCenterRequestCostCenterEntityList) *CreateCostCenterRequest
	GetCostCenterEntityList() []*CreateCostCenterRequestCostCenterEntityList
	SetNbid(v string) *CreateCostCenterRequest
	GetNbid() *string
}

type CreateCostCenterRequest struct {
	// This parameter is required.
	CostCenterEntityList []*CreateCostCenterRequestCostCenterEntityList `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty" type:"Repeated"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRequest) GetCostCenterEntityList() []*CreateCostCenterRequestCostCenterEntityList {
	return s.CostCenterEntityList
}

func (s *CreateCostCenterRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateCostCenterRequest) SetCostCenterEntityList(v []*CreateCostCenterRequestCostCenterEntityList) *CreateCostCenterRequest {
	s.CostCenterEntityList = v
	return s
}

func (s *CreateCostCenterRequest) SetNbid(v string) *CreateCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *CreateCostCenterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCostCenterRequestCostCenterEntityList struct {
	// This parameter is required.
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s CreateCostCenterRequestCostCenterEntityList) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRequestCostCenterEntityList) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRequestCostCenterEntityList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *CreateCostCenterRequestCostCenterEntityList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *CreateCostCenterRequestCostCenterEntityList) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetCostCenterName(v string) *CreateCostCenterRequestCostCenterEntityList {
	s.CostCenterName = &v
	return s
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetOwnerAccountId(v int64) *CreateCostCenterRequestCostCenterEntityList {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetParentCostCenterId(v int64) *CreateCostCenterRequestCostCenterEntityList {
	s.ParentCostCenterId = &v
	return s
}

func (s *CreateCostCenterRequestCostCenterEntityList) Validate() error {
	return dara.Validate(s)
}
