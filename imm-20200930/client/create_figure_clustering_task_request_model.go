// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClusteringTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFigureClusteringTaskRequest
	GetDatasetName() *string
	SetNotification(v *Notification) *CreateFigureClusteringTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateFigureClusteringTaskRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *CreateFigureClusteringTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateFigureClusteringTaskRequest
	GetUserData() *string
}

type CreateFigureClusteringTaskRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test": "val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClusteringTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFigureClusteringTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateFigureClusteringTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFigureClusteringTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateFigureClusteringTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFigureClusteringTaskRequest) SetDatasetName(v string) *CreateFigureClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetNotification(v *Notification) *CreateFigureClusteringTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetProjectName(v string) *CreateFigureClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateFigureClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateFigureClusteringTaskRequest) SetUserData(v string) *CreateFigureClusteringTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateFigureClusteringTaskRequest) Validate() error {
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
