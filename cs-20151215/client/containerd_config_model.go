// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerdConfig interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreImageDefinedVolume(v bool) *ContainerdConfig
	GetIgnoreImageDefinedVolume() *bool
	SetInsecureRegistries(v []*string) *ContainerdConfig
	GetInsecureRegistries() []*string
	SetLimitCore(v int64) *ContainerdConfig
	GetLimitCore() *int64
	SetLimitMemLock(v int64) *ContainerdConfig
	GetLimitMemLock() *int64
	SetLimitNoFile(v int64) *ContainerdConfig
	GetLimitNoFile() *int64
	SetMaxConcurrentDownloads(v int64) *ContainerdConfig
	GetMaxConcurrentDownloads() *int64
	SetRegistryMirrors(v []*string) *ContainerdConfig
	GetRegistryMirrors() []*string
}

type ContainerdConfig struct {
	IgnoreImageDefinedVolume *bool     `json:"ignoreImageDefinedVolume,omitempty" xml:"ignoreImageDefinedVolume,omitempty"`
	InsecureRegistries       []*string `json:"insecureRegistries,omitempty" xml:"insecureRegistries,omitempty" type:"Repeated"`
	LimitCore                *int64    `json:"limitCore,omitempty" xml:"limitCore,omitempty"`
	LimitMemLock             *int64    `json:"limitMemLock,omitempty" xml:"limitMemLock,omitempty"`
	LimitNoFile              *int64    `json:"limitNoFile,omitempty" xml:"limitNoFile,omitempty"`
	MaxConcurrentDownloads   *int64    `json:"maxConcurrentDownloads,omitempty" xml:"maxConcurrentDownloads,omitempty"`
	RegistryMirrors          []*string `json:"registryMirrors,omitempty" xml:"registryMirrors,omitempty" type:"Repeated"`
}

func (s ContainerdConfig) String() string {
	return dara.Prettify(s)
}

func (s ContainerdConfig) GoString() string {
	return s.String()
}

func (s *ContainerdConfig) GetIgnoreImageDefinedVolume() *bool {
	return s.IgnoreImageDefinedVolume
}

func (s *ContainerdConfig) GetInsecureRegistries() []*string {
	return s.InsecureRegistries
}

func (s *ContainerdConfig) GetLimitCore() *int64 {
	return s.LimitCore
}

func (s *ContainerdConfig) GetLimitMemLock() *int64 {
	return s.LimitMemLock
}

func (s *ContainerdConfig) GetLimitNoFile() *int64 {
	return s.LimitNoFile
}

func (s *ContainerdConfig) GetMaxConcurrentDownloads() *int64 {
	return s.MaxConcurrentDownloads
}

func (s *ContainerdConfig) GetRegistryMirrors() []*string {
	return s.RegistryMirrors
}

func (s *ContainerdConfig) SetIgnoreImageDefinedVolume(v bool) *ContainerdConfig {
	s.IgnoreImageDefinedVolume = &v
	return s
}

func (s *ContainerdConfig) SetInsecureRegistries(v []*string) *ContainerdConfig {
	s.InsecureRegistries = v
	return s
}

func (s *ContainerdConfig) SetLimitCore(v int64) *ContainerdConfig {
	s.LimitCore = &v
	return s
}

func (s *ContainerdConfig) SetLimitMemLock(v int64) *ContainerdConfig {
	s.LimitMemLock = &v
	return s
}

func (s *ContainerdConfig) SetLimitNoFile(v int64) *ContainerdConfig {
	s.LimitNoFile = &v
	return s
}

func (s *ContainerdConfig) SetMaxConcurrentDownloads(v int64) *ContainerdConfig {
	s.MaxConcurrentDownloads = &v
	return s
}

func (s *ContainerdConfig) SetRegistryMirrors(v []*string) *ContainerdConfig {
	s.RegistryMirrors = v
	return s
}

func (s *ContainerdConfig) Validate() error {
	return dara.Validate(s)
}
