// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *GetEventCallbackResponseBody
	GetAuthKey() *string
	SetAuthSwitch(v string) *GetEventCallbackResponseBody
	GetAuthSwitch() *string
	SetCallbackQueueName(v string) *GetEventCallbackResponseBody
	GetCallbackQueueName() *string
	SetCallbackType(v string) *GetEventCallbackResponseBody
	GetCallbackType() *string
	SetCallbackURL(v string) *GetEventCallbackResponseBody
	GetCallbackURL() *string
	SetEventTypeList(v string) *GetEventCallbackResponseBody
	GetEventTypeList() *string
	SetRequestId(v string) *GetEventCallbackResponseBody
	GetRequestId() *string
}

type GetEventCallbackResponseBody struct {
	// The authentication key. This parameter is returned only for HTTP callbacks.
	//
	// example:
	//
	// TestKey001
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// Specifies whether callback authentication is enabled. This parameter is returned only for **HTTP*	- callbacks. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AuthSwitch *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	// The name of the Simple Message Queue (SMQ) queue to which callback messages are sent.
	//
	// example:
	//
	// ice-callback-queue
	CallbackQueueName *string `json:"CallbackQueueName,omitempty" xml:"CallbackQueueName,omitempty"`
	// The callback method. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **MNS**
	//
	// example:
	//
	// HTTP
	CallbackType *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	// The callback URL to which event notifications are sent.
	//
	// example:
	//
	// http://xxx.yyy/callback
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The type of the callback event. Multiple values are separated with commas (,). For more information about callback event types, see [Event notification content](https://help.aliyun.com/document_detail/610204.html).
	//
	// example:
	//
	// ProduceMediaComplete,TranscodeComplete
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEventCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventCallbackResponseBody) GetAuthKey() *string {
	return s.AuthKey
}

func (s *GetEventCallbackResponseBody) GetAuthSwitch() *string {
	return s.AuthSwitch
}

func (s *GetEventCallbackResponseBody) GetCallbackQueueName() *string {
	return s.CallbackQueueName
}

func (s *GetEventCallbackResponseBody) GetCallbackType() *string {
	return s.CallbackType
}

func (s *GetEventCallbackResponseBody) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *GetEventCallbackResponseBody) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *GetEventCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventCallbackResponseBody) SetAuthKey(v string) *GetEventCallbackResponseBody {
	s.AuthKey = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetAuthSwitch(v string) *GetEventCallbackResponseBody {
	s.AuthSwitch = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetCallbackQueueName(v string) *GetEventCallbackResponseBody {
	s.CallbackQueueName = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetCallbackType(v string) *GetEventCallbackResponseBody {
	s.CallbackType = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetCallbackURL(v string) *GetEventCallbackResponseBody {
	s.CallbackURL = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetEventTypeList(v string) *GetEventCallbackResponseBody {
	s.EventTypeList = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetRequestId(v string) *GetEventCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
