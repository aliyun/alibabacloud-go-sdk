// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationDateClusteringTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetDatasetName() *string
	SetDateOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetDateOptionsShrink() *string
	SetLocationOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetLocationOptionsShrink() *string
	SetNotificationShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetNotificationShrink() *string
	SetProjectName(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetProjectName() *string
	SetTagsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateLocationDateClusteringTaskShrinkRequest
	GetUserData() *string
}

type CreateLocationDateClusteringTaskShrinkRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The date configurations for clustering.
	//
	// >  Adjusting these configurations affects existing spatiotemporal clusters for the dataset.
	//
	// This parameter is required.
	DateOptionsShrink *string `json:"DateOptions,omitempty" xml:"DateOptions,omitempty"`
	// The geolocation configurations for clustering.
	//
	// >  Adjusting these configurations affects existing spatiotemporal clusters for the dataset.
	//
	// This parameter is required.
	LocationOptionsShrink *string `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// test-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateLocationDateClusteringTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetDateOptionsShrink() *string {
	return s.DateOptionsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetLocationOptionsShrink() *string {
	return s.LocationOptionsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDatasetName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetDateOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.DateOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetLocationOptionsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.LocationOptionsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetNotificationShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetProjectName(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetTagsShrink(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) SetUserData(v string) *CreateLocationDateClusteringTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateLocationDateClusteringTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
