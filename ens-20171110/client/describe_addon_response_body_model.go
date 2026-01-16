// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*DescribeAddonResponseBodyAddons) *DescribeAddonResponseBody
	GetAddons() []*DescribeAddonResponseBodyAddons
	SetRequestId(v string) *DescribeAddonResponseBody
	GetRequestId() *string
}

type DescribeAddonResponseBody struct {
	// if can be null:
	// false
	Addons []*DescribeAddonResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponseBody) GetAddons() []*DescribeAddonResponseBodyAddons {
	return s.Addons
}

func (s *DescribeAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddonResponseBody) SetAddons(v []*DescribeAddonResponseBodyAddons) *DescribeAddonResponseBody {
	s.Addons = v
	return s
}

func (s *DescribeAddonResponseBody) SetRequestId(v string) *DescribeAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddonResponseBody) Validate() error {
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

type DescribeAddonResponseBodyAddons struct {
	// example:
	//
	// true
	CleanupCloudResources *string                                        `json:"CleanupCloudResources,omitempty" xml:"CleanupCloudResources,omitempty"`
	ConfigSchema          []*DescribeAddonResponseBodyAddonsConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAddonResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponseBodyAddons) GetCleanupCloudResources() *string {
	return s.CleanupCloudResources
}

func (s *DescribeAddonResponseBodyAddons) GetConfigSchema() []*DescribeAddonResponseBodyAddonsConfigSchema {
	return s.ConfigSchema
}

func (s *DescribeAddonResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *DescribeAddonResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonResponseBodyAddons) SetCleanupCloudResources(v string) *DescribeAddonResponseBodyAddons {
	s.CleanupCloudResources = &v
	return s
}

func (s *DescribeAddonResponseBodyAddons) SetConfigSchema(v []*DescribeAddonResponseBodyAddonsConfigSchema) *DescribeAddonResponseBodyAddons {
	s.ConfigSchema = v
	return s
}

func (s *DescribeAddonResponseBodyAddons) SetName(v string) *DescribeAddonResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *DescribeAddonResponseBodyAddons) SetVersion(v string) *DescribeAddonResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *DescribeAddonResponseBodyAddons) Validate() error {
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

type DescribeAddonResponseBodyAddonsConfigSchema struct {
	// example:
	//
	// 859e9d595b2974ed79c444658d1dea89
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 4155709cd12a09bdb8cbaca71bf03233
	ConfigVersion *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"key1":"val1"}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DescribeAddonResponseBodyAddonsConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponseBodyAddonsConfigSchema) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) GetName() *string {
	return s.Name
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) GetParams() interface{} {
	return s.Params
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) SetAppVersion(v string) *DescribeAddonResponseBodyAddonsConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) SetConfigVersion(v string) *DescribeAddonResponseBodyAddonsConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) SetName(v string) *DescribeAddonResponseBodyAddonsConfigSchema {
	s.Name = &v
	return s
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) SetParams(v interface{}) *DescribeAddonResponseBodyAddonsConfigSchema {
	s.Params = v
	return s
}

func (s *DescribeAddonResponseBodyAddonsConfigSchema) Validate() error {
	return dara.Validate(s)
}
