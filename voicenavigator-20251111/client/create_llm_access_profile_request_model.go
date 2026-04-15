// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLlmAccessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateLlmAccessProfileRequest
	GetInstanceId() *string
	SetProfile(v *CreateLlmAccessProfileRequestProfile) *CreateLlmAccessProfileRequest
	GetProfile() *CreateLlmAccessProfileRequestProfile
}

type CreateLlmAccessProfileRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Profile    *CreateLlmAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
}

func (s CreateLlmAccessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLlmAccessProfileRequest) GetProfile() *CreateLlmAccessProfileRequestProfile {
	return s.Profile
}

func (s *CreateLlmAccessProfileRequest) SetInstanceId(v string) *CreateLlmAccessProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLlmAccessProfileRequest) SetProfile(v *CreateLlmAccessProfileRequestProfile) *CreateLlmAccessProfileRequest {
	s.Profile = v
	return s
}

func (s *CreateLlmAccessProfileRequest) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLlmAccessProfileRequestProfile struct {
	// example:
	//
	// 70tKleNtMGaaIem7us7Miw-Tf3kPzE6l
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api.llm.enpoint.example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateLlmAccessProfileRequestProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileRequestProfile) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileRequestProfile) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateLlmAccessProfileRequestProfile) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateLlmAccessProfileRequestProfile) GetName() *string {
	return s.Name
}

func (s *CreateLlmAccessProfileRequestProfile) SetApiKey(v string) *CreateLlmAccessProfileRequestProfile {
	s.ApiKey = &v
	return s
}

func (s *CreateLlmAccessProfileRequestProfile) SetEndpoint(v string) *CreateLlmAccessProfileRequestProfile {
	s.Endpoint = &v
	return s
}

func (s *CreateLlmAccessProfileRequestProfile) SetName(v string) *CreateLlmAccessProfileRequestProfile {
	s.Name = &v
	return s
}

func (s *CreateLlmAccessProfileRequestProfile) Validate() error {
	return dara.Validate(s)
}
