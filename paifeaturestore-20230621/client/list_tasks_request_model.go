// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectId(v string) *ListTasksRequest
	GetObjectId() *string
	SetObjectType(v string) *ListTasksRequest
	GetObjectType() *string
	SetPageNumber(v int32) *ListTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListTasksRequest
	GetProjectId() *string
	SetStatus(v string) *ListTasksRequest
	GetStatus() *string
	SetTaskIds(v []*string) *ListTasksRequest
	GetTaskIds() []*string
	SetType(v string) *ListTasksRequest
	GetType() *string
}

type ListTasksRequest struct {
	// example:
	//
	// 4
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Running
	Status  *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListTasksRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListTasksRequest) SetObjectId(v string) *ListTasksRequest {
	s.ObjectId = &v
	return s
}

func (s *ListTasksRequest) SetObjectType(v string) *ListTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v string) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTaskIds(v []*string) *ListTasksRequest {
	s.TaskIds = v
	return s
}

func (s *ListTasksRequest) SetType(v string) *ListTasksRequest {
	s.Type = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
