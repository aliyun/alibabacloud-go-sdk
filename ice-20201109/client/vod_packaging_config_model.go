// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVodPackagingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDrmProvider(v *VodPackagingConfigDrmProvider) *VodPackagingConfig
	GetDrmProvider() *VodPackagingConfigDrmProvider
	SetManifestName(v string) *VodPackagingConfig
	GetManifestName() *string
	SetSegmentDuration(v int64) *VodPackagingConfig
	GetSegmentDuration() *int64
	SetStreamSelection(v *VodPackagingConfigStreamSelection) *VodPackagingConfig
	GetStreamSelection() *VodPackagingConfigStreamSelection
}

type VodPackagingConfig struct {
	DrmProvider     *VodPackagingConfigDrmProvider     `json:"DrmProvider,omitempty" xml:"DrmProvider,omitempty" type:"Struct"`
	ManifestName    *string                            `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	SegmentDuration *int64                             `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	StreamSelection *VodPackagingConfigStreamSelection `json:"StreamSelection,omitempty" xml:"StreamSelection,omitempty" type:"Struct"`
}

func (s VodPackagingConfig) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingConfig) GoString() string {
	return s.String()
}

func (s *VodPackagingConfig) GetDrmProvider() *VodPackagingConfigDrmProvider {
	return s.DrmProvider
}

func (s *VodPackagingConfig) GetManifestName() *string {
	return s.ManifestName
}

func (s *VodPackagingConfig) GetSegmentDuration() *int64 {
	return s.SegmentDuration
}

func (s *VodPackagingConfig) GetStreamSelection() *VodPackagingConfigStreamSelection {
	return s.StreamSelection
}

func (s *VodPackagingConfig) SetDrmProvider(v *VodPackagingConfigDrmProvider) *VodPackagingConfig {
	s.DrmProvider = v
	return s
}

func (s *VodPackagingConfig) SetManifestName(v string) *VodPackagingConfig {
	s.ManifestName = &v
	return s
}

func (s *VodPackagingConfig) SetSegmentDuration(v int64) *VodPackagingConfig {
	s.SegmentDuration = &v
	return s
}

func (s *VodPackagingConfig) SetStreamSelection(v *VodPackagingConfigStreamSelection) *VodPackagingConfig {
	s.StreamSelection = v
	return s
}

func (s *VodPackagingConfig) Validate() error {
	return dara.Validate(s)
}

type VodPackagingConfigDrmProvider struct {
	EncryptionMethod *string   `json:"EncryptionMethod,omitempty" xml:"EncryptionMethod,omitempty"`
	IV               *string   `json:"IV,omitempty" xml:"IV,omitempty"`
	SystemIds        []*string `json:"SystemIds,omitempty" xml:"SystemIds,omitempty" type:"Repeated"`
	Url              *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VodPackagingConfigDrmProvider) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingConfigDrmProvider) GoString() string {
	return s.String()
}

func (s *VodPackagingConfigDrmProvider) GetEncryptionMethod() *string {
	return s.EncryptionMethod
}

func (s *VodPackagingConfigDrmProvider) GetIV() *string {
	return s.IV
}

func (s *VodPackagingConfigDrmProvider) GetSystemIds() []*string {
	return s.SystemIds
}

func (s *VodPackagingConfigDrmProvider) GetUrl() *string {
	return s.Url
}

func (s *VodPackagingConfigDrmProvider) SetEncryptionMethod(v string) *VodPackagingConfigDrmProvider {
	s.EncryptionMethod = &v
	return s
}

func (s *VodPackagingConfigDrmProvider) SetIV(v string) *VodPackagingConfigDrmProvider {
	s.IV = &v
	return s
}

func (s *VodPackagingConfigDrmProvider) SetSystemIds(v []*string) *VodPackagingConfigDrmProvider {
	s.SystemIds = v
	return s
}

func (s *VodPackagingConfigDrmProvider) SetUrl(v string) *VodPackagingConfigDrmProvider {
	s.Url = &v
	return s
}

func (s *VodPackagingConfigDrmProvider) Validate() error {
	return dara.Validate(s)
}

type VodPackagingConfigStreamSelection struct {
	MaxVideoBitsPerSecond *int64  `json:"MaxVideoBitsPerSecond,omitempty" xml:"MaxVideoBitsPerSecond,omitempty"`
	MinVideoBitsPerSecond *int64  `json:"MinVideoBitsPerSecond,omitempty" xml:"MinVideoBitsPerSecond,omitempty"`
	StreamOrder           *string `json:"StreamOrder,omitempty" xml:"StreamOrder,omitempty"`
}

func (s VodPackagingConfigStreamSelection) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingConfigStreamSelection) GoString() string {
	return s.String()
}

func (s *VodPackagingConfigStreamSelection) GetMaxVideoBitsPerSecond() *int64 {
	return s.MaxVideoBitsPerSecond
}

func (s *VodPackagingConfigStreamSelection) GetMinVideoBitsPerSecond() *int64 {
	return s.MinVideoBitsPerSecond
}

func (s *VodPackagingConfigStreamSelection) GetStreamOrder() *string {
	return s.StreamOrder
}

func (s *VodPackagingConfigStreamSelection) SetMaxVideoBitsPerSecond(v int64) *VodPackagingConfigStreamSelection {
	s.MaxVideoBitsPerSecond = &v
	return s
}

func (s *VodPackagingConfigStreamSelection) SetMinVideoBitsPerSecond(v int64) *VodPackagingConfigStreamSelection {
	s.MinVideoBitsPerSecond = &v
	return s
}

func (s *VodPackagingConfigStreamSelection) SetStreamOrder(v string) *VodPackagingConfigStreamSelection {
	s.StreamOrder = &v
	return s
}

func (s *VodPackagingConfigStreamSelection) Validate() error {
	return dara.Validate(s)
}
