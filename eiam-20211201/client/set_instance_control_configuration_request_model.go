// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceControlConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlElements(v []*SetInstanceControlConfigurationRequestControlElements) *SetInstanceControlConfigurationRequest
	GetControlElements() []*SetInstanceControlConfigurationRequestControlElements
	SetInstanceId(v string) *SetInstanceControlConfigurationRequest
	GetInstanceId() *string
}

type SetInstanceControlConfigurationRequest struct {
	// The instance control elements.
	ControlElements []*SetInstanceControlConfigurationRequestControlElements `json:"ControlElements,omitempty" xml:"ControlElements,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetInstanceControlConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceControlConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceControlConfigurationRequest) GetControlElements() []*SetInstanceControlConfigurationRequestControlElements {
	return s.ControlElements
}

func (s *SetInstanceControlConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetInstanceControlConfigurationRequest) SetControlElements(v []*SetInstanceControlConfigurationRequestControlElements) *SetInstanceControlConfigurationRequest {
	s.ControlElements = v
	return s
}

func (s *SetInstanceControlConfigurationRequest) SetInstanceId(v string) *SetInstanceControlConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceControlConfigurationRequest) Validate() error {
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

type SetInstanceControlConfigurationRequestControlElements struct {
	// The name of the instance control element.
	//
	// example:
	//
	// human_verification
	ElementName *string `json:"ElementName,omitempty" xml:"ElementName,omitempty"`
	// The CAPTCHA authenticate configuration.
	HumanVerificationConfig *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig `json:"HumanVerificationConfig,omitempty" xml:"HumanVerificationConfig,omitempty" type:"Struct"`
	// The status of the instance control element.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetInstanceControlConfigurationRequestControlElements) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceControlConfigurationRequestControlElements) GoString() string {
	return s.String()
}

func (s *SetInstanceControlConfigurationRequestControlElements) GetElementName() *string {
	return s.ElementName
}

func (s *SetInstanceControlConfigurationRequestControlElements) GetHumanVerificationConfig() *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig {
	return s.HumanVerificationConfig
}

func (s *SetInstanceControlConfigurationRequestControlElements) GetStatus() *string {
	return s.Status
}

func (s *SetInstanceControlConfigurationRequestControlElements) SetElementName(v string) *SetInstanceControlConfigurationRequestControlElements {
	s.ElementName = &v
	return s
}

func (s *SetInstanceControlConfigurationRequestControlElements) SetHumanVerificationConfig(v *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) *SetInstanceControlConfigurationRequestControlElements {
	s.HumanVerificationConfig = v
	return s
}

func (s *SetInstanceControlConfigurationRequestControlElements) SetStatus(v string) *SetInstanceControlConfigurationRequestControlElements {
	s.Status = &v
	return s
}

func (s *SetInstanceControlConfigurationRequestControlElements) Validate() error {
	if s.HumanVerificationConfig != nil {
		if err := s.HumanVerificationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig struct {
	// The CAPTCHA type.
	//
	// example:
	//
	// urn:alibaba:idaas:humanverification:alibaba-cloud-slider-verification
	HumanVerificationType *string `json:"HumanVerificationType,omitempty" xml:"HumanVerificationType,omitempty"`
}

func (s SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) GoString() string {
	return s.String()
}

func (s *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) GetHumanVerificationType() *string {
	return s.HumanVerificationType
}

func (s *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) SetHumanVerificationType(v string) *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig {
	s.HumanVerificationType = &v
	return s
}

func (s *SetInstanceControlConfigurationRequestControlElementsHumanVerificationConfig) Validate() error {
	return dara.Validate(s)
}
