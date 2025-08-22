// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModuleVersion(v *GetRegistryModuleVersionResponseBodyModuleVersion) *GetRegistryModuleVersionResponseBody
	GetModuleVersion() *GetRegistryModuleVersionResponseBodyModuleVersion
	SetRequestId(v string) *GetRegistryModuleVersionResponseBody
	GetRequestId() *string
}

type GetRegistryModuleVersionResponseBody struct {
	ModuleVersion *GetRegistryModuleVersionResponseBodyModuleVersion `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty" type:"Struct"`
	// example:
	//
	// 62DF26B0-53F0-5747-9D7F-FEF444FB4E24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRegistryModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleVersionResponseBody) GetModuleVersion() *GetRegistryModuleVersionResponseBodyModuleVersion {
	return s.ModuleVersion
}

func (s *GetRegistryModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegistryModuleVersionResponseBody) SetModuleVersion(v *GetRegistryModuleVersionResponseBodyModuleVersion) *GetRegistryModuleVersionResponseBody {
	s.ModuleVersion = v
	return s
}

func (s *GetRegistryModuleVersionResponseBody) SetRequestId(v string) *GetRegistryModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRegistryModuleVersionResponseBodyModuleVersion struct {
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DetailUrl     *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	Downloads     *string `json:"downloads,omitempty" xml:"downloads,omitempty"`
	ModuleName    *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	Provider      *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Source        *string `json:"source,omitempty" xml:"source,omitempty"`
	SourceUrl     *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetRegistryModuleVersionResponseBodyModuleVersion) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleVersionResponseBodyModuleVersion) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetDownloads() *string {
	return s.Downloads
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetProvider() *string {
	return s.Provider
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetSource() *string {
	return s.Source
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) GetVersion() *string {
	return s.Version
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetCreateTime(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.CreateTime = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetDetailUrl(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.DetailUrl = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetDownloads(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.Downloads = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetModuleName(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.ModuleName = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetNamespaceName(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.NamespaceName = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetProvider(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.Provider = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetSource(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.Source = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetSourceUrl(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.SourceUrl = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) SetVersion(v string) *GetRegistryModuleVersionResponseBodyModuleVersion {
	s.Version = &v
	return s
}

func (s *GetRegistryModuleVersionResponseBodyModuleVersion) Validate() error {
	return dara.Validate(s)
}
