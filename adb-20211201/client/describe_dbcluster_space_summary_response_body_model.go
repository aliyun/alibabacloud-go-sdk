// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSpaceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBClusterSpaceSummaryResponseBodyData) *DescribeDBClusterSpaceSummaryResponseBody
	GetData() *DescribeDBClusterSpaceSummaryResponseBodyData
	SetRequestId(v string) *DescribeDBClusterSpaceSummaryResponseBody
	GetRequestId() *string
}

type DescribeDBClusterSpaceSummaryResponseBody struct {
	// The queried storage overview information.
	Data *DescribeDBClusterSpaceSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) GetData() *DescribeDBClusterSpaceSummaryResponseBodyData {
	return s.Data
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) SetData(v *DescribeDBClusterSpaceSummaryResponseBodyData) *DescribeDBClusterSpaceSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) SetRequestId(v string) *DescribeDBClusterSpaceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterSpaceSummaryResponseBodyData struct {
	// The cold data.
	ColdData *DescribeDBClusterSpaceSummaryResponseBodyDataColdData `json:"ColdData,omitempty" xml:"ColdData,omitempty" type:"Struct"`
	// The data growth.
	DataGrowth *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth `json:"DataGrowth,omitempty" xml:"DataGrowth,omitempty" type:"Struct"`
	// The hot data.
	HotData *DescribeDBClusterSpaceSummaryResponseBodyDataHotData `json:"HotData,omitempty" xml:"HotData,omitempty" type:"Struct"`
	// The total data size. Unit: bytes.
	//
	// >  Formula: Total data size = Hot data size+ Cold data size.
	//
	// example:
	//
	// 8388608
	TotalSize *string `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) GetColdData() *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	return s.ColdData
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) GetDataGrowth() *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth {
	return s.DataGrowth
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) GetHotData() *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	return s.HotData
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetColdData(v *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.ColdData = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetDataGrowth(v *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.DataGrowth = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetHotData(v *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.HotData = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) SetTotalSize(v string) *DescribeDBClusterSpaceSummaryResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyData) Validate() error {
	if s.ColdData != nil {
		if err := s.ColdData.Validate(); err != nil {
			return err
		}
	}
	if s.DataGrowth != nil {
		if err := s.DataGrowth.Validate(); err != nil {
			return err
		}
	}
	if s.HotData != nil {
		if err := s.HotData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterSpaceSummaryResponseBodyDataColdData struct {
	// The data size of table records. Unit: bytes.
	//
	// example:
	//
	// 1048576
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data size of regular indexes. Unit: bytes.
	//
	// example:
	//
	// 1048576
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The data size of other data. Unit: bytes.
	//
	// example:
	//
	// 1048576
	OtherSize *int64 `json:"OtherSize,omitempty" xml:"OtherSize,omitempty"`
	// The data size of primary key indexes. Unit: bytes.
	//
	// example:
	//
	// 1048576
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	// The cold data size. Unit: bytes.
	//
	// >  Formula: Cold data size = Data size of table records + Data size of regular indexes + Data size of primary key indexes + Data size of other data.
	//
	// example:
	//
	// 4194304
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataColdData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GetOtherSize() *int64 {
	return s.OtherSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GetPrimaryKeyIndexSize() *int64 {
	return s.PrimaryKeyIndexSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetDataSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.DataSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetOtherSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.OtherSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetPrimaryKeyIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) SetTotalSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataColdData {
	s.TotalSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataColdData) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth struct {
	// The data growth within the last day. Unit: bytes.
	//
	// >  Formula: Data growth within the last day = Current data size - Data size one day ago.
	//
	// example:
	//
	// 1048576
	DayGrowth *int64 `json:"DayGrowth,omitempty" xml:"DayGrowth,omitempty"`
	// The daily data growth within the last seven days. Unit: bytes.
	//
	// >  Formula: Daily data growth within the last seven days = (Current data size - Data size seven days ago)/7.
	//
	// example:
	//
	// 1048576
	WeekGrowth *int64 `json:"WeekGrowth,omitempty" xml:"WeekGrowth,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) GetDayGrowth() *int64 {
	return s.DayGrowth
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) GetWeekGrowth() *int64 {
	return s.WeekGrowth
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) SetDayGrowth(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth {
	s.DayGrowth = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) SetWeekGrowth(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth {
	s.WeekGrowth = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataDataGrowth) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterSpaceSummaryResponseBodyDataHotData struct {
	// The data size of table records. Unit: bytes.
	//
	// example:
	//
	// 1048576
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data size of regular indexes. Unit: bytes.
	//
	// example:
	//
	// 1048576
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The data size of other data. Unit: bytes.
	//
	// example:
	//
	// 1048576
	OtherSize *int64 `json:"OtherSize,omitempty" xml:"OtherSize,omitempty"`
	// The data size of primary key indexes. Unit: bytes.
	//
	// example:
	//
	// 1048576
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	// The hot data size. Unit: bytes.
	//
	// >  Formula: Hot data size = Data size of table records + Data size of regular indexes + Data size of primary key indexes + Data size of other data.
	//
	// example:
	//
	// 4194304
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataHotData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GetOtherSize() *int64 {
	return s.OtherSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GetPrimaryKeyIndexSize() *int64 {
	return s.PrimaryKeyIndexSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetDataSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.DataSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetOtherSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.OtherSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetPrimaryKeyIndexSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) SetTotalSize(v int64) *DescribeDBClusterSpaceSummaryResponseBodyDataHotData {
	s.TotalSize = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponseBodyDataHotData) Validate() error {
	return dara.Validate(s)
}
