// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v *AddressForStory) *CreateStoryRequest
	GetAddress() *AddressForStory
	SetCustomId(v string) *CreateStoryRequest
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *CreateStoryRequest
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *CreateStoryRequest
	GetDatasetName() *string
	SetMaxFileCount(v int64) *CreateStoryRequest
	GetMaxFileCount() *int64
	SetMinFileCount(v int64) *CreateStoryRequest
	GetMinFileCount() *int64
	SetNotification(v *Notification) *CreateStoryRequest
	GetNotification() *Notification
	SetNotifyTopicName(v string) *CreateStoryRequest
	GetNotifyTopicName() *string
	SetObjectId(v string) *CreateStoryRequest
	GetObjectId() *string
	SetProjectName(v string) *CreateStoryRequest
	GetProjectName() *string
	SetStoryEndTime(v string) *CreateStoryRequest
	GetStoryEndTime() *string
	SetStoryName(v string) *CreateStoryRequest
	GetStoryName() *string
	SetStoryStartTime(v string) *CreateStoryRequest
	GetStoryStartTime() *string
	SetStorySubType(v string) *CreateStoryRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateStoryRequest
	GetStoryType() *string
	SetTags(v map[string]interface{}) *CreateStoryRequest
	GetTags() map[string]interface{}
	SetUserData(v string) *CreateStoryRequest
	GetUserData() *string
}

type CreateStoryRequest struct {
	// The address information for the story. IMM filters photos whose shooting locations match the specified address to generate the story. This parameter takes effect only when StoryType is set to TravelMemory.
	//
	// > Due to regulatory requirements, parsing GPS information into addresses is not supported in Hong Kong (China), Macao (China), Taiwan (China), or regions outside of mainland China.
	Address *AddressForStory `json:"Address,omitempty" xml:"Address,omitempty"`
	// A custom identifier for the story. This ID can be different from ObjectId. You can use this ID to retrieve or sort stories.
	//
	// example:
	//
	// test
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels. These labels contain custom information about the story and can be used for retrieval.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of photos in the generated story. The actual number of photos is between the values of MinFileCount and MaxFileCount. The value must be an integer greater than the value of MinFileCount. To ensure the quality of the generated story, the internal algorithm limits the maximum number of photos to 1,500. If you set MaxFileCount to a value greater than 1,500, the setting does not take effect.
	//
	// example:
	//
	// 3
	MaxFileCount *int64 `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	// The minimum number of photos in the generated story. The actual number of photos is between the values of MinFileCount and MaxFileCount. The value must be an integer greater than 1. If the number of available candidate photos is less than this value, an empty story is returned.
	//
	// example:
	//
	// 1
	MinFileCount *int64 `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	// The notification configuration. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the topic for asynchronous notifications.
	//
	// example:
	//
	// test-topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// The ID for the story object. This parameter is optional. If you do not specify an ID, IMM generates one. You can use the story ID to query or update the story. If you specify an ID that already exists, the corresponding story is updated.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The end time of the photo collection for the story. This parameter and StoryStartTime define a time range. IMM filters candidate photos within this time range to generate the story. The value must be a string in the RFC 3339 format.
	//
	// example:
	//
	// 2021-12-30T16:00:00Z
	StoryEndTime *string `json:"StoryEndTime,omitempty" xml:"StoryEndTime,omitempty"`
	// The name of the story.
	//
	// example:
	//
	// name1
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	// The start time of the photo collection for the story. This parameter and StoryEndTime define a time range. IMM filters candidate photos within this time range to generate the story. The value must be a string in the RFC 3339 format.
	//
	// example:
	//
	// 2016-12-30T16:00:00Z
	StoryStartTime *string `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	// The subtype of the story to generate. For more information about story subtypes and their valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// Solo
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The type of the story to generate. For more information about story types and their valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	// This parameter provides a tagging mechanism that can be used in the following scenarios:
	//
	// - Set custom data that is returned in MNS messages.
	//
	// - Use as a search condition to search for tasks.
	//
	// - Use as a variable in TargetURI.
	//
	// example:
	//
	// {"key":"val"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information that is returned in an asynchronous notification message. You can use this information to associate the notification message with your services. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryRequest) GetAddress() *AddressForStory {
	return s.Address
}

func (s *CreateStoryRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *CreateStoryRequest) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *CreateStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateStoryRequest) GetMaxFileCount() *int64 {
	return s.MaxFileCount
}

func (s *CreateStoryRequest) GetMinFileCount() *int64 {
	return s.MinFileCount
}

func (s *CreateStoryRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateStoryRequest) GetNotifyTopicName() *string {
	return s.NotifyTopicName
}

func (s *CreateStoryRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateStoryRequest) GetStoryEndTime() *string {
	return s.StoryEndTime
}

func (s *CreateStoryRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *CreateStoryRequest) GetStoryStartTime() *string {
	return s.StoryStartTime
}

func (s *CreateStoryRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *CreateStoryRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *CreateStoryRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateStoryRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateStoryRequest) SetAddress(v *AddressForStory) *CreateStoryRequest {
	s.Address = v
	return s
}

func (s *CreateStoryRequest) SetCustomId(v string) *CreateStoryRequest {
	s.CustomId = &v
	return s
}

func (s *CreateStoryRequest) SetCustomLabels(v map[string]interface{}) *CreateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateStoryRequest) SetDatasetName(v string) *CreateStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateStoryRequest) SetMaxFileCount(v int64) *CreateStoryRequest {
	s.MaxFileCount = &v
	return s
}

func (s *CreateStoryRequest) SetMinFileCount(v int64) *CreateStoryRequest {
	s.MinFileCount = &v
	return s
}

func (s *CreateStoryRequest) SetNotification(v *Notification) *CreateStoryRequest {
	s.Notification = v
	return s
}

func (s *CreateStoryRequest) SetNotifyTopicName(v string) *CreateStoryRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateStoryRequest) SetObjectId(v string) *CreateStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateStoryRequest) SetProjectName(v string) *CreateStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryEndTime(v string) *CreateStoryRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryRequest) SetStoryName(v string) *CreateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryStartTime(v string) *CreateStoryRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryRequest) SetStorySubType(v string) *CreateStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryRequest) SetStoryType(v string) *CreateStoryRequest {
	s.StoryType = &v
	return s
}

func (s *CreateStoryRequest) SetTags(v map[string]interface{}) *CreateStoryRequest {
	s.Tags = v
	return s
}

func (s *CreateStoryRequest) SetUserData(v string) *CreateStoryRequest {
	s.UserData = &v
	return s
}

func (s *CreateStoryRequest) Validate() error {
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
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
