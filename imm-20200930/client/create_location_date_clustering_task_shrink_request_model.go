// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationDateClusteringTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetDatasetName() *string
	SetDateOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetDateOptionsShrink() *string
	SetLocationOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetLocationOptionsShrink() *string
	SetNotificationShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetUserData() *string
}

type CreateLocationDateClusteringTaskShrinkRequest struct {
	// The dataset name. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The date clustering settings.
	//
	// 	Notice: Modifying this setting also affects existing spatio-temporal clusters in your `Dataset`.
	//
	// This parameter is required.
	DateOptionsShrink *string `json:"DateOptions,omitempty" xml:"DateOptions,omitempty"`
	// The location clustering settings.
	//
	// 	Notice: Modifying this setting also affects existing spatio-temporal clusters in your `Dataset`.
	//
	// This parameter is required.
	LocationOptionsShrink *string `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty"`
	// The message notification configuration. For more information, see Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom information that is returned in the asynchronous notification message. This helps you associate the notification message with your system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateLocationDateClusteringTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetDateOptionsShrink() *string {
	return s.DateOptionsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetLocationOptionsShrink() *string {
	return s.LocationOptionsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDateOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DateOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetLocationOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.LocationOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetProjectName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetUserData(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
