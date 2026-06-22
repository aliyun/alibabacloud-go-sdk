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
	// The name of the dataset. To get the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The file to be indexed, in JSON format. For more information, see the struct definition.
	//
	// This parameter is required.
	FileShrink *string `json:"File,omitempty" xml:"File,omitempty"`
	// The message notification configuration. For more information, see Notification. For the format of the asynchronous notification message, see the Metadata Indexing section in [Asynchronous notification message formats](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. To get the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom information that is returned in the asynchronous notification message. This helps you associate the notification with your services. The maximum length is 2048 bytes.
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
