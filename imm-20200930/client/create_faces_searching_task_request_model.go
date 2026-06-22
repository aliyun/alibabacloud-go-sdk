// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFacesSearchingTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFacesSearchingTaskRequest
	GetDatasetName() *string
	SetMaxResult(v int64) *CreateFacesSearchingTaskRequest
	GetMaxResult() *int64
	SetNotification(v *Notification) *CreateFacesSearchingTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateFacesSearchingTaskRequest
	GetProjectName() *string
	SetSources(v []*CreateFacesSearchingTaskRequestSources) *CreateFacesSearchingTaskRequest
	GetSources() []*CreateFacesSearchingTaskRequestSources
	SetUserData(v string) *CreateFacesSearchingTaskRequest
	GetUserData() *string
}

type CreateFacesSearchingTaskRequest struct {
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
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of images.
	Sources []*CreateFacesSearchingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// Custom user information. This information is returned in the asynchronous notification message to help you associate the message with your system. The maximum length is 2048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFacesSearchingTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFacesSearchingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFacesSearchingTaskRequest) GetMaxResult() *int64 {
	return s.MaxResult
}

func (s *CreateFacesSearchingTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateFacesSearchingTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFacesSearchingTaskRequest) GetSources() []*CreateFacesSearchingTaskRequestSources {
	return s.Sources
}

func (s *CreateFacesSearchingTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFacesSearchingTaskRequest) SetDatasetName(v string) *CreateFacesSearchingTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetMaxResult(v int64) *CreateFacesSearchingTaskRequest {
	s.MaxResult = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetNotification(v *Notification) *CreateFacesSearchingTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetProjectName(v string) *CreateFacesSearchingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetSources(v []*CreateFacesSearchingTaskRequestSources) *CreateFacesSearchingTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateFacesSearchingTaskRequest) SetUserData(v string) *CreateFacesSearchingTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateFacesSearchingTaskRequest) Validate() error {
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFacesSearchingTaskRequestSources struct {
	// The OSS URI of the image.
	//
	// The OSS URI must follow the format oss\\://${Bucket}/${Object}. `${Bucket}` is the name of the OSS bucket in the same region as the current project. `${Object}` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateFacesSearchingTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateFacesSearchingTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateFacesSearchingTaskRequestSources) SetURI(v string) *CreateFacesSearchingTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateFacesSearchingTaskRequestSources) Validate() error {
	return dara.Validate(s)
}
