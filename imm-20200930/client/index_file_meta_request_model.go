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
	File *InputFile `json:"File,omitempty" xml:"File,omitempty"`
	// The message notification configuration. For more information, see Notification. For the format of the asynchronous notification message, see the Metadata Indexing section in [Asynchronous notification message formats](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
