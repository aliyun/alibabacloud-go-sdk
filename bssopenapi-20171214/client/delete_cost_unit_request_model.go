// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerUid(v int64) *DeleteCostUnitRequest
	GetOwnerUid() *int64
	SetUnitId(v int64) *DeleteCostUnitRequest
	GetUnitId() *int64
}

type DeleteCostUnitRequest struct {
	// The user ID of the cost center owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2135342
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the cost center. A value of -1 indicates the root cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 376348
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s DeleteCostUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitRequest) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *DeleteCostUnitRequest) GetUnitId() *int64 {
	return s.UnitId
}

func (s *DeleteCostUnitRequest) SetOwnerUid(v int64) *DeleteCostUnitRequest {
	s.OwnerUid = &v
	return s
}

func (s *DeleteCostUnitRequest) SetUnitId(v int64) *DeleteCostUnitRequest {
	s.UnitId = &v
	return s
}

func (s *DeleteCostUnitRequest) Validate() error {
	return dara.Validate(s)
}
