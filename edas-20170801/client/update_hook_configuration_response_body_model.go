// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHookConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateHookConfigurationResponseBody
	GetCode() *int32
	SetHooksConfiguration(v []*UpdateHookConfigurationResponseBodyHooksConfiguration) *UpdateHookConfigurationResponseBody
	GetHooksConfiguration() []*UpdateHookConfigurationResponseBodyHooksConfiguration
	SetMessage(v string) *UpdateHookConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHookConfigurationResponseBody
	GetRequestId() *string
}

type UpdateHookConfigurationResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the mounted script.
	HooksConfiguration []*UpdateHookConfigurationResponseBodyHooksConfiguration `json:"HooksConfiguration,omitempty" xml:"HooksConfiguration,omitempty" type:"Repeated"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// d498****-1dd8ec229862
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHookConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHookConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHookConfigurationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateHookConfigurationResponseBody) GetHooksConfiguration() []*UpdateHookConfigurationResponseBodyHooksConfiguration {
	return s.HooksConfiguration
}

func (s *UpdateHookConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHookConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHookConfigurationResponseBody) SetCode(v int32) *UpdateHookConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHookConfigurationResponseBody) SetHooksConfiguration(v []*UpdateHookConfigurationResponseBodyHooksConfiguration) *UpdateHookConfigurationResponseBody {
	s.HooksConfiguration = v
	return s
}

func (s *UpdateHookConfigurationResponseBody) SetMessage(v string) *UpdateHookConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHookConfigurationResponseBody) SetRequestId(v string) *UpdateHookConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHookConfigurationResponseBody) Validate() error {
	if s.HooksConfiguration != nil {
		for _, item := range s.HooksConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHookConfigurationResponseBodyHooksConfiguration struct {
	// Indicates whether a mount failure is ignored. Valid values:
	//
	// 	- **true**: A mount failure is ignored.
	//
	// 	- **false**: A mount failure is not ignored.
	//
	// example:
	//
	// true
	IgnoreFail *bool `json:"IgnoreFail,omitempty" xml:"IgnoreFail,omitempty"`
	// The name of the mounted script.
	//
	// example:
	//
	// postprepareInstanceEnvironmentOnScaleOut
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the mounted script.
	//
	// example:
	//
	// ls
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s UpdateHookConfigurationResponseBodyHooksConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateHookConfigurationResponseBodyHooksConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) GetIgnoreFail() *bool {
	return s.IgnoreFail
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) GetName() *string {
	return s.Name
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) GetScript() *string {
	return s.Script
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) SetIgnoreFail(v bool) *UpdateHookConfigurationResponseBodyHooksConfiguration {
	s.IgnoreFail = &v
	return s
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) SetName(v string) *UpdateHookConfigurationResponseBodyHooksConfiguration {
	s.Name = &v
	return s
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) SetScript(v string) *UpdateHookConfigurationResponseBodyHooksConfiguration {
	s.Script = &v
	return s
}

func (s *UpdateHookConfigurationResponseBodyHooksConfiguration) Validate() error {
	return dara.Validate(s)
}
