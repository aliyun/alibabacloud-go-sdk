// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpEventConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHttpEventConfigResponseBody
	GetCode() *string
	SetData(v *GetHttpEventConfigResponseBodyData) *GetHttpEventConfigResponseBody
	GetData() *GetHttpEventConfigResponseBodyData
	SetMessage(v string) *GetHttpEventConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHttpEventConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHttpEventConfigResponseBody
	GetSuccess() *bool
}

type GetHttpEventConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHttpEventConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHttpEventConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpEventConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpEventConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHttpEventConfigResponseBody) GetData() *GetHttpEventConfigResponseBodyData {
	return s.Data
}

func (s *GetHttpEventConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHttpEventConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpEventConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHttpEventConfigResponseBody) SetCode(v string) *GetHttpEventConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetHttpEventConfigResponseBody) SetData(v *GetHttpEventConfigResponseBodyData) *GetHttpEventConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetHttpEventConfigResponseBody) SetMessage(v string) *GetHttpEventConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetHttpEventConfigResponseBody) SetRequestId(v string) *GetHttpEventConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpEventConfigResponseBody) SetSuccess(v bool) *GetHttpEventConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetHttpEventConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHttpEventConfigResponseBodyData struct {
	// example:
	//
	// http://123213.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// "eventConfig": { "taskCompleteEvent":true}
	EventConfig *GetHttpEventConfigResponseBodyDataEventConfig `json:"EventConfig,omitempty" xml:"EventConfig,omitempty" type:"Struct"`
	// example:
	//
	// GET
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// example:
	//
	// true
	SwitchFlag *bool `json:"SwitchFlag,omitempty" xml:"SwitchFlag,omitempty"`
}

func (s GetHttpEventConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHttpEventConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHttpEventConfigResponseBodyData) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *GetHttpEventConfigResponseBodyData) GetEventConfig() *GetHttpEventConfigResponseBodyDataEventConfig {
	return s.EventConfig
}

func (s *GetHttpEventConfigResponseBodyData) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *GetHttpEventConfigResponseBodyData) GetSwitchFlag() *bool {
	return s.SwitchFlag
}

func (s *GetHttpEventConfigResponseBodyData) SetCallbackUrl(v string) *GetHttpEventConfigResponseBodyData {
	s.CallbackUrl = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyData) SetEventConfig(v *GetHttpEventConfigResponseBodyDataEventConfig) *GetHttpEventConfigResponseBodyData {
	s.EventConfig = v
	return s
}

func (s *GetHttpEventConfigResponseBodyData) SetHttpMethod(v string) *GetHttpEventConfigResponseBodyData {
	s.HttpMethod = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyData) SetSwitchFlag(v bool) *GetHttpEventConfigResponseBodyData {
	s.SwitchFlag = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyData) Validate() error {
	if s.EventConfig != nil {
		if err := s.EventConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHttpEventConfigResponseBodyDataEventConfig struct {
	// example:
	//
	// true
	ComplaintHandledEvent *bool `json:"ComplaintHandledEvent,omitempty" xml:"ComplaintHandledEvent,omitempty"`
	// example:
	//
	// true
	FileReviewedEvent *bool `json:"FileReviewedEvent,omitempty" xml:"FileReviewedEvent,omitempty"`
	// example:
	//
	// true
	SyncTaskCompleteEvent *bool `json:"SyncTaskCompleteEvent,omitempty" xml:"SyncTaskCompleteEvent,omitempty"`
	// example:
	//
	// true
	TaskCompleteEvent *bool `json:"TaskCompleteEvent,omitempty" xml:"TaskCompleteEvent,omitempty"`
}

func (s GetHttpEventConfigResponseBodyDataEventConfig) String() string {
	return dara.Prettify(s)
}

func (s GetHttpEventConfigResponseBodyDataEventConfig) GoString() string {
	return s.String()
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) GetComplaintHandledEvent() *bool {
	return s.ComplaintHandledEvent
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) GetFileReviewedEvent() *bool {
	return s.FileReviewedEvent
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) GetSyncTaskCompleteEvent() *bool {
	return s.SyncTaskCompleteEvent
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) GetTaskCompleteEvent() *bool {
	return s.TaskCompleteEvent
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) SetComplaintHandledEvent(v bool) *GetHttpEventConfigResponseBodyDataEventConfig {
	s.ComplaintHandledEvent = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) SetFileReviewedEvent(v bool) *GetHttpEventConfigResponseBodyDataEventConfig {
	s.FileReviewedEvent = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) SetSyncTaskCompleteEvent(v bool) *GetHttpEventConfigResponseBodyDataEventConfig {
	s.SyncTaskCompleteEvent = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) SetTaskCompleteEvent(v bool) *GetHttpEventConfigResponseBodyDataEventConfig {
	s.TaskCompleteEvent = &v
	return s
}

func (s *GetHttpEventConfigResponseBodyDataEventConfig) Validate() error {
	return dara.Validate(s)
}
