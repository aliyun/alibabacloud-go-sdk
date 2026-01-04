// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody
	GetAddons() []*ListAddonsResponseBodyAddons
	SetRequestId(v string) *ListAddonsResponseBody
	GetRequestId() *string
}

type ListAddonsResponseBody struct {
	Addons []*ListAddonsResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) GetAddons() []*ListAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonsResponseBody) SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonsResponseBody) SetRequestId(v string) *ListAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonsResponseBody) Validate() error {
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

type ListAddonsResponseBodyAddons struct {
	// example:
	//
	// true
	CleanupCloudResources *bool                                       `json:"CleanupCloudResources,omitempty" xml:"CleanupCloudResources,omitempty"`
	ConfigSchema          []*ListAddonsResponseBodyAddonsConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddons) GetCleanupCloudResources() *bool {
	return s.CleanupCloudResources
}

func (s *ListAddonsResponseBodyAddons) GetConfigSchema() []*ListAddonsResponseBodyAddonsConfigSchema {
	return s.ConfigSchema
}

func (s *ListAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyAddons) SetCleanupCloudResources(v bool) *ListAddonsResponseBodyAddons {
	s.CleanupCloudResources = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetConfigSchema(v []*ListAddonsResponseBodyAddonsConfigSchema) *ListAddonsResponseBodyAddons {
	s.ConfigSchema = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetName(v string) *ListAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetVersion(v string) *ListAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) Validate() error {
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

type ListAddonsResponseBodyAddonsConfigSchema struct {
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
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ListAddonsResponseBodyAddonsConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddonsConfigSchema) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) GetParams() map[string]interface{} {
	return s.Params
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) SetAppVersion(v string) *ListAddonsResponseBodyAddonsConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) SetConfigVersion(v string) *ListAddonsResponseBodyAddonsConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) SetName(v string) *ListAddonsResponseBodyAddonsConfigSchema {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) SetParams(v map[string]interface{}) *ListAddonsResponseBodyAddonsConfigSchema {
	s.Params = v
	return s
}

func (s *ListAddonsResponseBodyAddonsConfigSchema) Validate() error {
	return dara.Validate(s)
}
