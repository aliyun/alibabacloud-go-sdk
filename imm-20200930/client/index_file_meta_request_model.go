// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *IndexFileMetaRequest
	GetDatasetName() *string
	SetFile(v *InputFile) *IndexFileMetaRequest
	GetFile() *InputFile
	SetNotification(v *Notification) *IndexFileMetaRequest
	GetNotification() *Notification
	SetProjectName(v string) *IndexFileMetaRequest
	GetProjectName() *string
	SetUserData(v string) *IndexFileMetaRequest
	GetUserData() *string
}

type IndexFileMetaRequest struct {
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
	File *InputFile `json:"File,omitempty" xml:"File,omitempty"`
	// The notification settings. For more information, see the "Metadata indexing" section of the [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html) topic.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s IndexFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *IndexFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *IndexFileMetaRequest) GetFile() *InputFile {
	return s.File
}

func (s *IndexFileMetaRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *IndexFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *IndexFileMetaRequest) GetUserData() *string {
	return s.UserData
}

func (s *IndexFileMetaRequest) SetDatasetName(v string) *IndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *IndexFileMetaRequest) SetFile(v *InputFile) *IndexFileMetaRequest {
	s.File = v
	return s
}

func (s *IndexFileMetaRequest) SetNotification(v *Notification) *IndexFileMetaRequest {
	s.Notification = v
	return s
}

func (s *IndexFileMetaRequest) SetProjectName(v string) *IndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *IndexFileMetaRequest) SetUserData(v string) *IndexFileMetaRequest {
	s.UserData = &v
	return s
}

func (s *IndexFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
