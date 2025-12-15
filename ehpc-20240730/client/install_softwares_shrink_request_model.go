// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSoftwaresShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackagesShrink(v string) *InstallSoftwaresShrinkRequest
	GetAdditionalPackagesShrink() *string
	SetClusterId(v string) *InstallSoftwaresShrinkRequest
	GetClusterId() *string
}

type InstallSoftwaresShrinkRequest struct {
	// The information about the software systems that you want to install.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwaresShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallSoftwaresShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresShrinkRequest) GetAdditionalPackagesShrink() *string {
	return s.AdditionalPackagesShrink
}

func (s *InstallSoftwaresShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallSoftwaresShrinkRequest) SetAdditionalPackagesShrink(v string) *InstallSoftwaresShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *InstallSoftwaresShrinkRequest) SetClusterId(v string) *InstallSoftwaresShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallSoftwaresShrinkRequest) Validate() error {
	return dara.Validate(s)
}
