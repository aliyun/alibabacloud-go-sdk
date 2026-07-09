// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterDtoList(v []*QueryCostCenterResponseBodyCostCenterDtoList) *QueryCostCenterResponseBody
	GetCostCenterDtoList() []*QueryCostCenterResponseBodyCostCenterDtoList
	SetCurrentPage(v int32) *QueryCostCenterResponseBody
	GetCurrentPage() *int32
	SetMetadata(v interface{}) *QueryCostCenterResponseBody
	GetMetadata() interface{}
	SetPageSize(v int32) *QueryCostCenterResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryCostCenterResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryCostCenterResponseBody
	GetTotalCount() *int32
}

type QueryCostCenterResponseBody struct {
	// The list of financial unit data.
	CostCenterDtoList []*QueryCostCenterResponseBodyCostCenterDtoList `json:"CostCenterDtoList,omitempty" xml:"CostCenterDtoList,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The metadata of the response struct.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponseBody) GetCostCenterDtoList() []*QueryCostCenterResponseBodyCostCenterDtoList {
	return s.CostCenterDtoList
}

func (s *QueryCostCenterResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCostCenterResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryCostCenterResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostCenterResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryCostCenterResponseBody) SetCostCenterDtoList(v []*QueryCostCenterResponseBodyCostCenterDtoList) *QueryCostCenterResponseBody {
	s.CostCenterDtoList = v
	return s
}

func (s *QueryCostCenterResponseBody) SetCurrentPage(v int32) *QueryCostCenterResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetMetadata(v interface{}) *QueryCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterResponseBody) SetPageSize(v int32) *QueryCostCenterResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetRequestId(v string) *QueryCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetTotalCount(v int32) *QueryCostCenterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryCostCenterResponseBody) Validate() error {
	if s.CostCenterDtoList != nil {
		for _, item := range s.CostCenterDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCostCenterResponseBodyCostCenterDtoList struct {
	// The code of the financial unit.
	//
	// example:
	//
	// 15945703968#
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// The ID of the financial unit.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// The name of the financial unit. The name must be unique within the same account.
	//
	// example:
	//
	// 资源组
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// The level of the financial unit node.
	//
	// example:
	//
	// loose
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the user who owns the financial unit.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// The ID of the parent financial unit.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// The ID of the previous financial unit.
	//
	// example:
	//
	// 485996
	PrevCostCenterId *int64 `json:"PrevCostCenterId,omitempty" xml:"PrevCostCenterId,omitempty"`
	Priority         *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s QueryCostCenterResponseBodyCostCenterDtoList) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResponseBodyCostCenterDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetLevel() *int32 {
	return s.Level
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetPrevCostCenterId() *int64 {
	return s.PrevCostCenterId
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterCode(v string) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterName(v string) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetLevel(v int32) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.Level = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetOwnerAccountId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetParentCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetPrevCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.PrevCostCenterId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetPriority(v int32) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.Priority = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) Validate() error {
	return dara.Validate(s)
}
