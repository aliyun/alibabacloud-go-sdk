// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAddonInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCleanupCloudResources(v bool) *GetClusterAddonInstanceResponseBody
	GetCleanupCloudResources() *bool
	SetConfigSchema(v []*GetClusterAddonInstanceResponseBodyConfigSchema) *GetClusterAddonInstanceResponseBody
	GetConfigSchema() []*GetClusterAddonInstanceResponseBodyConfigSchema
	SetName(v string) *GetClusterAddonInstanceResponseBody
	GetName() *string
	SetRequestId(v string) *GetClusterAddonInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetClusterAddonInstanceResponseBody
	GetStatus() *string
	SetVersion(v string) *GetClusterAddonInstanceResponseBody
	GetVersion() *string
}

type GetClusterAddonInstanceResponseBody struct {
	CleanupCloudResources *bool                                              `json:"CleanupCloudResources,omitempty" xml:"CleanupCloudResources,omitempty"`
	ConfigSchema          []*GetClusterAddonInstanceResponseBodyConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterAddonInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceResponseBody) GetCleanupCloudResources() *bool {
	return s.CleanupCloudResources
}

func (s *GetClusterAddonInstanceResponseBody) GetConfigSchema() []*GetClusterAddonInstanceResponseBodyConfigSchema {
	return s.ConfigSchema
}

func (s *GetClusterAddonInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetClusterAddonInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterAddonInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetClusterAddonInstanceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetClusterAddonInstanceResponseBody) SetCleanupCloudResources(v bool) *GetClusterAddonInstanceResponseBody {
	s.CleanupCloudResources = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetConfigSchema(v []*GetClusterAddonInstanceResponseBodyConfigSchema) *GetClusterAddonInstanceResponseBody {
	s.ConfigSchema = v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetName(v string) *GetClusterAddonInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetRequestId(v string) *GetClusterAddonInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetStatus(v string) *GetClusterAddonInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) SetVersion(v string) *GetClusterAddonInstanceResponseBody {
	s.Version = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBody) Validate() error {
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

type GetClusterAddonInstanceResponseBodyConfigSchema struct {
	// example:
	//
	// d0ead1f4e28de0f9e3c86588409a88a4
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// d0ead1f4e28de0f9e3c86588409a88a4
	ConfigVersion *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"k1":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetClusterAddonInstanceResponseBodyConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceResponseBodyConfigSchema) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) GetName() *string {
	return s.Name
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) GetParams() *string {
	return s.Params
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) SetAppVersion(v string) *GetClusterAddonInstanceResponseBodyConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) SetConfigVersion(v string) *GetClusterAddonInstanceResponseBodyConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) SetName(v string) *GetClusterAddonInstanceResponseBodyConfigSchema {
	s.Name = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) SetParams(v string) *GetClusterAddonInstanceResponseBodyConfigSchema {
	s.Params = &v
	return s
}

func (s *GetClusterAddonInstanceResponseBodyConfigSchema) Validate() error {
	return dara.Validate(s)
}
