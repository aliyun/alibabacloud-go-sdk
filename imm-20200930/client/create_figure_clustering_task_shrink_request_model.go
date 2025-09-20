// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClusteringTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFigureClusteringTaskShrinkRequest
	GetDatasetName() *string
	SetNotificationShrink(v string) *CreateFigureClusteringTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateFigureClusteringTaskShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *CreateFigureClusteringTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateFigureClusteringTaskShrinkRequest
	GetUserData() *string
}

type CreateFigureClusteringTaskShrinkRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test": "val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClusteringTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFigureClusteringTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateFigureClusteringTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFigureClusteringTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateFigureClusteringTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetProjectName(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) SetUserData(v string) *CreateFigureClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateFigureClusteringTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
