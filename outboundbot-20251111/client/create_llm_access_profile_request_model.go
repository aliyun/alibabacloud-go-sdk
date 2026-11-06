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
	// The instance ID.
	//
	// example:
	//
	// 33014787-cc13-49d3-ab2f-a98aa8f15fbb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model access configuration.
	Profile *CreateLlmAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
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
	// The API key.
	//
	// example:
	//
	// example-api-key
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The model service registration URL.
	//
	// example:
	//
	// https://adbpg-ai-gateway10-cn-beijing.opentrust.net/1383274970238153-193
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// Sample configuration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
