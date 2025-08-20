// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTableObjectsResponseBodyData) *GetTableObjectsResponseBody
	GetData() *GetTableObjectsResponseBodyData
	SetPageNumber(v int64) *GetTableObjectsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetTableObjectsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetTableObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetTableObjectsResponseBody
	GetTotalCount() *int64
}

type GetTableObjectsResponseBody struct {
	// The data returned.
	Data *GetTableObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of the returned page. The value is an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 863D51B7-5321-41D8-A0B6-A088B0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponseBody) GetData() *GetTableObjectsResponseBodyData {
	return s.Data
}

func (s *GetTableObjectsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableObjectsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableObjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTableObjectsResponseBody) SetData(v *GetTableObjectsResponseBodyData) *GetTableObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableObjectsResponseBody) SetPageNumber(v int64) *GetTableObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetPageSize(v int64) *GetTableObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetRequestId(v string) *GetTableObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableObjectsResponseBody) SetTotalCount(v int64) *GetTableObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTableObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableObjectsResponseBodyData struct {
	// The number of the returned page. The value is an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details of the tables.
	TableSummaryModels []*TableSummaryModel `json:"TableSummaryModels,omitempty" xml:"TableSummaryModels,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableObjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableObjectsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableObjectsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableObjectsResponseBodyData) GetTableSummaryModels() []*TableSummaryModel {
	return s.TableSummaryModels
}

func (s *GetTableObjectsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTableObjectsResponseBodyData) SetPageNumber(v int64) *GetTableObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetPageSize(v int64) *GetTableObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetTableSummaryModels(v []*TableSummaryModel) *GetTableObjectsResponseBodyData {
	s.TableSummaryModels = v
	return s
}

func (s *GetTableObjectsResponseBodyData) SetTotalCount(v int64) *GetTableObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetTableObjectsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
