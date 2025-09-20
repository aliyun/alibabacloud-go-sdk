// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusteringTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateSimilarImageClusteringTaskRequest
	GetDatasetName() *string
	SetNotification(v *Notification) *CreateSimilarImageClusteringTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateSimilarImageClusteringTaskRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *CreateSimilarImageClusteringTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateSimilarImageClusteringTaskRequest
	GetUserData() *string
}

type CreateSimilarImageClusteringTaskRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
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
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateSimilarImageClusteringTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateSimilarImageClusteringTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateSimilarImageClusteringTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSimilarImageClusteringTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateSimilarImageClusteringTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateSimilarImageClusteringTaskRequest) SetDatasetName(v string) *CreateSimilarImageClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetNotification(v *Notification) *CreateSimilarImageClusteringTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetProjectName(v string) *CreateSimilarImageClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateSimilarImageClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) SetUserData(v string) *CreateSimilarImageClusteringTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskRequest) Validate() error {
	return dara.Validate(s)
}
