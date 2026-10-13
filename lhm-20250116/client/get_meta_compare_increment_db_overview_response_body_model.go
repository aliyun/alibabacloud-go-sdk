// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCompareIncrementDbOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaCompareIncrementDbOverviewResponseBodyData) *GetMetaCompareIncrementDbOverviewResponseBody
	GetData() *GetMetaCompareIncrementDbOverviewResponseBodyData
	SetErrCode(v string) *GetMetaCompareIncrementDbOverviewResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetMetaCompareIncrementDbOverviewResponseBody
	GetErrMessage() *string
	SetSuccess(v bool) *GetMetaCompareIncrementDbOverviewResponseBody
	GetSuccess() *bool
}

type GetMetaCompareIncrementDbOverviewResponseBody struct {
	// The list of data returned by the API operation. For the element structure, see the descriptions of the child fields.
	Data *GetMetaCompareIncrementDbOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message code.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed. If the call fails, troubleshoot the issue based on errCode and errMessage.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMetaCompareIncrementDbOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCompareIncrementDbOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) GetData() *GetMetaCompareIncrementDbOverviewResponseBodyData {
	return s.Data
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) SetData(v *GetMetaCompareIncrementDbOverviewResponseBodyData) *GetMetaCompareIncrementDbOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) SetErrCode(v string) *GetMetaCompareIncrementDbOverviewResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) SetErrMessage(v string) *GetMetaCompareIncrementDbOverviewResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) SetSuccess(v bool) *GetMetaCompareIncrementDbOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaCompareIncrementDbOverviewResponseBodyData struct {
	// The unique identifier of the current metadata comparison job.
	//
	// example:
	//
	// 100256
	CompareJobId *int64 `json:"compareJobId,omitempty" xml:"compareJobId,omitempty"`
	// The metadata version time corresponding to the current comparison.
	//
	// example:
	//
	// 2025-06-15T10:30:00Z
	CurVersionTime *string `json:"curVersionTime,omitempty" xml:"curVersionTime,omitempty"`
	// The number of partitions added compared to the last comparison.
	//
	// example:
	//
	// 15
	DiffAddPartitionNum *int64 `json:"diffAddPartitionNum,omitempty" xml:"diffAddPartitionNum,omitempty"`
	// The number of tables added compared to the last comparison.
	//
	// example:
	//
	// 5
	DiffAddTableNum *int64 `json:"diffAddTableNum,omitempty" xml:"diffAddTableNum,omitempty"`
	// The number of partitions deleted compared to the last comparison.
	//
	// example:
	//
	// 8
	DiffDelPartitionNum *int64 `json:"diffDelPartitionNum,omitempty" xml:"diffDelPartitionNum,omitempty"`
	// The number of tables deleted compared to the last comparison.
	//
	// example:
	//
	// 2
	DiffDelTableNum *int64 `json:"diffDelTableNum,omitempty" xml:"diffDelTableNum,omitempty"`
	// The number of partitions whose properties changed compared to the last comparison.
	//
	// example:
	//
	// 7
	DiffModifyPartitionNum *int64 `json:"diffModifyPartitionNum,omitempty" xml:"diffModifyPartitionNum,omitempty"`
	// The number of tables whose structure underwent schema evolution compared to the last comparison.
	//
	// example:
	//
	// 5
	DiffModifyTableNum *int64 `json:"diffModifyTableNum,omitempty" xml:"diffModifyTableNum,omitempty"`
	// The total number of differences at the partition level, which is the sum of the number of added, deleted, and modified partitions.
	//
	// example:
	//
	// 30
	DiffPartitionNum *int64 `json:"diffPartitionNum,omitempty" xml:"diffPartitionNum,omitempty"`
	// The total number of differences at the table level, which is the sum of the number of added, deleted, and modified tables.
	//
	// example:
	//
	// 12
	DiffTableNum *int64 `json:"diffTableNum,omitempty" xml:"diffTableNum,omitempty"`
	// The metadata version time corresponding to the last comparison.
	//
	// example:
	//
	// 2025-06-14T10:30:00Z
	PrevVersionTime *string `json:"prevVersionTime,omitempty" xml:"prevVersionTime,omitempty"`
}

func (s GetMetaCompareIncrementDbOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCompareIncrementDbOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetCompareJobId() *int64 {
	return s.CompareJobId
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetCurVersionTime() *string {
	return s.CurVersionTime
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffAddPartitionNum() *int64 {
	return s.DiffAddPartitionNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffAddTableNum() *int64 {
	return s.DiffAddTableNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffDelPartitionNum() *int64 {
	return s.DiffDelPartitionNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffDelTableNum() *int64 {
	return s.DiffDelTableNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffModifyPartitionNum() *int64 {
	return s.DiffModifyPartitionNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffModifyTableNum() *int64 {
	return s.DiffModifyTableNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffPartitionNum() *int64 {
	return s.DiffPartitionNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetDiffTableNum() *int64 {
	return s.DiffTableNum
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) GetPrevVersionTime() *string {
	return s.PrevVersionTime
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetCompareJobId(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.CompareJobId = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetCurVersionTime(v string) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.CurVersionTime = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffAddPartitionNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffAddPartitionNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffAddTableNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffAddTableNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffDelPartitionNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffDelPartitionNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffDelTableNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffDelTableNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffModifyPartitionNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffModifyPartitionNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffModifyTableNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffModifyTableNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffPartitionNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffPartitionNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetDiffTableNum(v int64) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.DiffTableNum = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) SetPrevVersionTime(v string) *GetMetaCompareIncrementDbOverviewResponseBodyData {
	s.PrevVersionTime = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
