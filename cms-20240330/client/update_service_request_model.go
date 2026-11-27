// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v string) *UpdateServiceRequest
	GetAttributes() *string
	SetDescription(v string) *UpdateServiceRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateServiceRequest
	GetDisplayName() *string
	SetServiceStatus(v string) *UpdateServiceRequest
	GetServiceStatus() *string
	SetDryRun(v bool) *UpdateServiceRequest
	GetDryRun() *bool
}

type UpdateServiceRequest struct {
	// The extended attributes.
	//
	// example:
	//
	// {"language":"java"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The service description. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// api-monitor-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The service status. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// Stopped
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// Specifies whether to perform a dry run of the request. Valid values: true: sends a check request and does not perform the actual operation. false: sends a normal request. Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) GetAttributes() *string {
	return s.Attributes
}

func (s *UpdateServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateServiceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateServiceRequest) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *UpdateServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceRequest) SetAttributes(v string) *UpdateServiceRequest {
	s.Attributes = &v
	return s
}

func (s *UpdateServiceRequest) SetDescription(v string) *UpdateServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateServiceRequest) SetDisplayName(v string) *UpdateServiceRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceStatus(v string) *UpdateServiceRequest {
	s.ServiceStatus = &v
	return s
}

func (s *UpdateServiceRequest) SetDryRun(v bool) *UpdateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceRequest) Validate() error {
	return dara.Validate(s)
}
