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
	InspectionTaskList []*ListNisInspectionTasksResponseBodyInspectionTaskList `json:"InspectionTaskList,omitempty" xml:"InspectionTaskList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// hKrS+MVXkuOgztXnvdml194Cz/lMNdmr+DEh0th6dVlNEo/F148UPCh2itDku7Qj
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2024-06-18 00:14:46
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InspectionName *string `json:"InspectionName,omitempty" xml:"InspectionName,omitempty"`
	// example:
	//
	// basic
	InspectionProject *string `json:"InspectionProject,omitempty" xml:"InspectionProject,omitempty"`
	// example:
	//
	// ni-8svm******hzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// example:
	//
	// nir-b4c4c9******8a25e
	LastUpdateReportId *string `json:"LastUpdateReportId,omitempty" xml:"LastUpdateReportId,omitempty"`
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
