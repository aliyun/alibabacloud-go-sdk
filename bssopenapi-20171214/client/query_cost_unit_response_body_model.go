// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCostUnitResponseBody
	GetCode() *string
	SetData(v *QueryCostUnitResponseBodyData) *QueryCostUnitResponseBody
	GetData() *QueryCostUnitResponseBodyData
	SetMessage(v string) *QueryCostUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCostUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCostUnitResponseBody
	GetSuccess() *bool
}

type QueryCostUnitResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCostUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCostUnitResponseBody) GetData() *QueryCostUnitResponseBodyData {
	return s.Data
}

func (s *QueryCostUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCostUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCostUnitResponseBody) SetCode(v string) *QueryCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetData(v *QueryCostUnitResponseBodyData) *QueryCostUnitResponseBody {
	s.Data = v
	return s
}

func (s *QueryCostUnitResponseBody) SetMessage(v string) *QueryCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetRequestId(v string) *QueryCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetSuccess(v bool) *QueryCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCostUnitResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCostUnitResponseBodyData struct {
	// The cost centers.
	CostUnitDtoList []*QueryCostUnitResponseBodyDataCostUnitDtoList `json:"CostUnitDtoList,omitempty" xml:"CostUnitDtoList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostUnitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBodyData) GetCostUnitDtoList() []*QueryCostUnitResponseBodyDataCostUnitDtoList {
	return s.CostUnitDtoList
}

func (s *QueryCostUnitResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryCostUnitResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostUnitResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryCostUnitResponseBodyData) SetCostUnitDtoList(v []*QueryCostUnitResponseBodyDataCostUnitDtoList) *QueryCostUnitResponseBodyData {
	s.CostUnitDtoList = v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetPageNum(v int32) *QueryCostUnitResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetPageSize(v int32) *QueryCostUnitResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetTotalCount(v int32) *QueryCostUnitResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryCostUnitResponseBodyDataCostUnitDtoList struct {
	// The user ID of the cost center owner.
	//
	// example:
	//
	// 2343464
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the parent cost center. A value of -1 indicates the root cost center.
	//
	// example:
	//
	// -1
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	// The ID of the cost center.
	//
	// example:
	//
	// 23534
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	// The name of the cost center.
	//
	// example:
	//
	// test
	UnitName *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s QueryCostUnitResponseBodyDataCostUnitDtoList) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResponseBodyDataCostUnitDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) GetParentUnitId() *int64 {
	return s.ParentUnitId
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) GetUnitId() *int64 {
	return s.UnitId
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) GetUnitName() *string {
	return s.UnitName
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetOwnerUid(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetParentUnitId(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetUnitId(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetUnitName(v string) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitName = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) Validate() error {
	return dara.Validate(s)
}
