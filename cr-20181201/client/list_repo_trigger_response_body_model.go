// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepoTriggerResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoTriggerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListRepoTriggerResponseBody
	GetRequestId() *string
	SetTriggers(v []*ListRepoTriggerResponseBodyTriggers) *ListRepoTriggerResponseBody
	GetTriggers() []*ListRepoTriggerResponseBodyTriggers
}

type ListRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2CA76D52-A8F0-4D0B-854E-FBD9F6C99049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The triggers of the repository.
	Triggers []*ListRepoTriggerResponseBodyTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListRepoTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoTriggerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoTriggerResponseBody) GetTriggers() []*ListRepoTriggerResponseBodyTriggers {
	return s.Triggers
}

func (s *ListRepoTriggerResponseBody) SetCode(v string) *ListRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetIsSuccess(v bool) *ListRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetRequestId(v string) *ListRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetTriggers(v []*ListRepoTriggerResponseBodyTriggers) *ListRepoTriggerResponseBody {
	s.Triggers = v
	return s
}

func (s *ListRepoTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRepoTriggerResponseBodyTriggers struct {
	// The type of the event that activates the trigger. Valid values:
	//
	// 	- `BUILD_SUCCESS`: The trigger is activated when an image building task is successful.
	//
	// 	- `BUILD_Fail`: The trigger is activated when an image building task fails.
	//
	// example:
	//
	// BUILD_SUCCESS
	RepoEvent *string `json:"RepoEvent,omitempty" xml:"RepoEvent,omitempty"`
	// The ID of the trigger.
	//
	// example:
	//
	// crw-vriyql9eq7ep****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	//
	// example:
	//
	// test
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// example:
	//
	// *
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LISTTAG`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// example:
	//
	// ALL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// example:
	//
	// https://www.test.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s ListRepoTriggerResponseBodyTriggers) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTriggerResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponseBodyTriggers) GetRepoEvent() *string {
	return s.RepoEvent
}

func (s *ListRepoTriggerResponseBodyTriggers) GetTriggerId() *string {
	return s.TriggerId
}

func (s *ListRepoTriggerResponseBodyTriggers) GetTriggerName() *string {
	return s.TriggerName
}

func (s *ListRepoTriggerResponseBodyTriggers) GetTriggerTag() *string {
	return s.TriggerTag
}

func (s *ListRepoTriggerResponseBodyTriggers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListRepoTriggerResponseBodyTriggers) GetTriggerUrl() *string {
	return s.TriggerUrl
}

func (s *ListRepoTriggerResponseBodyTriggers) SetRepoEvent(v string) *ListRepoTriggerResponseBodyTriggers {
	s.RepoEvent = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerId(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerId = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerName(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerName = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerTag(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerTag = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerType(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerType = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerUrl(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerUrl = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) Validate() error {
	return dara.Validate(s)
}
