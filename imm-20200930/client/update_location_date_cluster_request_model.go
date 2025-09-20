// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationDateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomId(v string) *UpdateLocationDateClusterRequest
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *UpdateLocationDateClusterRequest
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *UpdateLocationDateClusterRequest
	GetDatasetName() *string
	SetObjectId(v string) *UpdateLocationDateClusterRequest
	GetObjectId() *string
	SetProjectName(v string) *UpdateLocationDateClusterRequest
	GetProjectName() *string
	SetTitle(v string) *UpdateLocationDateClusterRequest
	GetTitle() *string
}

type UpdateLocationDateClusterRequest struct {
	// The custom ID of the cluster. When the cluster is indexed into the dataset, the custom ID is stored as the data attribute. You can map the custom ID to other data in your business system. For example, you can pass the custom ID to map a URI to an ID. We recommend that you specify a globally unique value. The value can be up to 1,024 bytes in size.
	//
	// example:
	//
	// member-id-0001
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels. The parameter stores custom key-value labels, which can be used to filter data. You can specify up to 100 custom labels for a cluster.
	//
	// example:
	//
	// {
	//
	//       "UserScore": "5"
	//
	// }
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the cluster that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// location-date-cluster-71dd4f32-9597-4085-a2ab-3a7b0fd0aff9
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the cluster. The name can be used to search for the cluster. The value can be up to 1,024 bytes in size.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLocationDateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationDateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *UpdateLocationDateClusterRequest) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *UpdateLocationDateClusterRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateLocationDateClusterRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateLocationDateClusterRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateLocationDateClusterRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateLocationDateClusterRequest) SetCustomId(v string) *UpdateLocationDateClusterRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetCustomLabels(v map[string]interface{}) *UpdateLocationDateClusterRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetDatasetName(v string) *UpdateLocationDateClusterRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetObjectId(v string) *UpdateLocationDateClusterRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetProjectName(v string) *UpdateLocationDateClusterRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) SetTitle(v string) *UpdateLocationDateClusterRequest {
	s.Title = &v
	return s
}

func (s *UpdateLocationDateClusterRequest) Validate() error {
	return dara.Validate(s)
}
