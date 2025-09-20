// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationDateClusteringTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *CreateLocationDateClusteringTaskRequest
	GetDatasetName() *string
	SetDateOptions(v *CreateLocationDateClusteringTaskRequestDateOptions) *CreateLocationDateClusteringTaskRequest
	GetDateOptions() *CreateLocationDateClusteringTaskRequestDateOptions
	SetLocationOptions(v *CreateLocationDateClusteringTaskRequestLocationOptions) *CreateLocationDateClusteringTaskRequest
	GetLocationOptions() *CreateLocationDateClusteringTaskRequestLocationOptions
	SetNotification(v *Notification) *CreateLocationDateClusteringTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateLocationDateClusteringTaskRequest
	GetProjectName() *string
	SetTags(v map[string]interface{}) *CreateLocationDateClusteringTaskRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateLocationDateClusteringTaskRequest
	GetUserData() *string
}

type CreateLocationDateClusteringTaskRequest struct {
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
	DateOptions *CreateLocationDateClusteringTaskRequestDateOptions `json:"DateOptions,omitempty" xml:"DateOptions,omitempty" type:"Struct"`
	// The geolocation configurations for clustering.
	//
	// >  Adjusting these configurations affects existing spatiotemporal clusters for the dataset.
	//
	// This parameter is required.
	LocationOptions *CreateLocationDateClusteringTaskRequestLocationOptions `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty" type:"Struct"`
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

func (s CreateLocationDateClusteringTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateLocationDateClusteringTaskRequest) GetDateOptions() *CreateLocationDateClusteringTaskRequestDateOptions {
	return s.DateOptions
}

func (s *CreateLocationDateClusteringTaskRequest) GetLocationOptions() *CreateLocationDateClusteringTaskRequestLocationOptions {
	return s.LocationOptions
}

func (s *CreateLocationDateClusteringTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateLocationDateClusteringTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateLocationDateClusteringTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateLocationDateClusteringTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLocationDateClusteringTaskRequest) SetDatasetName(v string) *CreateLocationDateClusteringTaskRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetDateOptions(v *CreateLocationDateClusteringTaskRequestDateOptions) *CreateLocationDateClusteringTaskRequest {
	s.DateOptions = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetLocationOptions(v *CreateLocationDateClusteringTaskRequestLocationOptions) *CreateLocationDateClusteringTaskRequest {
	s.LocationOptions = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetNotification(v *Notification) *CreateLocationDateClusteringTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetProjectName(v string) *CreateLocationDateClusteringTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetTags(v map[string]interface{}) *CreateLocationDateClusteringTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) SetUserData(v string) *CreateLocationDateClusteringTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateLocationDateClusteringTaskRequestDateOptions struct {
	// The maximum number of days allowed in a gap for a single spatiotemporal cluster. Valid values: 0 to 99999.
	//
	// For example, if travel photos were produced on March 4, 5, and 7, 2024, but not on Marh 6, 2024, and you set the parameter to 1, IMM considers the travel spanning the date range from March 4, 2024 to March 7, 2024 and includes photos within the data range in the same cluster.````
	//
	// We recommend that you set the parameter to a value within the range from 0 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GapDays *int64 `json:"GapDays,omitempty" xml:"GapDays,omitempty"`
	// The maximum number of days that a single spatiotemporal cluster can span. Valid values: 1 to 99999. IMM does not create a cluster that spans more than the maximum number of days.
	//
	// For example, if you want to create travel photo clusters, you may want to exclude photos that were taken within 15 consecutive days in the same city, because it is likely that these photos were not taken during a travel. In this case, you can set the parameter to 15 to exclude this time range and location from the clustering task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	MaxDays *int64 `json:"MaxDays,omitempty" xml:"MaxDays,omitempty"`
	// The minimum number of days that a single spatiotemporal cluster can span. Valid values: 1 to 99999. IMM does not create a cluster that spans less than the minimum number of days.
	//
	// For example, if you do not want a one-day tour cluster, you can set the parameter to 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinDays *int64 `json:"MinDays,omitempty" xml:"MinDays,omitempty"`
}

func (s CreateLocationDateClusteringTaskRequestDateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequestDateOptions) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) GetGapDays() *int64 {
	return s.GapDays
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) GetMaxDays() *int64 {
	return s.MaxDays
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) GetMinDays() *int64 {
	return s.MinDays
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetGapDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.GapDays = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetMaxDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.MaxDays = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) SetMinDays(v int64) *CreateLocationDateClusteringTaskRequestDateOptions {
	s.MinDays = &v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestDateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLocationDateClusteringTaskRequestLocationOptions struct {
	// The administrative division levels. You can specify multiple administrative division levels.
	//
	// For example, you uploaded photos that were taken from March 3, 2024 to March 5, 2024 in Hangzhou and photos that were taken from March 6, 2024 to March 8, 2024 in Jiaxing. When you call the operation and set the parameter to `["city", "province"]`, the following spatiotemporal clusters are created from these photos:
	//
	// 	- March 3, 2024 to March 5, 2024, Hangzhou
	//
	// 	- March 6, 2024 to March 8, 2024, Jiaxing
	//
	// 	- March 3, 2024 to March 8, 2024, Zhejiang
	//
	// This parameter is required.
	LocationDateClusterLevels []*string `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty" type:"Repeated"`
}

func (s CreateLocationDateClusteringTaskRequestLocationOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskRequestLocationOptions) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskRequestLocationOptions) GetLocationDateClusterLevels() []*string {
	return s.LocationDateClusterLevels
}

func (s *CreateLocationDateClusteringTaskRequestLocationOptions) SetLocationDateClusterLevels(v []*string) *CreateLocationDateClusteringTaskRequestLocationOptions {
	s.LocationDateClusterLevels = v
	return s
}

func (s *CreateLocationDateClusteringTaskRequestLocationOptions) Validate() error {
	return dara.Validate(s)
}
