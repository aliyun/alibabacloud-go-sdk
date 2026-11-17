// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultServiceId(v string) *UpdateAiModelProviderRequest
	GetDefaultServiceId() *string
	SetDisplayName(v string) *UpdateAiModelProviderRequest
	GetDisplayName() *string
	SetServiceIds(v []*string) *UpdateAiModelProviderRequest
	GetServiceIds() []*string
}

type UpdateAiModelProviderRequest struct {
	// The optional default AI service ID. If not empty, it must be a service in the updated association set. Omitting this parameter retains the saved default. An explicit empty string clears the default. A non-empty value sets the default. The value can be up to 64 characters in length.
	//
	// example:
	//
	// svc-d00000000000000000000
	DefaultServiceId *string `json:"defaultServiceId,omitempty" xml:"defaultServiceId,omitempty"`
	// The display name of the model provider. This parameter is required and can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen Cloud / Alibaba Cloud Model Studio
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The complete set of AI service IDs to associate with the model provider. If you do not specify this parameter, the existing associations are retained. If you pass an empty array, all associations are cleared. Duplicate and empty values are normalized.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s UpdateAiModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderRequest) GetDefaultServiceId() *string {
	return s.DefaultServiceId
}

func (s *UpdateAiModelProviderRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAiModelProviderRequest) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *UpdateAiModelProviderRequest) SetDefaultServiceId(v string) *UpdateAiModelProviderRequest {
	s.DefaultServiceId = &v
	return s
}

func (s *UpdateAiModelProviderRequest) SetDisplayName(v string) *UpdateAiModelProviderRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAiModelProviderRequest) SetServiceIds(v []*string) *UpdateAiModelProviderRequest {
	s.ServiceIds = v
	return s
}

func (s *UpdateAiModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
