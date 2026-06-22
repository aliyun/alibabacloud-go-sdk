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
	// The dataset name. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The date clustering settings.
	//
	// 	Notice: Modifying this setting also affects existing spatio-temporal clusters in your `Dataset`.
	//
	// This parameter is required.
	DateOptions *CreateLocationDateClusteringTaskRequestDateOptions `json:"DateOptions,omitempty" xml:"DateOptions,omitempty" type:"Struct"`
	// The location clustering settings.
	//
	// 	Notice: Modifying this setting also affects existing spatio-temporal clusters in your `Dataset`.
	//
	// This parameter is required.
	LocationOptions *CreateLocationDateClusteringTaskRequestLocationOptions `json:"LocationOptions,omitempty" xml:"LocationOptions,omitempty" type:"Struct"`
	// The message notification configuration. For more information, see Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Custom tags used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Custom information that is returned in the asynchronous notification message. This helps you associate the notification message with your system. The maximum length is 2,048 bytes.
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
	if s.DateOptions != nil {
		if err := s.DateOptions.Validate(); err != nil {
			return err
		}
	}
	if s.LocationOptions != nil {
		if err := s.LocationOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLocationDateClusteringTaskRequestDateOptions struct {
	// The maximum number of gap days allowed in a single spatio-temporal group. The value must be in the range of 0 to 99,999.
	//
	// For example, a user has photos from March 4–5 and March 7, but not from March 6. If you assume that the photos from March 4–7 belong to the same trip, set this parameter to `1 day`. This allows the gap of `1 day` on March 6 to be included in the same spatio-temporal cluster.
	//
	// Set this parameter to a value from 0 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GapDays *int64 `json:"GapDays,omitempty" xml:"GapDays,omitempty"`
	// The maximum number of days in a single spatio-temporal group. The value must be in the range of 1 to 99,999. Clusters with more days than this value are not detected or stored.
	//
	// For example, if a user takes photos in the same location for more than 15 consecutive days, this location might be their residence rather than a travel destination. If you want to exclude this time period and location from the spatio-temporal clusters, set this parameter to 15.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	MaxDays *int64 `json:"MaxDays,omitempty" xml:"MaxDays,omitempty"`
	// The minimum number of days in a single spatio-temporal group. The value must be in the range of 1 to 99,999. Clusters with fewer days than this value are not detected or stored.
	//
	// For example, if you do not want to include one-day trips in the generated groups, set this parameter to 2.
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
	// A list of administrative levels for grouping. You can select multiple levels.
	//
	// For example, a user uploads photos taken in Hangzhou from March 3 to March 5 and photos taken in Jiaxing from March 6 to March 8. If you set this parameter to `["city", "province"]`, the following spatio-temporal clusters are generated:
	//
	// - March 3 to March 5, Hangzhou
	//
	// - March 6 to March 8, Jiaxing
	//
	// - March 3 to March 8, Zhejiang
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
