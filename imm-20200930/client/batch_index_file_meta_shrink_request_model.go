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
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The objects in Object Storage Service (OSS). Specify OSS objects by using a JSON array. You can specify up to 100 objects in an array.
	//
	// This parameter is required.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The user-defined data that you want to return in asynchronous messages. This parameter takes effect only when you specify the MNS settings in the Notification parameter. The maximum information length is 2,048 bytes.
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
