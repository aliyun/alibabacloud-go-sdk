// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveManifestConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdMarkers(v string) *LiveManifestConfig
	GetAdMarkers() *string
	SetDateTimeInterval(v int32) *LiveManifestConfig
	GetDateTimeInterval() *int32
	SetManifestDuration(v int32) *LiveManifestConfig
	GetManifestDuration() *int32
	SetMaxVideoBitrate(v int32) *LiveManifestConfig
	GetMaxVideoBitrate() *int32
	SetMinBufferTime(v int32) *LiveManifestConfig
	GetMinBufferTime() *int32
	SetMinUpdatePeriod(v int32) *LiveManifestConfig
	GetMinUpdatePeriod() *int32
	SetMinVideoBitrate(v int32) *LiveManifestConfig
	GetMinVideoBitrate() *int32
	SetPresentationDelay(v int32) *LiveManifestConfig
	GetPresentationDelay() *int32
	SetSegmentCount(v int32) *LiveManifestConfig
	GetSegmentCount() *int32
	SetSegmentTemplateFormat(v string) *LiveManifestConfig
	GetSegmentTemplateFormat() *string
	SetStreamOrder(v string) *LiveManifestConfig
	GetStreamOrder() *string
}

type LiveManifestConfig struct {
	AdMarkers             *string `json:"AdMarkers,omitempty" xml:"AdMarkers,omitempty"`
	DateTimeInterval      *int32  `json:"DateTimeInterval,omitempty" xml:"DateTimeInterval,omitempty"`
	ManifestDuration      *int32  `json:"ManifestDuration,omitempty" xml:"ManifestDuration,omitempty"`
	MaxVideoBitrate       *int32  `json:"MaxVideoBitrate,omitempty" xml:"MaxVideoBitrate,omitempty"`
	MinBufferTime         *int32  `json:"MinBufferTime,omitempty" xml:"MinBufferTime,omitempty"`
	MinUpdatePeriod       *int32  `json:"MinUpdatePeriod,omitempty" xml:"MinUpdatePeriod,omitempty"`
	MinVideoBitrate       *int32  `json:"MinVideoBitrate,omitempty" xml:"MinVideoBitrate,omitempty"`
	PresentationDelay     *int32  `json:"PresentationDelay,omitempty" xml:"PresentationDelay,omitempty"`
	SegmentCount          *int32  `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	SegmentTemplateFormat *string `json:"SegmentTemplateFormat,omitempty" xml:"SegmentTemplateFormat,omitempty"`
	StreamOrder           *string `json:"StreamOrder,omitempty" xml:"StreamOrder,omitempty"`
}

func (s LiveManifestConfig) String() string {
	return dara.Prettify(s)
}

func (s LiveManifestConfig) GoString() string {
	return s.String()
}

func (s *LiveManifestConfig) GetAdMarkers() *string {
	return s.AdMarkers
}

func (s *LiveManifestConfig) GetDateTimeInterval() *int32 {
	return s.DateTimeInterval
}

func (s *LiveManifestConfig) GetManifestDuration() *int32 {
	return s.ManifestDuration
}

func (s *LiveManifestConfig) GetMaxVideoBitrate() *int32 {
	return s.MaxVideoBitrate
}

func (s *LiveManifestConfig) GetMinBufferTime() *int32 {
	return s.MinBufferTime
}

func (s *LiveManifestConfig) GetMinUpdatePeriod() *int32 {
	return s.MinUpdatePeriod
}

func (s *LiveManifestConfig) GetMinVideoBitrate() *int32 {
	return s.MinVideoBitrate
}

func (s *LiveManifestConfig) GetPresentationDelay() *int32 {
	return s.PresentationDelay
}

func (s *LiveManifestConfig) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *LiveManifestConfig) GetSegmentTemplateFormat() *string {
	return s.SegmentTemplateFormat
}

func (s *LiveManifestConfig) GetStreamOrder() *string {
	return s.StreamOrder
}

func (s *LiveManifestConfig) SetAdMarkers(v string) *LiveManifestConfig {
	s.AdMarkers = &v
	return s
}

func (s *LiveManifestConfig) SetDateTimeInterval(v int32) *LiveManifestConfig {
	s.DateTimeInterval = &v
	return s
}

func (s *LiveManifestConfig) SetManifestDuration(v int32) *LiveManifestConfig {
	s.ManifestDuration = &v
	return s
}

func (s *LiveManifestConfig) SetMaxVideoBitrate(v int32) *LiveManifestConfig {
	s.MaxVideoBitrate = &v
	return s
}

func (s *LiveManifestConfig) SetMinBufferTime(v int32) *LiveManifestConfig {
	s.MinBufferTime = &v
	return s
}

func (s *LiveManifestConfig) SetMinUpdatePeriod(v int32) *LiveManifestConfig {
	s.MinUpdatePeriod = &v
	return s
}

func (s *LiveManifestConfig) SetMinVideoBitrate(v int32) *LiveManifestConfig {
	s.MinVideoBitrate = &v
	return s
}

func (s *LiveManifestConfig) SetPresentationDelay(v int32) *LiveManifestConfig {
	s.PresentationDelay = &v
	return s
}

func (s *LiveManifestConfig) SetSegmentCount(v int32) *LiveManifestConfig {
	s.SegmentCount = &v
	return s
}

func (s *LiveManifestConfig) SetSegmentTemplateFormat(v string) *LiveManifestConfig {
	s.SegmentTemplateFormat = &v
	return s
}

func (s *LiveManifestConfig) SetStreamOrder(v string) *LiveManifestConfig {
	s.StreamOrder = &v
	return s
}

func (s *LiveManifestConfig) Validate() error {
	return dara.Validate(s)
}
