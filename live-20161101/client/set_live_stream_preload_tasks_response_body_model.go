// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamPreloadTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedURL(v int32) *SetLiveStreamPreloadTasksResponseBody
	GetFailedURL() *int32
	SetPreloadTasksMessages(v *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) *SetLiveStreamPreloadTasksResponseBody
	GetPreloadTasksMessages() *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages
	SetRequestId(v string) *SetLiveStreamPreloadTasksResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetLiveStreamPreloadTasksResponseBody
	GetStatus() *string
	SetSuccessURL(v int32) *SetLiveStreamPreloadTasksResponseBody
	GetSuccessURL() *int32
	SetTotalURL(v int32) *SetLiveStreamPreloadTasksResponseBody
	GetTotalURL() *int32
}

type SetLiveStreamPreloadTasksResponseBody struct {
	// The number of URLs for which the prefetch task configuration failed.
	//
	// example:
	//
	// 0
	FailedURL *int32 `json:"FailedURL,omitempty" xml:"FailedURL,omitempty"`
	// The details of the prefetch task.
	PreloadTasksMessages *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages `json:"PreloadTasksMessages,omitempty" xml:"PreloadTasksMessages,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36E0E523-E0C6-5D95-A465-A8EA2DCBA2A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the prefetch task. Valid values:
	//
	// 	- Success
	//
	// 	- Failed
	//
	// >  Success is returned only if the prefetch task is configured for all specified streaming URLs.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of URLs for which the prefetch task is successfully configured.
	//
	// example:
	//
	// 1
	SuccessURL *int32 `json:"SuccessURL,omitempty" xml:"SuccessURL,omitempty"`
	// The total number of URLs.
	//
	// example:
	//
	// 1
	TotalURL *int32 `json:"TotalURL,omitempty" xml:"TotalURL,omitempty"`
}

func (s SetLiveStreamPreloadTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamPreloadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetFailedURL() *int32 {
	return s.FailedURL
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetPreloadTasksMessages() *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages {
	return s.PreloadTasksMessages
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetSuccessURL() *int32 {
	return s.SuccessURL
}

func (s *SetLiveStreamPreloadTasksResponseBody) GetTotalURL() *int32 {
	return s.TotalURL
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetFailedURL(v int32) *SetLiveStreamPreloadTasksResponseBody {
	s.FailedURL = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetPreloadTasksMessages(v *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) *SetLiveStreamPreloadTasksResponseBody {
	s.PreloadTasksMessages = v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetRequestId(v string) *SetLiveStreamPreloadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetStatus(v string) *SetLiveStreamPreloadTasksResponseBody {
	s.Status = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetSuccessURL(v int32) *SetLiveStreamPreloadTasksResponseBody {
	s.SuccessURL = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) SetTotalURL(v int32) *SetLiveStreamPreloadTasksResponseBody {
	s.TotalURL = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages struct {
	PreloadTasksMessage []*SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage `json:"PreloadTasksMessage,omitempty" xml:"PreloadTasksMessage,omitempty" type:"Repeated"`
}

func (s SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) GoString() string {
	return s.String()
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) GetPreloadTasksMessage() []*SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage {
	return s.PreloadTasksMessage
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) SetPreloadTasksMessage(v []*SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages {
	s.PreloadTasksMessage = v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessages) Validate() error {
	return dara.Validate(s)
}

type SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage struct {
	// Indicates whether the prefetch task is successful. Valid values:
	//
	// 	- Successfully
	//
	// 	- InternalError
	//
	// example:
	//
	// Successfully
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming URL.
	PlayUrl *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	// The ID of the prefetch task.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) GoString() string {
	return s.String()
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) GetDescription() *string {
	return s.Description
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) GetTaskId() *string {
	return s.TaskId
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) SetDescription(v string) *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage {
	s.Description = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) SetPlayUrl(v string) *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage {
	s.PlayUrl = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) SetTaskId(v string) *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage {
	s.TaskId = &v
	return s
}

func (s *SetLiveStreamPreloadTasksResponseBodyPreloadTasksMessagesPreloadTasksMessage) Validate() error {
	return dara.Validate(s)
}
