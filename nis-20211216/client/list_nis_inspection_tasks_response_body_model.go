// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskList(v []*ListNisInspectionTasksResponseBodyInspectionTaskList) *ListNisInspectionTasksResponseBody
	GetInspectionTaskList() []*ListNisInspectionTasksResponseBodyInspectionTaskList
	SetMaxResults(v int32) *ListNisInspectionTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNisInspectionTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNisInspectionTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNisInspectionTasksResponseBody
	GetTotalCount() *int32
}

type ListNisInspectionTasksResponseBody struct {
	// The list of network inspection tasks.
	InspectionTaskList []*ListNisInspectionTasksResponseBodyInspectionTaskList `json:"InspectionTaskList,omitempty" xml:"InspectionTaskList,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. If the value of this parameter is not empty, it indicates that there are more results to retrieve. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// hKrS+MVXkuOgztXnvdml194Cz/lMNdmr+DEh0th6dVlNEo/F148UPCh2itDku7Qj
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNisInspectionTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTasksResponseBody) GetInspectionTaskList() []*ListNisInspectionTasksResponseBodyInspectionTaskList {
	return s.InspectionTaskList
}

func (s *ListNisInspectionTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNisInspectionTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNisInspectionTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNisInspectionTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNisInspectionTasksResponseBody) SetInspectionTaskList(v []*ListNisInspectionTasksResponseBodyInspectionTaskList) *ListNisInspectionTasksResponseBody {
	s.InspectionTaskList = v
	return s
}

func (s *ListNisInspectionTasksResponseBody) SetMaxResults(v int32) *ListNisInspectionTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNisInspectionTasksResponseBody) SetNextToken(v string) *ListNisInspectionTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNisInspectionTasksResponseBody) SetRequestId(v string) *ListNisInspectionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNisInspectionTasksResponseBody) SetTotalCount(v int32) *ListNisInspectionTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNisInspectionTasksResponseBody) Validate() error {
	if s.InspectionTaskList != nil {
		for _, item := range s.InspectionTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNisInspectionTasksResponseBodyInspectionTaskList struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2024-06-18 00:14:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the network inspection task.
	//
	// example:
	//
	// NIS inspection
	InspectionName *string `json:"InspectionName,omitempty" xml:"InspectionName,omitempty"`
	// The type of inspection solution that the network inspection task uses. Valid values: basic and customized.
	//
	// example:
	//
	// basic
	InspectionProject *string `json:"InspectionProject,omitempty" xml:"InspectionProject,omitempty"`
	// The ID of the network inspection task.
	//
	// example:
	//
	// ni-8svm******hzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// The ID of the latest report.
	//
	// example:
	//
	// nir-b4c4c9******8a25e
	LastUpdateReportId *string `json:"LastUpdateReportId,omitempty" xml:"LastUpdateReportId,omitempty"`
	// The running status of the task. Valid values:
	//
	// Creating: The task is being created.
	//
	// - Active
	//
	// - Running
	//
	// - Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNisInspectionTasksResponseBodyInspectionTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTasksResponseBodyInspectionTaskList) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetInspectionName() *string {
	return s.InspectionName
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetInspectionProject() *string {
	return s.InspectionProject
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetLastUpdateReportId() *string {
	return s.LastUpdateReportId
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetCreateTime(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetInspectionName(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.InspectionName = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetInspectionProject(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.InspectionProject = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetInspectionTaskId(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.InspectionTaskId = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetLastUpdateReportId(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.LastUpdateReportId = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) SetStatus(v string) *ListNisInspectionTasksResponseBodyInspectionTaskList {
	s.Status = &v
	return s
}

func (s *ListNisInspectionTasksResponseBodyInspectionTaskList) Validate() error {
	return dara.Validate(s)
}
