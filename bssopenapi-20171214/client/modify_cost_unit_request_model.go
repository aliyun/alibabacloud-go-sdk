// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUnitEntityList(v []*ModifyCostUnitRequestUnitEntityList) *ModifyCostUnitRequest
	GetUnitEntityList() []*ModifyCostUnitRequestUnitEntityList
}

type ModifyCostUnitRequest struct {
	// The cost centers to be modified.
	UnitEntityList []*ModifyCostUnitRequestUnitEntityList `json:"UnitEntityList,omitempty" xml:"UnitEntityList,omitempty" type:"Repeated"`
}

func (s ModifyCostUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostUnitRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitRequest) GetUnitEntityList() []*ModifyCostUnitRequestUnitEntityList {
	return s.UnitEntityList
}

func (s *ModifyCostUnitRequest) SetUnitEntityList(v []*ModifyCostUnitRequestUnitEntityList) *ModifyCostUnitRequest {
	s.UnitEntityList = v
	return s
}

func (s *ModifyCostUnitRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyCostUnitRequestUnitEntityList struct {
	// The new name of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// newTest
	NewUnitName *string `json:"NewUnitName,omitempty" xml:"NewUnitName,omitempty"`
	// The user ID of the cost center owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1321312312
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2524352
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s ModifyCostUnitRequestUnitEntityList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostUnitRequestUnitEntityList) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitRequestUnitEntityList) GetNewUnitName() *string {
	return s.NewUnitName
}

func (s *ModifyCostUnitRequestUnitEntityList) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *ModifyCostUnitRequestUnitEntityList) GetUnitId() *int64 {
	return s.UnitId
}

func (s *ModifyCostUnitRequestUnitEntityList) SetNewUnitName(v string) *ModifyCostUnitRequestUnitEntityList {
	s.NewUnitName = &v
	return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *ModifyCostUnitRequestUnitEntityList {
	s.OwnerUid = &v
	return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetUnitId(v int64) *ModifyCostUnitRequestUnitEntityList {
	s.UnitId = &v
	return s
}

func (s *ModifyCostUnitRequestUnitEntityList) Validate() error {
	return dara.Validate(s)
}
