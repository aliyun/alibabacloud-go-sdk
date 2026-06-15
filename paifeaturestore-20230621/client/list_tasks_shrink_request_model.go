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
	// The ID of the object. You can call the ListModelFeatures or ListFeatureViews operation to obtain the ID.
	//
	// example:
	//
	// 4
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the object.
	//
	// ● `ModelFeature`: a model feature.
	//
	// ● `FeatureView`: a feature view.
	//
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the project. You can call the ListProjects operation to obtain the ID.
	//
	// example:
	//
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The task status.
	//
	// ● `Initializing`: The task is being initialized.
	//
	// ● `Running`: The task is running.
	//
	// ● `Success`: The task is successful.
	//
	// ● `Failure`: The task fails.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task IDs by which to filter tasks.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The task type by which to filter tasks.
	//
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
