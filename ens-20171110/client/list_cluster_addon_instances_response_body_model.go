// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListClusterAddonInstancesResponseBodyAddons) *ListClusterAddonInstancesResponseBody
	GetAddons() []*ListClusterAddonInstancesResponseBodyAddons
	SetRequestId(v string) *ListClusterAddonInstancesResponseBody
	GetRequestId() *string
}

type ListClusterAddonInstancesResponseBody struct {
	Addons []*ListClusterAddonInstancesResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterAddonInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponseBody) GetAddons() []*ListClusterAddonInstancesResponseBodyAddons {
	return s.Addons
}

func (s *ListClusterAddonInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterAddonInstancesResponseBody) SetAddons(v []*ListClusterAddonInstancesResponseBodyAddons) *ListClusterAddonInstancesResponseBody {
	s.Addons = v
	return s
}

func (s *ListClusterAddonInstancesResponseBody) SetRequestId(v string) *ListClusterAddonInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBody) Validate() error {
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

type ListClusterAddonInstancesResponseBodyAddons struct {
	// example:
	//
	// true
	CleanupCloudResources *bool                                                      `json:"CleanupCloudResources,omitempty" xml:"CleanupCloudResources,omitempty"`
	ConfigSchema          []*ListClusterAddonInstancesResponseBodyAddonsConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClusterAddonInstancesResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetCleanupCloudResources() *bool {
	return s.CleanupCloudResources
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetConfigSchema() []*ListClusterAddonInstancesResponseBodyAddonsConfigSchema {
	return s.ConfigSchema
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetStatus() *string {
	return s.Status
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetCleanupCloudResources(v bool) *ListClusterAddonInstancesResponseBodyAddons {
	s.CleanupCloudResources = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetConfigSchema(v []*ListClusterAddonInstancesResponseBodyAddonsConfigSchema) *ListClusterAddonInstancesResponseBodyAddons {
	s.ConfigSchema = v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetName(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetStatus(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.Status = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetVersion(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) Validate() error {
	if s.ConfigSchema != nil {
		for _, item := range s.ConfigSchema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterAddonInstancesResponseBodyAddonsConfigSchema struct {
	// example:
	//
	// 7380581386597434629002
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 4572581386436834662215
	ConfigVersion *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"key1": "val1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ListClusterAddonInstancesResponseBodyAddonsConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponseBodyAddonsConfigSchema) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) GetName() *string {
	return s.Name
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) GetParams() *string {
	return s.Params
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) SetAppVersion(v string) *ListClusterAddonInstancesResponseBodyAddonsConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) SetConfigVersion(v string) *ListClusterAddonInstancesResponseBodyAddonsConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) SetName(v string) *ListClusterAddonInstancesResponseBodyAddonsConfigSchema {
	s.Name = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) SetParams(v string) *ListClusterAddonInstancesResponseBodyAddonsConfigSchema {
	s.Params = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddonsConfigSchema) Validate() error {
	return dara.Validate(s)
}
