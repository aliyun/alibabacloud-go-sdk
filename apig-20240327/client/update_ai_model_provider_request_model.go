// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *UpdateAiModelProviderRequest
	GetDisplayName() *string
	SetServiceIds(v []*string) *UpdateAiModelProviderRequest
	GetServiceIds() []*string
}

type UpdateAiModelProviderRequest struct {
	// The display name of the model provider. This parameter is required. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen Cloud / Alibaba Cloud Bailian.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The complete set of AI service IDs to bind to the model provider. If this parameter is not specified, the existing bindings are retained. If an empty array is specified, all bindings are cleared. Duplicate and empty values are normalized.
	ServiceIds []*string `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s UpdateAiModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAiModelProviderRequest) GetServiceIds() []*string {
	return s.ServiceIds
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
