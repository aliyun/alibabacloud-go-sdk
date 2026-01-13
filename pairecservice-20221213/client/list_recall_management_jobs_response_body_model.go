// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecallManagementJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecallManagementJobsResponseBody
	GetNextToken() *string
	SetRecallManagementJobs(v []*ListRecallManagementJobsResponseBodyRecallManagementJobs) *ListRecallManagementJobsResponseBody
	GetRecallManagementJobs() []*ListRecallManagementJobsResponseBodyRecallManagementJobs
	SetRequestId(v string) *ListRecallManagementJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRecallManagementJobsResponseBody
	GetTotalCount() *string
}

type ListRecallManagementJobsResponseBody struct {
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ""
	NextToken            *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RecallManagementJobs []*ListRecallManagementJobsResponseBodyRecallManagementJobs `json:"RecallManagementJobs,omitempty" xml:"RecallManagementJobs,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecallManagementJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecallManagementJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementJobsResponseBody) GetRecallManagementJobs() []*ListRecallManagementJobsResponseBodyRecallManagementJobs {
	return s.RecallManagementJobs
}

func (s *ListRecallManagementJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecallManagementJobsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRecallManagementJobsResponseBody) SetMaxResults(v int32) *ListRecallManagementJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementJobsResponseBody) SetNextToken(v string) *ListRecallManagementJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementJobsResponseBody) SetRecallManagementJobs(v []*ListRecallManagementJobsResponseBodyRecallManagementJobs) *ListRecallManagementJobsResponseBody {
	s.RecallManagementJobs = v
	return s
}

func (s *ListRecallManagementJobsResponseBody) SetRequestId(v string) *ListRecallManagementJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecallManagementJobsResponseBody) SetTotalCount(v string) *ListRecallManagementJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecallManagementJobsResponseBody) Validate() error {
	if s.RecallManagementJobs != nil {
		for _, item := range s.RecallManagementJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecallManagementJobsResponseBodyRecallManagementJobs struct {
	// example:
	//
	// 2025-03-28T10:24Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	RecallManagementJobId  *string                                                                         `json:"RecallManagementJobId,omitempty" xml:"RecallManagementJobId,omitempty"`
	RecallManagerTableInfo *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo `json:"RecallManagerTableInfo,omitempty" xml:"RecallManagerTableInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2025-01-28T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecallManagementJobsResponseBodyRecallManagementJobs) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementJobsResponseBodyRecallManagementJobs) GoString() string {
	return s.String()
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) GetRecallManagementJobId() *string {
	return s.RecallManagementJobId
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) GetRecallManagerTableInfo() *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo {
	return s.RecallManagerTableInfo
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) GetStatus() *string {
	return s.Status
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) SetEndTime(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobs {
	s.EndTime = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) SetRecallManagementJobId(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobs {
	s.RecallManagementJobId = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) SetRecallManagerTableInfo(v *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) *ListRecallManagementJobsResponseBodyRecallManagementJobs {
	s.RecallManagerTableInfo = v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) SetStartTime(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobs {
	s.StartTime = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) SetStatus(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobs {
	s.Status = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobs) Validate() error {
	if s.RecallManagerTableInfo != nil {
		if err := s.RecallManagerTableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo struct {
	// example:
	//
	// ds=20250701
	DataVersion *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	// example:
	//
	// 2
	RecallManagementTableVersionId *string `json:"RecallManagementTableVersionId,omitempty" xml:"RecallManagementTableVersionId,omitempty"`
	// example:
	//
	// 1000
	SourceTableDataSize *string `json:"SourceTableDataSize,omitempty" xml:"SourceTableDataSize,omitempty"`
	// example:
	//
	// 100
	SourceTableRowCount *string `json:"SourceTableRowCount,omitempty" xml:"SourceTableRowCount,omitempty"`
}

func (s ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) GoString() string {
	return s.String()
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) GetDataVersion() *string {
	return s.DataVersion
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) GetRecallManagementTableVersionId() *string {
	return s.RecallManagementTableVersionId
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) GetSourceTableDataSize() *string {
	return s.SourceTableDataSize
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) GetSourceTableRowCount() *string {
	return s.SourceTableRowCount
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) SetDataVersion(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo {
	s.DataVersion = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) SetRecallManagementTableVersionId(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo {
	s.RecallManagementTableVersionId = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) SetSourceTableDataSize(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo {
	s.SourceTableDataSize = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) SetSourceTableRowCount(v string) *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo {
	s.SourceTableRowCount = &v
	return s
}

func (s *ListRecallManagementJobsResponseBodyRecallManagementJobsRecallManagerTableInfo) Validate() error {
	return dara.Validate(s)
}
