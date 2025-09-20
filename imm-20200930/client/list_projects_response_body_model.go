// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListProjectsResponseBody
	GetNextToken() *string
	SetProjects(v []*Project) *ListProjectsResponseBody
	GetProjects() []*Project
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDAx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The projects.
	Projects []*Project `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4A7A2D0E-D8B8-4DA0-8127-EB32C660
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBody) GetProjects() []*Project {
	return s.Projects
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*Project) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
