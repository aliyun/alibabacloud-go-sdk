// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListTasksResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListTasksResponseBody
	GetNextToken() *string
	SetProjectName(v string) *ListTasksResponseBody
	GetProjectName() *string
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*TaskInfo) *ListTasksResponseBody
	GetTasks() []*TaskInfo
}

type ListTasksResponseBody struct {
	// The length of the returned result list.
	//
	// example:
	//
	// 1
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. The pagination token is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter. This parameter has a value only when not all results are returned.
	//
	// You can specify the value of the NextToken parameter in the next request to list remaining results.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9847E7D0-A9A3-0053-84C6-BA16FFFA726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*TaskInfo `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTasks() []*TaskInfo {
	return s.Tasks
}

func (s *ListTasksResponseBody) SetMaxResults(v string) *ListTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetProjectName(v string) *ListTasksResponseBody {
	s.ProjectName = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*TaskInfo) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
