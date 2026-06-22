// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFacesSearchingTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFacesSearchingTaskShrinkRequest
	GetDatasetName() *string
	SetMaxResult(v int64) *CreateFacesSearchingTaskShrinkRequest
	GetMaxResult() *int64
	SetNotificationShrink(v string) *CreateFacesSearchingTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateFacesSearchingTaskShrinkRequest
	GetProjectName() *string
	SetSourcesShrink(v string) *CreateFacesSearchingTaskShrinkRequest
	GetSourcesShrink() *string
	SetUserData(v string) *CreateFacesSearchingTaskShrinkRequest
	GetUserData() *string
}

type CreateFacesSearchingTaskShrinkRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of most similar faces to return. Valid values: 0 to 100. Default value: 5.
	//
	// example:
	//
	// 100
	MaxResult *int64 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The notification configuration. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of images.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// Custom user information. This information is returned in the asynchronous notification message to help you associate the message with your system. The maximum length is 2048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFacesSearchingTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFacesSearchingTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetMaxResult() *int64 {
	return s.MaxResult
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateFacesSearchingTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetDatasetName(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetMaxResult(v int64) *CreateFacesSearchingTaskShrinkRequest {
	s.MaxResult = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetNotificationShrink(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetProjectName(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetSourcesShrink(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) SetUserData(v string) *CreateFacesSearchingTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateFacesSearchingTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
