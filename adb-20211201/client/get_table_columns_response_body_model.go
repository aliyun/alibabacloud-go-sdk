// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTableColumnsResponseBodyData) *GetTableColumnsResponseBody
	GetData() *GetTableColumnsResponseBodyData
	SetPageNumber(v int64) *GetTableColumnsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetTableColumnsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetTableColumnsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetTableColumnsResponseBody
	GetTotalCount() *int64
}

type GetTableColumnsResponseBody struct {
	// The returned data.
	Data *GetTableColumnsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBody) GetData() *GetTableColumnsResponseBodyData {
	return s.Data
}

func (s *GetTableColumnsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableColumnsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTableColumnsResponseBody) SetData(v *GetTableColumnsResponseBodyData) *GetTableColumnsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableColumnsResponseBody) SetPageNumber(v int64) *GetTableColumnsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetPageSize(v int64) *GetTableColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetRequestId(v string) *GetTableColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetTotalCount(v int64) *GetTableColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTableColumnsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableColumnsResponseBodyData struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the table.
	Table *TableDetailModel `json:"Table,omitempty" xml:"Table,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTableColumnsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableColumnsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableColumnsResponseBodyData) GetTable() *TableDetailModel {
	return s.Table
}

func (s *GetTableColumnsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTableColumnsResponseBodyData) SetPageNumber(v int64) *GetTableColumnsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetPageSize(v int64) *GetTableColumnsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetTable(v *TableDetailModel) *GetTableColumnsResponseBodyData {
	s.Table = v
	return s
}

func (s *GetTableColumnsResponseBodyData) SetTotalCount(v int64) *GetTableColumnsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetTableColumnsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
