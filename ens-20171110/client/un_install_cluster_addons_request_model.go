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
	SetClusterId(v string) *UnInstallClusterAddonsRequest
	GetClusterId() *string
}

type UnInstallClusterAddonsRequest struct {
	Addons []*UnInstallClusterAddonsRequestAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s *UnInstallClusterAddonsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnInstallClusterAddonsRequest) SetAddons(v []*UnInstallClusterAddonsRequestAddons) *UnInstallClusterAddonsRequest {
	s.Addons = v
	return s
}

func (s *UnInstallClusterAddonsRequest) SetClusterId(v string) *UnInstallClusterAddonsRequest {
	s.ClusterId = &v
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
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UnInstallClusterAddonsRequestAddons) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsRequestAddons) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequestAddons) GetName() *string {
	return s.Name
}

func (s *UnInstallClusterAddonsRequestAddons) SetName(v string) *UnInstallClusterAddonsRequestAddons {
	s.Name = &v
	return s
}

func (s *UnInstallClusterAddonsRequestAddons) Validate() error {
	return dara.Validate(s)
}
