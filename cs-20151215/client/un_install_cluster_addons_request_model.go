// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallClusterAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*UnInstallClusterAddonsRequestAddons) *UnInstallClusterAddonsRequest
	GetAddons() []*UnInstallClusterAddonsRequestAddons
}

type UnInstallClusterAddonsRequest struct {
	// The list of add-ons to uninstall.
	Addons []*UnInstallClusterAddonsRequestAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
}

func (s UnInstallClusterAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequest) GetAddons() []*UnInstallClusterAddonsRequestAddons {
	return s.Addons
}

func (s *UnInstallClusterAddonsRequest) SetAddons(v []*UnInstallClusterAddonsRequestAddons) *UnInstallClusterAddonsRequest {
	s.Addons = v
	return s
}

func (s *UnInstallClusterAddonsRequest) Validate() error {
	return dara.Validate(s)
}

type UnInstallClusterAddonsRequestAddons struct {
	// Specifies whether to clean up related cloud resources during uninstallation.
	//
	// 	- true: clean up
	//
	// 	- false: retain
	//
	// example:
	//
	// true
	CleanupCloudResources *bool `json:"cleanup_cloud_resources,omitempty" xml:"cleanup_cloud_resources,omitempty"`
	// The name of the add-on to uninstall. You can call the [ListClusterAddonInstances](https://help.aliyun.com/document_detail/2667940.html) operation to query the installed add-ons.
	//
	// example:
	//
	// ack-node-problem-detector
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UnInstallClusterAddonsRequestAddons) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsRequestAddons) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequestAddons) GetCleanupCloudResources() *bool {
	return s.CleanupCloudResources
}

func (s *UnInstallClusterAddonsRequestAddons) GetName() *string {
	return s.Name
}

func (s *UnInstallClusterAddonsRequestAddons) SetCleanupCloudResources(v bool) *UnInstallClusterAddonsRequestAddons {
	s.CleanupCloudResources = &v
	return s
}

func (s *UnInstallClusterAddonsRequestAddons) SetName(v string) *UnInstallClusterAddonsRequestAddons {
	s.Name = &v
	return s
}

func (s *UnInstallClusterAddonsRequestAddons) Validate() error {
	return dara.Validate(s)
}
