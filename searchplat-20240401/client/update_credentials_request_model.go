// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateCredentialsRequest
	GetEnabled() *bool
	SetDryRun(v bool) *UpdateCredentialsRequest
	GetDryRun() *bool
}

type UpdateCredentialsRequest struct {
	// Specifies whether the credential is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateCredentialsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCredentialsRequest) SetEnabled(v bool) *UpdateCredentialsRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateCredentialsRequest) SetDryRun(v bool) *UpdateCredentialsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
