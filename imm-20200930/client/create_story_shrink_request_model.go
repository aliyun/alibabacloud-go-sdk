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
	// The address of the story. IMM filters candidate photos to generate a story based on the value of this parameter. This parameter takes effect only if you set StoryType to TravelMemory.
	//
	// >  If you are located in Hong Kong (China), Macao (China), Taiwan (China), or overseas, you cannot specify an address in the Chinese mainland by using this parameter.
	AddressShrink *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The custom ID. A custom ID of a generated story may differ from the value of ObjectID and can be utilized for subsequent retrieval and sorting of stories.
	//
	// example:
	//
	// test
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels. Labels specify the custom information of the story. This enables retrieval based on your business requirements.
	//
	// example:
	//
	// {"Bucket": "examplebucket"}
	CustomLabelsShrink *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For information about how to obtain the name of a dataset, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of photo files in the story. The actual number of photo files ranges from the value of MinFileCount to the value of MaxFileCount. The value of this parameter must be an integer greater than the value of MinFileCount. To provide the desired effect, the algorithm limits the maximum number of photo files to 1,500. If you set MaxFileCount to a value greater than 1,500, this parameter does not take effect.
	//
	// example:
	//
	// 3
	MaxFileCount *int64 `json:"MaxFileCount,omitempty" xml:"MaxFileCount,omitempty"`
	// The minimum number of photo files in the story. The actual number of photo files ranges from the value of MinFileCount to the value of MaxFileCount. The value of this parameter must be an integer greater than 1. If the actual number of candidate photos is less than the value of this parameter, a null story is returned.
	//
	// example:
	//
	// 1
	MinFileCount *int64 `json:"MinFileCount,omitempty" xml:"MinFileCount,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The topic name of the asynchronous reverse notification.
	//
	// example:
	//
	// test-topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// The ID of the story. This parameter is optional. If you leave this parameter empty, IMM assigns a unique identifier to the story. You can query and update a story based on its ID. You can also manually create an ID for a story. After you create an ID for a story, you must specify this parameter to pass the ID into the system. This way, IMM can record the ID as the unique identifier of the story. If you pass an existing ID into the system, IMM updates the story that corresponds to the ID.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The end time of the photo collection for which you want to create the story. StoryStartTime and StoryEndTime form a time interval based on which IMM filters candidate photos to generate a story. The value must be a string in the RFC3339 format.
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
	// The start time of the photo collection for which you want to create the story. StoryStartTime and StoryEndTime form a time interval based on which IMM filters candidate photos to generate a story. The value must be a string in the RFC3339 format.
	//
	// example:
	//
	// 2016-12-30T16:00:00Z
	StoryStartTime *string `json:"StoryStartTime,omitempty" xml:"StoryStartTime,omitempty"`
	// The subtype of the story. For information about valid subtypes, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// Solo
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The type of the story. For information about valid types, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	// The tags. You can specify this parameter in one of the following scenarios:
	//
	// 	- Specify tags as custom data, which is returned in messages provided by Simple Message Queue.
	//
	// 	- Search for tasks by tag.
	//
	// 	- Specify tags as variables in destination URIs.
	//
	// example:
	//
	// {"key":"val"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The custom information, which is returned as asynchronous notifications to facilitate notification management in your system. The maximum information length is 2,048 bytes.
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
