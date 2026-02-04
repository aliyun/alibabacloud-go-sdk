// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceControlConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceControlConfiguration(v *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) *GetInstanceControlConfigurationResponseBody
	GetInstanceControlConfiguration() *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration
	SetRequestId(v string) *GetInstanceControlConfigurationResponseBody
	GetRequestId() *string
}

type GetInstanceControlConfigurationResponseBody struct {
	InstanceControlConfiguration *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration `json:"InstanceControlConfiguration,omitempty" xml:"InstanceControlConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceControlConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationResponseBody) GetInstanceControlConfiguration() *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration {
	return s.InstanceControlConfiguration
}

func (s *GetInstanceControlConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceControlConfigurationResponseBody) SetInstanceControlConfiguration(v *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) *GetInstanceControlConfigurationResponseBody {
	s.InstanceControlConfiguration = v
	return s
}

func (s *GetInstanceControlConfigurationResponseBody) SetRequestId(v string) *GetInstanceControlConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceControlConfigurationResponseBody) Validate() error {
	if s.InstanceControlConfiguration != nil {
		if err := s.InstanceControlConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration struct {
	// 实例控制配置项
	ControlElements []*GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements `json:"ControlElements,omitempty" xml:"ControlElements,omitempty" type:"Repeated"`
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) GetControlElements() []*GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements {
	return s.ControlElements
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) SetControlElements(v []*GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration {
	s.ControlElements = v
	return s
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfiguration) Validate() error {
	if s.ControlElements != nil {
		for _, item := range s.ControlElements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements struct {
	// 实例控制项名称，如human_verification。
	//
	// example:
	//
	// human_verification
	ElementName             *string                                                                                                        `json:"ElementName,omitempty" xml:"ElementName,omitempty"`
	HumanVerificationConfig *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig `json:"HumanVerificationConfig,omitempty" xml:"HumanVerificationConfig,omitempty" type:"Struct"`
	// 实例控制项状态，enabled或者disabled。
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) GetElementName() *string {
	return s.ElementName
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) GetHumanVerificationConfig() *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig {
	return s.HumanVerificationConfig
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) SetElementName(v string) *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements {
	s.ElementName = &v
	return s
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) SetHumanVerificationConfig(v *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements {
	s.HumanVerificationConfig = v
	return s
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) SetStatus(v string) *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements {
	s.Status = &v
	return s
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElements) Validate() error {
	if s.HumanVerificationConfig != nil {
		if err := s.HumanVerificationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig struct {
	// example:
	//
	// urn:alibaba:idaas:humanverification:alibaba-cloud-jigsaw-verification
	HumanVerificationType *string `json:"HumanVerificationType,omitempty" xml:"HumanVerificationType,omitempty"`
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) GetHumanVerificationType() *string {
	return s.HumanVerificationType
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) SetHumanVerificationType(v string) *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig {
	s.HumanVerificationType = &v
	return s
}

func (s *GetInstanceControlConfigurationResponseBodyInstanceControlConfigurationControlElementsHumanVerificationConfig) Validate() error {
	return dara.Validate(s)
}
