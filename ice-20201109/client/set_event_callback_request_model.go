// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *SetEventCallbackRequest
	GetAuthKey() *string
	SetAuthSwitch(v string) *SetEventCallbackRequest
	GetAuthSwitch() *string
	SetCallbackQueueName(v string) *SetEventCallbackRequest
	GetCallbackQueueName() *string
	SetCallbackType(v string) *SetEventCallbackRequest
	GetCallbackType() *string
	SetCallbackURL(v string) *SetEventCallbackRequest
	GetCallbackURL() *string
	SetEventTypeList(v string) *SetEventCallbackRequest
	GetEventTypeList() *string
}

type SetEventCallbackRequest struct {
	// The authentication key. The key can be up to 32 characters in length and must contain uppercase letters, lowercase letters, and digits. This parameter takes effect only if you set CallbackType to **HTTP**.
	//
	// example:
	//
	// TestKey001
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// Specifies whether to enable callback authentication. This parameter takes effect only if you set CallbackType to **HTTP**. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AuthSwitch *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	// The name of the Simple Message Queue (SMQ) queue in the region. The name must start with ice-callback-.
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
	// The callback URL. This parameter is required if you set CallbackType to **HTTP**. The callback URL cannot exceed 256 bytes in length. You can specify only one callback URL.
	//
	// example:
	//
	// http://xxx.yyy/callback
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The type of the callback event. You can specify multiple values separated with commas (,). ProduceMediaComplete: indicates that the editing and production task is complete.
	//
	// example:
	//
	// ProduceMediaComplete
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
}

func (s SetEventCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEventCallbackRequest) GoString() string {
	return s.String()
}

func (s *SetEventCallbackRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SetEventCallbackRequest) GetAuthSwitch() *string {
	return s.AuthSwitch
}

func (s *SetEventCallbackRequest) GetCallbackQueueName() *string {
	return s.CallbackQueueName
}

func (s *SetEventCallbackRequest) GetCallbackType() *string {
	return s.CallbackType
}

func (s *SetEventCallbackRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *SetEventCallbackRequest) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *SetEventCallbackRequest) SetAuthKey(v string) *SetEventCallbackRequest {
	s.AuthKey = &v
	return s
}

func (s *SetEventCallbackRequest) SetAuthSwitch(v string) *SetEventCallbackRequest {
	s.AuthSwitch = &v
	return s
}

func (s *SetEventCallbackRequest) SetCallbackQueueName(v string) *SetEventCallbackRequest {
	s.CallbackQueueName = &v
	return s
}

func (s *SetEventCallbackRequest) SetCallbackType(v string) *SetEventCallbackRequest {
	s.CallbackType = &v
	return s
}

func (s *SetEventCallbackRequest) SetCallbackURL(v string) *SetEventCallbackRequest {
	s.CallbackURL = &v
	return s
}

func (s *SetEventCallbackRequest) SetEventTypeList(v string) *SetEventCallbackRequest {
	s.EventTypeList = &v
	return s
}

func (s *SetEventCallbackRequest) Validate() error {
	return dara.Validate(s)
}
