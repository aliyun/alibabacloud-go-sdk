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
	// The dataset name. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The notification configuration. For more information, see Notification. For more information about the message format, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom tags that you can use to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"test": "val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom user data that is returned in the asynchronous notification message. You can use this data to associate the notification message with your services. The maximum length is 2048 bytes.
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
