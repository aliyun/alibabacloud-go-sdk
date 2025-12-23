// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListProjectsResponseBodyData) *ListProjectsResponseBody
	GetData() []*ListProjectsResponseBodyData
	SetNextToken(v string) *ListProjectsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListProjectsResponseBody
	GetTotalSize() *int32
}

type ListProjectsResponseBody struct {
	Data []*ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetData() []*ListProjectsResponseBodyData {
	return s.Data
}

func (s *ListProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListProjectsResponseBody) SetData(v []*ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalSize(v int32) *ListProjectsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyData struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s ListProjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListProjectsResponseBodyData) SetDescription(v string) *ListProjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetName(v string) *ListProjectsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectId(v string) *ListProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
