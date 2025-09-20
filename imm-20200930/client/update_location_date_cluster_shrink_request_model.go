// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocationDateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomId(v string) *UpdateLocationDateClusterShrinkRequest
	GetCustomId() *string
	SetCustomLabelsShrink(v string) *UpdateLocationDateClusterShrinkRequest
	GetCustomLabelsShrink() *string
	SetDatasetName(v string) *UpdateLocationDateClusterShrinkRequest
	GetDatasetName() *string
	SetObjectId(v string) *UpdateLocationDateClusterShrinkRequest
	GetObjectId() *string
	SetProjectName(v string) *UpdateLocationDateClusterShrinkRequest
	GetProjectName() *string
	SetTitle(v string) *UpdateLocationDateClusterShrinkRequest
	GetTitle() *string
}

type UpdateLocationDateClusterShrinkRequest struct {
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
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
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

func (s UpdateLocationDateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocationDateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLocationDateClusterShrinkRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *UpdateLocationDateClusterShrinkRequest) GetCustomLabelsShrink() *string {
	return s.CustomLabelsShrink
}

func (s *UpdateLocationDateClusterShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateLocationDateClusterShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateLocationDateClusterShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateLocationDateClusterShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateLocationDateClusterShrinkRequest) SetCustomId(v string) *UpdateLocationDateClusterShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetCustomLabelsShrink(v string) *UpdateLocationDateClusterShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetDatasetName(v string) *UpdateLocationDateClusterShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetObjectId(v string) *UpdateLocationDateClusterShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetProjectName(v string) *UpdateLocationDateClusterShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) SetTitle(v string) *UpdateLocationDateClusterShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateLocationDateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
