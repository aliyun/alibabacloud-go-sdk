// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchIndexFileMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchIndexFileMetaShrinkRequest
	GetDatasetName() *string
	SetFilesShrink(v string) *BatchIndexFileMetaShrinkRequest
	GetFilesShrink() *string
	SetNotificationShrink(v string) *BatchIndexFileMetaShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *BatchIndexFileMetaShrinkRequest
	GetProjectName() *string
	SetUserData(v string) *BatchIndexFileMetaShrinkRequest
	GetUserData() *string
}

type BatchIndexFileMetaShrinkRequest struct {
	// The dataset name. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// A list of OSS files. This is an array in JSON format that can contain up to 100 files.
	//
	// This parameter is required.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// The notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see the metadata indexing section in [Asynchronous notification message formats](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom user data. This parameter takes effect only when you specify an MNS configuration for the Notification parameter. The data is returned in the asynchronous notification message, which you can use to associate the message with your services. The maximum length is 2048 bytes.
	//
	// example:
	//
	// {
	//
	//       "id": "test-id",
	//
	//       "name": "test-name"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchIndexFileMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchIndexFileMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchIndexFileMetaShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *BatchIndexFileMetaShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *BatchIndexFileMetaShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchIndexFileMetaShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *BatchIndexFileMetaShrinkRequest) SetDatasetName(v string) *BatchIndexFileMetaShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetFilesShrink(v string) *BatchIndexFileMetaShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetNotificationShrink(v string) *BatchIndexFileMetaShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetProjectName(v string) *BatchIndexFileMetaShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) SetUserData(v string) *BatchIndexFileMetaShrinkRequest {
	s.UserData = &v
	return s
}

func (s *BatchIndexFileMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
