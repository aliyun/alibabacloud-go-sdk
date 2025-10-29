// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScanRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListDataQualityScanRunsResponseBodyPageInfo) *ListDataQualityScanRunsResponseBody
	GetPageInfo() *ListDataQualityScanRunsResponseBodyPageInfo
	SetRequestId(v string) *ListDataQualityScanRunsResponseBody
	GetRequestId() *string
}

type ListDataQualityScanRunsResponseBody struct {
	// The page information.
	PageInfo *ListDataQualityScanRunsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityScanRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsResponseBody) GetPageInfo() *ListDataQualityScanRunsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListDataQualityScanRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityScanRunsResponseBody) SetPageInfo(v *ListDataQualityScanRunsResponseBodyPageInfo) *ListDataQualityScanRunsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListDataQualityScanRunsResponseBody) SetRequestId(v string) *ListDataQualityScanRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityScanRunsResponseBodyPageInfo struct {
	// The list of data quality monitor run records.
	DataQualityScanRuns []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns `json:"DataQualityScanRuns,omitempty" xml:"DataQualityScanRuns,omitempty" type:"Repeated"`
	// The page number of the results. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records returned.
	//
	// example:
	//
	// 324
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityScanRunsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) GetDataQualityScanRuns() []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	return s.DataQualityScanRuns
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) SetDataQualityScanRuns(v []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) *ListDataQualityScanRunsResponseBodyPageInfo {
	s.DataQualityScanRuns = v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) SetPageNumber(v int32) *ListDataQualityScanRunsResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) SetPageSize(v int32) *ListDataQualityScanRunsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) SetTotalCount(v int32) *ListDataQualityScanRunsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfo) Validate() error {
	if s.DataQualityScanRuns != nil {
		for _, item := range s.DataQualityScanRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns struct {
	// The time when the data quality monitor starts running.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the data quality monitor stops.
	//
	// example:
	//
	// 1710239005403
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the data quality monitor running record.
	//
	// example:
	//
	// 3155
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters configured for the instance.
	Parameters []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The status of the instance.
	//
	// 	- Pass
	//
	// 	- Running
	//
	// 	- Error
	//
	// 	- Warn
	//
	// 	- Fail
	//
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GetParameters() []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters {
	return s.Parameters
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) SetCreateTime(v int64) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) SetFinishTime(v int64) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	s.FinishTime = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) SetId(v int64) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	s.Id = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) SetParameters(v []*ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	s.Parameters = v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) SetStatus(v string) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns {
	s.Status = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRuns) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters struct {
	// The parameter name.
	//
	// example:
	//
	// dt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value. You can use a scheduling time expression.
	//
	// example:
	//
	// $[yyyy-mm-dd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) GetName() *string {
	return s.Name
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) SetName(v string) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters {
	s.Name = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) SetValue(v string) *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters {
	s.Value = &v
	return s
}

func (s *ListDataQualityScanRunsResponseBodyPageInfoDataQualityScanRunsParameters) Validate() error {
	return dara.Validate(s)
}
