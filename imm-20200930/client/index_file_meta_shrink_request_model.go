// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *IndexFileMetaShrinkRequest
	GetDatasetName() *string
	SetFileShrink(v string) *IndexFileMetaShrinkRequest
	GetFileShrink() *string
	SetNotificationShrink(v string) *IndexFileMetaShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *IndexFileMetaShrinkRequest
	GetProjectName() *string
	SetUserData(v string) *IndexFileMetaShrinkRequest
	GetUserData() *string
}

type IndexFileMetaShrinkRequest struct {
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The file for which you want to create a metadata index. The value must be in the JSON format.
	//
	// This parameter is required.
	FileShrink *string `json:"File,omitempty" xml:"File,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom user information, which is returned in an asynchronous notification. The maximum length of a notification is 2048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s IndexFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *IndexFileMetaShrinkRequest) GetFileShrink() *string {
	return s.FileShrink
}

func (s *IndexFileMetaShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *IndexFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *IndexFileMetaShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *IndexFileMetaShrinkRequest) SetDatasetName(v string) *IndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetFileShrink(v string) *IndexFileMetaShrinkRequest {
	s.FileShrink = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetNotificationShrink(v string) *IndexFileMetaShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetProjectName(v string) *IndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) SetUserData(v string) *IndexFileMetaShrinkRequest {
	s.UserData = &v
	return s
}

func (s *IndexFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
