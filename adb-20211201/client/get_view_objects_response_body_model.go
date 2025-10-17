// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetViewObjectsResponseBodyData) *GetViewObjectsResponseBody
	GetData() *GetViewObjectsResponseBodyData
	SetPageNumber(v int64) *GetViewObjectsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetViewObjectsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetViewObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetViewObjectsResponseBody
	GetTotalCount() *int64
}

type GetViewObjectsResponseBody struct {
	// The returned data.
	Data *GetViewObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: **1**.
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
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetViewObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetViewObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponseBody) GetData() *GetViewObjectsResponseBodyData {
	return s.Data
}

func (s *GetViewObjectsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetViewObjectsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetViewObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetViewObjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetViewObjectsResponseBody) SetData(v *GetViewObjectsResponseBodyData) *GetViewObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetViewObjectsResponseBody) SetPageNumber(v int64) *GetViewObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetPageSize(v int64) *GetViewObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetRequestId(v string) *GetViewObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViewObjectsResponseBody) SetTotalCount(v int64) *GetViewObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetViewObjectsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetViewObjectsResponseBodyData struct {
	// The page number. Pages start from page 1. Default value: **1**.
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
	// The queried views.
	TableSummaryModels []*TableSummaryModel `json:"TableSummaryModels,omitempty" xml:"TableSummaryModels,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetViewObjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetViewObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetViewObjectsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetViewObjectsResponseBodyData) GetTableSummaryModels() []*TableSummaryModel {
	return s.TableSummaryModels
}

func (s *GetViewObjectsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetViewObjectsResponseBodyData) SetPageNumber(v int64) *GetViewObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetPageSize(v int64) *GetViewObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetTableSummaryModels(v []*TableSummaryModel) *GetViewObjectsResponseBodyData {
	s.TableSummaryModels = v
	return s
}

func (s *GetViewObjectsResponseBodyData) SetTotalCount(v int64) *GetViewObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetViewObjectsResponseBodyData) Validate() error {
	if s.TableSummaryModels != nil {
		for _, item := range s.TableSummaryModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
