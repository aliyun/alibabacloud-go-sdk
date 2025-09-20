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
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of the most similar faces that you want to return. Valid values: 1 to 100. Default value: 5.
	//
	// example:
	//
	// 100
	MaxResult *int64 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The images.
	Sources []*CreateFacesSearchingTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
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
	return dara.Validate(s)
}

type CreateFacesSearchingTaskRequestSources struct {
	// The OSS URI of the image.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
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
