// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerUid(v int64) *QueryCostUnitRequest
	GetOwnerUid() *int64
	SetPageNum(v int32) *QueryCostUnitRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryCostUnitRequest
	GetPageSize() *int32
	SetParentUnitId(v int64) *QueryCostUnitRequest
	GetParentUnitId() *int64
}

type QueryCostUnitRequest struct {
	// The owner of the financial unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28394563429587
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent financial unit. A value of -1 indicates the root of the financial unit tree.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
}

func (s QueryCostUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitRequest) GoString() string {
	return s.String()
}

func (s *QueryCostUnitRequest) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *QueryCostUnitRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryCostUnitRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostUnitRequest) GetParentUnitId() *int64 {
	return s.ParentUnitId
}

func (s *QueryCostUnitRequest) SetOwnerUid(v int64) *QueryCostUnitRequest {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitRequest) SetPageNum(v int32) *QueryCostUnitRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitRequest) SetPageSize(v int32) *QueryCostUnitRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitRequest) SetParentUnitId(v int64) *QueryCostUnitRequest {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitRequest) Validate() error {
	return dara.Validate(s)
}
