// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigurationRecorderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceTypes(v string) *UpdateConfigurationRecorderRequest
	GetResourceTypes() *string
}

type UpdateConfigurationRecorderRequest struct {
	// The resource types. Separate multiple resource types with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s UpdateConfigurationRecorderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigurationRecorderRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationRecorderRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *UpdateConfigurationRecorderRequest) SetResourceTypes(v string) *UpdateConfigurationRecorderRequest {
	s.ResourceTypes = &v
	return s
}

func (s *UpdateConfigurationRecorderRequest) Validate() error {
	return dara.Validate(s)
}
