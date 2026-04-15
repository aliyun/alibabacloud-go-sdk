// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLlmAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessProfileId(v string) *UpdateLlmAccessProfileRequest
	GetAccessProfileId() *string
	SetInstanceId(v string) *UpdateLlmAccessProfileRequest
	GetInstanceId() *string
	SetProfile(v *UpdateLlmAccessProfileRequestProfile) *UpdateLlmAccessProfileRequest
	GetProfile() *UpdateLlmAccessProfileRequestProfile
}

type UpdateLlmAccessProfileRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4060
	InstanceId *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Profile    *UpdateLlmAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
}

func (s UpdateLlmAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileRequest) GetAccessProfileId() *string {
	return s.AccessProfileId
}

func (s *UpdateLlmAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLlmAccessProfileRequest) GetProfile() *UpdateLlmAccessProfileRequestProfile {
	return s.Profile
}

func (s *UpdateLlmAccessProfileRequest) SetAccessProfileId(v string) *UpdateLlmAccessProfileRequest {
	s.AccessProfileId = &v
	return s
}

func (s *UpdateLlmAccessProfileRequest) SetInstanceId(v string) *UpdateLlmAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLlmAccessProfileRequest) SetProfile(v *UpdateLlmAccessProfileRequestProfile) *UpdateLlmAccessProfileRequest {
	s.Profile = v
	return s
}

func (s *UpdateLlmAccessProfileRequest) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLlmAccessProfileRequestProfile struct {
	// example:
	//
	// akm-091e4c2c-9938-4a0d-ade2-b9c477fbcbdb
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api.llm.enpoint.example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLlmAccessProfileRequestProfile) String() string {
	return dara.Prettify(s)
}

func (s UpdateLlmAccessProfileRequestProfile) GoString() string {
	return s.String()
}

func (s *UpdateLlmAccessProfileRequestProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateLlmAccessProfileRequestProfile) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateLlmAccessProfileRequestProfile) GetName() *string {
	return s.Name
}

func (s *UpdateLlmAccessProfileRequestProfile) SetApiKey(v string) *UpdateLlmAccessProfileRequestProfile {
	s.ApiKey = &v
	return s
}

func (s *UpdateLlmAccessProfileRequestProfile) SetEndpoint(v string) *UpdateLlmAccessProfileRequestProfile {
	s.Endpoint = &v
	return s
}

func (s *UpdateLlmAccessProfileRequestProfile) SetName(v string) *UpdateLlmAccessProfileRequestProfile {
	s.Name = &v
	return s
}

func (s *UpdateLlmAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
