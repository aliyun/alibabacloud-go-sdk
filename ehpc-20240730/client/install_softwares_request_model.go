// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSoftwaresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackages(v []*InstallSoftwaresRequestAdditionalPackages) *InstallSoftwaresRequest
	GetAdditionalPackages() []*InstallSoftwaresRequestAdditionalPackages
	SetClusterId(v string) *InstallSoftwaresRequest
	GetClusterId() *string
}

type InstallSoftwaresRequest struct {
	// The information about the software systems that you want to install.
	AdditionalPackages []*InstallSoftwaresRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwaresRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresRequest) GetAdditionalPackages() []*InstallSoftwaresRequestAdditionalPackages {
	return s.AdditionalPackages
}

func (s *InstallSoftwaresRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallSoftwaresRequest) SetAdditionalPackages(v []*InstallSoftwaresRequestAdditionalPackages) *InstallSoftwaresRequest {
	s.AdditionalPackages = v
	return s
}

func (s *InstallSoftwaresRequest) SetClusterId(v string) *InstallSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallSoftwaresRequest) Validate() error {
	if s.AdditionalPackages != nil {
		for _, item := range s.AdditionalPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallSoftwaresRequestAdditionalPackages struct {
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstallSoftwaresRequestAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s InstallSoftwaresRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresRequestAdditionalPackages) GetName() *string {
	return s.Name
}

func (s *InstallSoftwaresRequestAdditionalPackages) GetVersion() *string {
	return s.Version
}

func (s *InstallSoftwaresRequestAdditionalPackages) SetName(v string) *InstallSoftwaresRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *InstallSoftwaresRequestAdditionalPackages) SetVersion(v string) *InstallSoftwaresRequestAdditionalPackages {
	s.Version = &v
	return s
}

func (s *InstallSoftwaresRequestAdditionalPackages) Validate() error {
	return dara.Validate(s)
}
