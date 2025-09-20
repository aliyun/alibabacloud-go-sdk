// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClustersMergingTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetDatasetName() *string
	SetFrom(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetFrom() *string
	SetFromsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetFromsShrink() *string
	SetNotificationShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetTagsShrink() *string
	SetTo(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetTo() *string
	SetUserData(v string) *CreateFigureClustersMergingTaskShrinkRequest
	GetUserData() *string
}

type CreateFigureClustersMergingTaskShrinkRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the source group. You must specify either From or Froms, but not both.
	//
	// example:
	//
	// Cluster-2ab85905-23ba-4632-b2d8-1c21cfe****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The IDs of source clustering groups. You must specify either From or Froms, but not both. You can specify up to 100 task IDs.
	FromsShrink *string `json:"Froms,omitempty" xml:"Froms,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags, which can be used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"key":"val"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the destination clustering group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster-4a3a71c1-c092-4788-8826-2f65d17****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The custom data, which is returned in an asynchronous notification and facilitates notification management. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClustersMergingTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClustersMergingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetFrom() *string {
	return s.From
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetFromsShrink() *string {
	return s.FromsShrink
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetTo() *string {
	return s.To
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetDatasetName(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetFrom(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.From = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetFromsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.FromsShrink = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetNotificationShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetProjectName(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetTagsShrink(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetTo(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.To = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) SetUserData(v string) *CreateFigureClustersMergingTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateFigureClustersMergingTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
