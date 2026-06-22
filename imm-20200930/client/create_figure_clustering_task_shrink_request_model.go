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
	// The dataset name. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The notification configuration. For more information, see Notification. For more information about the message format, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom tags that you can use to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"test": "val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom user data that is returned in the asynchronous notification message. You can use this data to associate the notification message with your services. The maximum length is 2048 bytes.
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
