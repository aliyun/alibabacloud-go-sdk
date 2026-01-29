// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUnitEntityList(v []*CreateCostUnitRequestUnitEntityList) *CreateCostUnitRequest
	GetUnitEntityList() []*CreateCostUnitRequestUnitEntityList
}

type CreateCostUnitRequest struct {
	// The list of cost centers.
	UnitEntityList []*CreateCostUnitRequestUnitEntityList `json:"UnitEntityList,omitempty" xml:"UnitEntityList,omitempty" type:"Repeated"`
}

func (s CreateCostUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateCostUnitRequest) GetUnitEntityList() []*CreateCostUnitRequestUnitEntityList {
	return s.UnitEntityList
}

func (s *CreateCostUnitRequest) SetUnitEntityList(v []*CreateCostUnitRequestUnitEntityList) *CreateCostUnitRequest {
	s.UnitEntityList = v
	return s
}

func (s *CreateCostUnitRequest) Validate() error {
	if s.UnitEntityList != nil {
		for _, item := range s.UnitEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCostUnitRequestUnitEntityList struct {
	// The user ID of the owner of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 982375623
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the parent cost center. A value of -1 indicates the root cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	// The name of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	UnitName *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s CreateCostUnitRequestUnitEntityList) String() string {
	return dara.Prettify(s)
}

func (s CreateCostUnitRequestUnitEntityList) GoString() string {
	return s.String()
}

func (s *CreateCostUnitRequestUnitEntityList) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *CreateCostUnitRequestUnitEntityList) GetParentUnitId() *int64 {
	return s.ParentUnitId
}

func (s *CreateCostUnitRequestUnitEntityList) GetUnitName() *string {
	return s.UnitName
}

func (s *CreateCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *CreateCostUnitRequestUnitEntityList {
	s.OwnerUid = &v
	return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetParentUnitId(v int64) *CreateCostUnitRequestUnitEntityList {
	s.ParentUnitId = &v
	return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetUnitName(v string) *CreateCostUnitRequestUnitEntityList {
	s.UnitName = &v
	return s
}

func (s *CreateCostUnitRequestUnitEntityList) Validate() error {
	return dara.Validate(s)
}
