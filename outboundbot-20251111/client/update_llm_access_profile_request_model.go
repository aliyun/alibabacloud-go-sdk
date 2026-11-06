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
	// The access profile ID.
	//
	// example:
	//
	// ed06946e-dcbb-4f3a-baf8-a5feae0198e9
	AccessProfileId *string `json:"AccessProfileId,omitempty" xml:"AccessProfileId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model access configuration.
	Profile *UpdateLlmAccessProfileRequestProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
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
	// wss://example.com/api/realtime
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// Sample configuration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
