// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchIndexFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchIndexFileMetaRequest
	GetDatasetName() *string
	SetFiles(v []*InputFile) *BatchIndexFileMetaRequest
	GetFiles() []*InputFile
	SetNotification(v *Notification) *BatchIndexFileMetaRequest
	GetNotification() *Notification
	SetProjectName(v string) *BatchIndexFileMetaRequest
	GetProjectName() *string
	SetUserData(v string) *BatchIndexFileMetaRequest
	GetUserData() *string
}

type BatchIndexFileMetaRequest struct {
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
	Files []*InputFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see the metadata indexing section in [Asynchronous notification message formats](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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

func (s BatchIndexFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchIndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchIndexFileMetaRequest) GetFiles() []*InputFile {
	return s.Files
}

func (s *BatchIndexFileMetaRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *BatchIndexFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchIndexFileMetaRequest) GetUserData() *string {
	return s.UserData
}

func (s *BatchIndexFileMetaRequest) SetDatasetName(v string) *BatchIndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetFiles(v []*InputFile) *BatchIndexFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetNotification(v *Notification) *BatchIndexFileMetaRequest {
	s.Notification = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetProjectName(v string) *BatchIndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetUserData(v string) *BatchIndexFileMetaRequest {
	s.UserData = &v
	return s
}

func (s *BatchIndexFileMetaRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
