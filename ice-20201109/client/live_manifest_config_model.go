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
	// The type of ad markers to include in the manifest.
	//
	// 	- NONE: Removes all ad markers.
	//
	// 	- DATE_RANGE: Inserts EXT-X-DATERANGE tags (HLS spec). Valid for HLS/HLS-CMAF endpoints.
	//
	// 	- XML: Inserts XML-based ad markers (DASH spec). Valid for DASH endpoints.
	//
	// example:
	//
	// NONE
	AdMarkers *string `json:"AdMarkers,omitempty" xml:"AdMarkers,omitempty"`
	// The interval, in seconds, at which to insert the EXT-X-PROGRAM-DATE-TIME tag into the playlist. By default, no tags are inserted. Valid values: 1 to 3600. Applies only to HLS and HLS-CMAF endpoints.
	//
	// example:
	//
	// 5
	DateTimeInterval *int32 `json:"DateTimeInterval,omitempty" xml:"DateTimeInterval,omitempty"`
	// The duration of the startover window, in seconds. It defines the maximum time a viewer can seek backward in the live stream. Valid values: 1 to 3600. Default value: 60. Applies only to DASH endpoints.
	//
	// example:
	//
	// 60
	ManifestDuration *int32 `json:"ManifestDuration,omitempty" xml:"ManifestDuration,omitempty"`
	// The maximum bitrate threshold (in bits per second) that video tracks must be at or below to be available for playback from this endpoint. It must be a positive integer. If not set, no maximum bitrate is enforced.
	//
	// example:
	//
	// 8000000
	MaxVideoBitrate *int32 `json:"MaxVideoBitrate,omitempty" xml:"MaxVideoBitrate,omitempty"`
	// The minimum buffer time, in seconds. Valid values: 1 to 30. Default value: the duration of two segments. Applies only to DASH endpoints.
	//
	// Note: Setting this value too low may cause playback to stutter. We recommend a value no less than two segment durations.
	//
	// example:
	//
	// 8
	MinBufferTime *int32 `json:"MinBufferTime,omitempty" xml:"MinBufferTime,omitempty"`
	// The minimum update period for the manifest, in seconds. Valid values: 1 to 3600. Default value: the duration of two segments. Applies only to DASH endpoints.
	//
	// Note: For smooth playback, set this value to be less than MinBufferTime.
	//
	// example:
	//
	// 8
	MinUpdatePeriod *int32 `json:"MinUpdatePeriod,omitempty" xml:"MinUpdatePeriod,omitempty"`
	// The minimum bitrate threshold (in bits per second) that video tracks must be at or above to be available for playback from this endpoint. It must be a positive integer. If not set, no minimum bitrate is enforced.
	//
	// example:
	//
	// 1000000
	MinVideoBitrate *int32 `json:"MinVideoBitrate,omitempty" xml:"MinVideoBitrate,omitempty"`
	// The suggested presentation delay, in seconds. Valid values: 1 to 60. Default value: the duration of three segments.
	//
	// example:
	//
	// 12
	PresentationDelay *int32 `json:"PresentationDelay,omitempty" xml:"PresentationDelay,omitempty"`
	// The number of segments to include in the playlist. Applies to HLS and HLS-CMAF protocols. If not set, the channel\\"s default configuration is used. Valid values: 2 to 100.
	//
	// example:
	//
	// 3
	SegmentCount *int32 `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	// The format of the segment template. Only NUMBER_TIMELINE is supported (default). Applies only to DASH endpoints.
	//
	// example:
	//
	// NUMBER_TIMELINE
	SegmentTemplateFormat *string `json:"SegmentTemplateFormat,omitempty" xml:"SegmentTemplateFormat,omitempty"`
	// The order of streams in the master playlist. Valid values:
	//
	// 	- ORIGINAL: Preserves the original order of the input streams.
	//
	// 	- VIDEO_BITRATE_ASCENDING: sorts the streams in ascending order of bitrates, from lowest to highest.
	//
	// 	- VIDEO_BITRATE_DESCENDING: sorts the streams in descending order of bitrates, from highest to lowest.
	//
	// example:
	//
	// ORIGINAL
	StreamOrder *string `json:"StreamOrder,omitempty" xml:"StreamOrder,omitempty"`
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
