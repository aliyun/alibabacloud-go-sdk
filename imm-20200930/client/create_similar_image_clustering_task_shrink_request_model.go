// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusteringTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateSimilarImageClusteringTaskShrinkRequest
	GetDatasetName() *string
	SetNotificationShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateSimilarImageClusteringTaskShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateSimilarImageClusteringTaskShrinkRequest
	GetUserData() *string
}

type CreateSimilarImageClusteringTaskShrinkRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The configuration of the notification message. For more information, click Notification. For more information about the format of asynchronous notification messages, see [Asynchronous notification messages](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags that are used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information. This information is returned in the asynchronous notification message to help you associate the message with your system. The value can be up to 2,048 bytes in length.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateSimilarImageClusteringTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetProjectName(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) SetUserData(v string) *CreateSimilarImageClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
