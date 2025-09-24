// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerUid(v int64) *QueryCostUnitResourceRequest
	GetOwnerUid() *int64
	SetPageNum(v int32) *QueryCostUnitResourceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryCostUnitResourceRequest
	GetPageSize() *int32
	SetUnitId(v int64) *QueryCostUnitResourceRequest
	GetUnitId() *int64
}

type QueryCostUnitResourceRequest struct {
	// The user ID of the cost center owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23453245
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 235325
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s QueryCostUnitResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceRequest) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *QueryCostUnitResourceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryCostUnitResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostUnitResourceRequest) GetUnitId() *int64 {
	return s.UnitId
}

func (s *QueryCostUnitResourceRequest) SetOwnerUid(v int64) *QueryCostUnitResourceRequest {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetPageNum(v int32) *QueryCostUnitResourceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetPageSize(v int32) *QueryCostUnitResourceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetUnitId(v int64) *QueryCostUnitResourceRequest {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResourceRequest) Validate() error {
	return dara.Validate(s)
}
