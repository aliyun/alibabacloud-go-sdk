// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSoftwaresShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackagesShrink(v string) *UninstallSoftwaresShrinkRequest
	GetAdditionalPackagesShrink() *string
	SetClusterId(v string) *UninstallSoftwaresShrinkRequest
	GetClusterId() *string
}

type UninstallSoftwaresShrinkRequest struct {
	// The information about the software systems that you want to uninstall.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwaresShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallSoftwaresShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresShrinkRequest) GetAdditionalPackagesShrink() *string {
	return s.AdditionalPackagesShrink
}

func (s *UninstallSoftwaresShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallSoftwaresShrinkRequest) SetAdditionalPackagesShrink(v string) *UninstallSoftwaresShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *UninstallSoftwaresShrinkRequest) SetClusterId(v string) *UninstallSoftwaresShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallSoftwaresShrinkRequest) Validate() error {
	return dara.Validate(s)
}
