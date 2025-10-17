// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDatabaseObjectsResponseBodyData) *GetDatabaseObjectsResponseBody
	GetData() *GetDatabaseObjectsResponseBodyData
	SetPageNumber(v int64) *GetDatabaseObjectsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetDatabaseObjectsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetDatabaseObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetDatabaseObjectsResponseBody
	GetTotalCount() *int64
}

type GetDatabaseObjectsResponseBody struct {
	// The returned data.
	Data *GetDatabaseObjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 863D51B7-5321-41D8-A0B6-A088B0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDatabaseObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponseBody) GetData() *GetDatabaseObjectsResponseBodyData {
	return s.Data
}

func (s *GetDatabaseObjectsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetDatabaseObjectsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDatabaseObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseObjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDatabaseObjectsResponseBody) SetData(v *GetDatabaseObjectsResponseBodyData) *GetDatabaseObjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetPageNumber(v int64) *GetDatabaseObjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetPageSize(v int64) *GetDatabaseObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetRequestId(v string) *GetDatabaseObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) SetTotalCount(v int64) *GetDatabaseObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDatabaseObjectsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseObjectsResponseBodyData struct {
	// The queried databases.
	DatabaseSummaryModels []*DatabaseSummaryModel `json:"DatabaseSummaryModels,omitempty" xml:"DatabaseSummaryModels,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDatabaseObjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsResponseBodyData) GetDatabaseSummaryModels() []*DatabaseSummaryModel {
	return s.DatabaseSummaryModels
}

func (s *GetDatabaseObjectsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetDatabaseObjectsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDatabaseObjectsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDatabaseObjectsResponseBodyData) SetDatabaseSummaryModels(v []*DatabaseSummaryModel) *GetDatabaseObjectsResponseBodyData {
	s.DatabaseSummaryModels = v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetPageNumber(v int64) *GetDatabaseObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetPageSize(v int64) *GetDatabaseObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) SetTotalCount(v int64) *GetDatabaseObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetDatabaseObjectsResponseBodyData) Validate() error {
	if s.DatabaseSummaryModels != nil {
		for _, item := range s.DatabaseSummaryModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
