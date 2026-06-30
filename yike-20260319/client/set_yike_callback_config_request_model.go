// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeCallbackConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackConfig(v string) *SetYikeCallbackConfigRequest
	GetCallbackConfig() *string
	SetCallbackUrl(v string) *SetYikeCallbackConfigRequest
	GetCallbackUrl() *string
}

type SetYikeCallbackConfigRequest struct {
	// The callback configuration. CallbackEventList supports multiple events, including:
	//
	// - EventType: the event type. Valid values:
	//
	//   - UserCreated: The user is created.
	//
	//   - UserCreditAdded: Credits are added to the user.
	//
	// - UserData: a JSON object that contains user-defined data.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"CallbackEventList":[{"EventType":"UserCreditAdded","UserData":"{}"}]}
	CallbackConfig *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// The HTTP callback URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// http//example.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
}

func (s SetYikeCallbackConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetYikeCallbackConfigRequest) GoString() string {
	return s.String()
}

func (s *SetYikeCallbackConfigRequest) GetCallbackConfig() *string {
	return s.CallbackConfig
}

func (s *SetYikeCallbackConfigRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SetYikeCallbackConfigRequest) SetCallbackConfig(v string) *SetYikeCallbackConfigRequest {
	s.CallbackConfig = &v
	return s
}

func (s *SetYikeCallbackConfigRequest) SetCallbackUrl(v string) *SetYikeCallbackConfigRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SetYikeCallbackConfigRequest) Validate() error {
	return dara.Validate(s)
}
