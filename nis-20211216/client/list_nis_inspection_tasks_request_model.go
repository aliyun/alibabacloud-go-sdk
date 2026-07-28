// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionName(v string) *ListNisInspectionTasksRequest
	GetInspectionName() *string
	SetInspectionProject(v string) *ListNisInspectionTasksRequest
	GetInspectionProject() *string
	SetInspectionTaskId(v string) *ListNisInspectionTasksRequest
	GetInspectionTaskId() *string
	SetMaxResults(v int32) *ListNisInspectionTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNisInspectionTasksRequest
	GetNextToken() *string
	SetStatus(v string) *ListNisInspectionTasksRequest
	GetStatus() *string
}

type ListNisInspectionTasksRequest struct {
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
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the NextToken value returned from a previous call.
	//
	// example:
	//
	// hKrS+MVXkuOgztXnvdml1/R9jhHkiH8eW3CfaOYU0CEL7yiT0zae6J8v1zYNg+d1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The running status of the task. Valid values:
	//
	// - Creating
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

func (s ListNisInspectionTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTasksRequest) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTasksRequest) GetInspectionName() *string {
	return s.InspectionName
}

func (s *ListNisInspectionTasksRequest) GetInspectionProject() *string {
	return s.InspectionProject
}

func (s *ListNisInspectionTasksRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *ListNisInspectionTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNisInspectionTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNisInspectionTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListNisInspectionTasksRequest) SetInspectionName(v string) *ListNisInspectionTasksRequest {
	s.InspectionName = &v
	return s
}

func (s *ListNisInspectionTasksRequest) SetInspectionProject(v string) *ListNisInspectionTasksRequest {
	s.InspectionProject = &v
	return s
}

func (s *ListNisInspectionTasksRequest) SetInspectionTaskId(v string) *ListNisInspectionTasksRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *ListNisInspectionTasksRequest) SetMaxResults(v int32) *ListNisInspectionTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNisInspectionTasksRequest) SetNextToken(v string) *ListNisInspectionTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListNisInspectionTasksRequest) SetStatus(v string) *ListNisInspectionTasksRequest {
	s.Status = &v
	return s
}

func (s *ListNisInspectionTasksRequest) Validate() error {
	return dara.Validate(s)
}
