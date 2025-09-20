// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClustersMergingTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateFigureClustersMergingTaskRequest
	GetDatasetName() *string
	SetFrom(v string) *CreateFigureClustersMergingTaskRequest
	GetFrom() *string
	SetFroms(v []*string) *CreateFigureClustersMergingTaskRequest
	GetFroms() []*string
	SetNotification(v *Notification) *CreateFigureClustersMergingTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateFigureClustersMergingTaskRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *CreateFigureClustersMergingTaskRequest
	GetTags() map[string]interface{}
	SetTo(v string) *CreateFigureClustersMergingTaskRequest
	GetTo() *string
	SetUserData(v string) *CreateFigureClustersMergingTaskRequest
	GetUserData() *string
}

type CreateFigureClustersMergingTaskRequest struct {
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the source group. You must specify either From or Froms, but not both.
	//
	// example:
	//
	// Cluster-2ab85905-23ba-4632-b2d8-1c21cfe****
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The IDs of source clustering groups. You must specify either From or Froms, but not both. You can specify up to 100 task IDs.
	Froms []*string `json:"Froms,omitempty" xml:"Froms,omitempty" type:"Repeated"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The custom tags, which can be used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"key":"val"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the destination clustering group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster-4a3a71c1-c092-4788-8826-2f65d17****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The custom data, which is returned in an asynchronous notification and facilitates notification management. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateFigureClustersMergingTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClustersMergingTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateFigureClustersMergingTaskRequest) GetFrom() *string {
	return s.From
}

func (s *CreateFigureClustersMergingTaskRequest) GetFroms() []*string {
	return s.Froms
}

func (s *CreateFigureClustersMergingTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateFigureClustersMergingTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFigureClustersMergingTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateFigureClustersMergingTaskRequest) GetTo() *string {
	return s.To
}

func (s *CreateFigureClustersMergingTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateFigureClustersMergingTaskRequest) SetDatasetName(v string) *CreateFigureClustersMergingTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetFrom(v string) *CreateFigureClustersMergingTaskRequest {
	s.From = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetFroms(v []*string) *CreateFigureClustersMergingTaskRequest {
	s.Froms = v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetNotification(v *Notification) *CreateFigureClustersMergingTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetProjectName(v string) *CreateFigureClustersMergingTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetTags(v map[string]interface{}) *CreateFigureClustersMergingTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetTo(v string) *CreateFigureClustersMergingTaskRequest {
	s.To = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) SetUserData(v string) *CreateFigureClustersMergingTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateFigureClustersMergingTaskRequest) Validate() error {
	return dara.Validate(s)
}
