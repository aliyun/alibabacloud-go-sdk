// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressShrink(v string) *CreateStoryShrinkRequest
	GetAddressShrink() *string
	SetCustomId(v string) *CreateStoryShrinkRequest
	GetCustomId() *string
	SetCustomLabelsShrink(v string) *CreateStoryShrinkRequest
	GetCustomLabelsShrink() *string
	SetDatasetName(v string) *CreateStoryShrinkRequest
	GetDatasetName() *string
	SetMaxFileCount(v int64) *CreateStoryShrinkRequest
	GetMaxFileCount() *int64
	SetMinFileCount(v int64) *CreateStoryShrinkRequest
	GetMinFileCount() *int64
	SetNotificationShrink(v string) *CreateStoryShrinkRequest
	GetNotificationShrink() *string
	SetNotifyTopicName(v string) *CreateStoryShrinkRequest
	GetNotifyTopicName() *string
	SetObjectId(v string) *CreateStoryShrinkRequest
	GetObjectId() *string
	SetProjectName(v string) *CreateStoryShrinkRequest
	GetProjectName() *string
	SetStoryEndTime(v string) *CreateStoryShrinkRequest
	GetStoryEndTime() *string
	SetStoryName(v string) *CreateStoryShrinkRequest
	GetStoryName() *string
	SetStoryStartTime(v string) *CreateStoryShrinkRequest
	GetStoryStartTime() *string
	SetStorySubType(v string) *CreateStoryShrinkRequest
	GetStorySubType() *string
	SetStoryType(v string) *CreateStoryShrinkRequest
	GetStoryType() *string
	SetTagsShrink(v string) *CreateStoryShrinkRequest
	GetTagsShrink() *string
	SetUserData(v string) *CreateStoryShrinkRequest
	GetUserData() *string
}

type CreateStoryShrinkRequest struct {
	// The address information for the story. IMM filters photos whose shooting locations match the specified address to generate the story. This parameter takes effect only when StoryType is set to TravelMemory.
	//
	// > Due to regulatory requirements, parsing GPS information into addresses is not supported in Hong Kong (China), Macao (China), Taiwan (China), or regions outside of mainland China.
	AddressShrink *string `json:"Address,omitempty" xml:"Address,omitempty"`
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
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
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
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information that is returned in an asynchronous notification message. You can use this information to associate the notification message with your services. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateStoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryShrinkRequest) GetAddressShrink() *string {
	return s.AddressShrink
}

func (s *CreateStoryShrinkRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *CreateStoryShrinkRequest) GetCustomLabelsShrink() *string {
	return s.CustomLabelsShrink
}

func (s *CreateStoryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateStoryShrinkRequest) GetMaxFileCount() *int64 {
	return s.MaxFileCount
}

func (s *CreateStoryShrinkRequest) GetMinFileCount() *int64 {
	return s.MinFileCount
}

func (s *CreateStoryShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateStoryShrinkRequest) GetNotifyTopicName() *string {
	return s.NotifyTopicName
}

func (s *CreateStoryShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateStoryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateStoryShrinkRequest) GetStoryEndTime() *string {
	return s.StoryEndTime
}

func (s *CreateStoryShrinkRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *CreateStoryShrinkRequest) GetStoryStartTime() *string {
	return s.StoryStartTime
}

func (s *CreateStoryShrinkRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *CreateStoryShrinkRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *CreateStoryShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateStoryShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateStoryShrinkRequest) SetAddressShrink(v string) *CreateStoryShrinkRequest {
	s.AddressShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetCustomId(v string) *CreateStoryShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetCustomLabelsShrink(v string) *CreateStoryShrinkRequest {
	s.CustomLabelsShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetDatasetName(v string) *CreateStoryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetMaxFileCount(v int64) *CreateStoryShrinkRequest {
	s.MaxFileCount = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetMinFileCount(v int64) *CreateStoryShrinkRequest {
	s.MinFileCount = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetNotificationShrink(v string) *CreateStoryShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetNotifyTopicName(v string) *CreateStoryShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetObjectId(v string) *CreateStoryShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetProjectName(v string) *CreateStoryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryEndTime(v string) *CreateStoryShrinkRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryName(v string) *CreateStoryShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryStartTime(v string) *CreateStoryShrinkRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStorySubType(v string) *CreateStoryShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetStoryType(v string) *CreateStoryShrinkRequest {
	s.StoryType = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetTagsShrink(v string) *CreateStoryShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateStoryShrinkRequest) SetUserData(v string) *CreateStoryShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateStoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
