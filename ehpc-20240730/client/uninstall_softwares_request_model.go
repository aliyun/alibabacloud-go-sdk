// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSoftwaresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackages(v []*UninstallSoftwaresRequestAdditionalPackages) *UninstallSoftwaresRequest
	GetAdditionalPackages() []*UninstallSoftwaresRequestAdditionalPackages
	SetClusterId(v string) *UninstallSoftwaresRequest
	GetClusterId() *string
}

type UninstallSoftwaresRequest struct {
	// The information about the software systems that you want to uninstall.
	AdditionalPackages []*UninstallSoftwaresRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwaresRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresRequest) GetAdditionalPackages() []*UninstallSoftwaresRequestAdditionalPackages {
	return s.AdditionalPackages
}

func (s *UninstallSoftwaresRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallSoftwaresRequest) SetAdditionalPackages(v []*UninstallSoftwaresRequestAdditionalPackages) *UninstallSoftwaresRequest {
	s.AdditionalPackages = v
	return s
}

func (s *UninstallSoftwaresRequest) SetClusterId(v string) *UninstallSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallSoftwaresRequest) Validate() error {
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

type UninstallSoftwaresRequestAdditionalPackages struct {
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

func (s UninstallSoftwaresRequestAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s UninstallSoftwaresRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresRequestAdditionalPackages) GetName() *string {
	return s.Name
}

func (s *UninstallSoftwaresRequestAdditionalPackages) GetVersion() *string {
	return s.Version
}

func (s *UninstallSoftwaresRequestAdditionalPackages) SetName(v string) *UninstallSoftwaresRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *UninstallSoftwaresRequestAdditionalPackages) SetVersion(v string) *UninstallSoftwaresRequestAdditionalPackages {
	s.Version = &v
	return s
}

func (s *UninstallSoftwaresRequestAdditionalPackages) Validate() error {
	return dara.Validate(s)
}
