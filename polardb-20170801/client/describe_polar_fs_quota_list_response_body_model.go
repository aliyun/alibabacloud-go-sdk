// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribePolarFsQuotaListResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribePolarFsQuotaListResponseBody
	GetPageRecordCount() *string
	SetPageSize(v string) *DescribePolarFsQuotaListResponseBody
	GetPageSize() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsQuotaListResponseBody
	GetPolarFsInstanceId() *string
	SetQuotaItems(v []*DescribePolarFsQuotaListResponseBodyQuotaItems) *DescribePolarFsQuotaListResponseBody
	GetQuotaItems() []*DescribePolarFsQuotaListResponseBodyQuotaItems
	SetRequestId(v string) *DescribePolarFsQuotaListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribePolarFsQuotaListResponseBody
	GetTotalRecordCount() *string
}

type DescribePolarFsQuotaListResponseBody struct {
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pfs-test****
	PolarFsInstanceId *string                                           `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	QuotaItems        []*DescribePolarFsQuotaListResponseBodyQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarFsQuotaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribePolarFsQuotaListResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribePolarFsQuotaListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePolarFsQuotaListResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsQuotaListResponseBody) GetQuotaItems() []*DescribePolarFsQuotaListResponseBodyQuotaItems {
	return s.QuotaItems
}

func (s *DescribePolarFsQuotaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsQuotaListResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribePolarFsQuotaListResponseBody) SetPageNumber(v string) *DescribePolarFsQuotaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetPageRecordCount(v string) *DescribePolarFsQuotaListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetPageSize(v string) *DescribePolarFsQuotaListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetPolarFsInstanceId(v string) *DescribePolarFsQuotaListResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetQuotaItems(v []*DescribePolarFsQuotaListResponseBodyQuotaItems) *DescribePolarFsQuotaListResponseBody {
	s.QuotaItems = v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetRequestId(v string) *DescribePolarFsQuotaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) SetTotalRecordCount(v string) *DescribePolarFsQuotaListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBody) Validate() error {
	if s.QuotaItems != nil {
		for _, item := range s.QuotaItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsQuotaListResponseBodyQuotaItems struct {
	// example:
	//
	// 1073741824
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// Inodes
	//
	// example:
	//
	// 100
	Inodes *int64 `json:"Inodes,omitempty" xml:"Inodes,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 104857600
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// example:
	//
	// 1
	UsedInodes *int64 `json:"UsedInodes,omitempty" xml:"UsedInodes,omitempty"`
}

func (s DescribePolarFsQuotaListResponseBodyQuotaItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaListResponseBodyQuotaItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) GetInodes() *int64 {
	return s.Inodes
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) GetUsedInodes() *int64 {
	return s.UsedInodes
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) SetCapacity(v int64) *DescribePolarFsQuotaListResponseBodyQuotaItems {
	s.Capacity = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) SetInodes(v int64) *DescribePolarFsQuotaListResponseBodyQuotaItems {
	s.Inodes = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) SetPath(v string) *DescribePolarFsQuotaListResponseBodyQuotaItems {
	s.Path = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) SetUsedCapacity(v int64) *DescribePolarFsQuotaListResponseBodyQuotaItems {
	s.UsedCapacity = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) SetUsedInodes(v int64) *DescribePolarFsQuotaListResponseBodyQuotaItems {
	s.UsedInodes = &v
	return s
}

func (s *DescribePolarFsQuotaListResponseBodyQuotaItems) Validate() error {
	return dara.Validate(s)
}
