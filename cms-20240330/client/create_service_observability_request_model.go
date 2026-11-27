// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceObservabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreateServiceObservabilityRequest
	GetDryRun() *bool
}

type CreateServiceObservabilityRequest struct {
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: The system checks the request but does not perform the actual operation.
	//
	// - false: The system processes the request normally.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateServiceObservabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceObservabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceObservabilityRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceObservabilityRequest) SetDryRun(v bool) *CreateServiceObservabilityRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceObservabilityRequest) Validate() error {
	return dara.Validate(s)
}
