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
	// The list of add-ons.
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
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UnInstallClusterAddonsRequestAddons struct {
	// Specifies whether to delete the associated cloud resources when uninstalling the add-on. Valid values:
	//
	// - true: Deletes the resources.
	//
	// - false: Retains the resources.
	//
	// example:
	//
	// true
	CleanupCloudResources *bool `json:"cleanup_cloud_resources,omitempty" xml:"cleanup_cloud_resources,omitempty"`
	// The name of the add-on to uninstall. You can call the [ListClusterAddonInstances](https://help.aliyun.com/document_detail/2667940.html) operation to query the add-ons installed in the cluster.
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
