// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectId(v string) *ListTasksShrinkRequest
	GetObjectId() *string
	SetObjectType(v string) *ListTasksShrinkRequest
	GetObjectType() *string
	SetPageNumber(v int32) *ListTasksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListTasksShrinkRequest
	GetProjectId() *string
	SetStatus(v string) *ListTasksShrinkRequest
	GetStatus() *string
	SetTaskIdsShrink(v string) *ListTasksShrinkRequest
	GetTaskIdsShrink() *string
	SetType(v string) *ListTasksShrinkRequest
	GetType() *string
}

type ListTasksShrinkRequest struct {
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
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListTasksShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListTasksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *ListTasksShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListTasksShrinkRequest) SetObjectId(v string) *ListTasksShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetObjectType(v string) *ListTasksShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageNumber(v int32) *ListTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageSize(v int32) *ListTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectId(v string) *ListTasksShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskIdsShrink(v string) *ListTasksShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetType(v string) *ListTasksShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
