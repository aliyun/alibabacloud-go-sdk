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
	SetPartHoldBackMs(v int32) *LiveManifestConfig
	GetPartHoldBackMs() *int32
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
	// The ad markers supported in the playlist. Valid values:
	//
	// - NONE: removes ad markers.
	//
	// - DATE_RANGE: uses the EXT-X-DATERANGE tag defined in the HLS specification. This value is available when the endpoint protocol is HLS/HLS_CMAF.
	//
	// - XML: uses the XML ad markers defined in the DASH specification. This value is available when the endpoint protocol is DASH.
	//
	// example:
	//
	// NONE
	AdMarkers *string `json:"AdMarkers,omitempty" xml:"AdMarkers,omitempty"`
	// The interval (in seconds) for inserting the EXT-X-PROGRAM-DATE-TIME time tag. By default, the tag is not inserted. Valid values: 1 to 3600. This parameter applies to the HLS/HLS_CMAF protocol.
	//
	// example:
	//
	// 5
	DateTimeInterval *int32 `json:"DateTimeInterval,omitempty" xml:"DateTimeInterval,omitempty"`
	// The maximum time-shift duration during live streaming. Unit: seconds. Valid values: 1 to 3600. Default value: 60. This parameter applies to DASH.
	//
	// example:
	//
	// 60
	ManifestDuration *int32 `json:"ManifestDuration,omitempty" xml:"ManifestDuration,omitempty"`
	// The maximum input bitrate threshold (unit: bit/s). A video track must have a bitrate less than or equal to this threshold to be played from this endpoint. Valid values: integers greater than 0. By default, this parameter is empty and no maximum bitrate limit is set.
	//
	// example:
	//
	// 8000000
	MaxVideoBitrate *int32 `json:"MaxVideoBitrate,omitempty" xml:"MaxVideoBitrate,omitempty"`
	// The minimum buffer time. Unit: seconds. Valid values: 1 to 30. Default value: 2 segment durations. This parameter applies only to DASH.
	//
	// >  An excessively small minimum buffer time may cause playback stuttering. Set this parameter to a value no less than 2 segment durations.
	//
	// example:
	//
	// 8
	MinBufferTime *int32 `json:"MinBufferTime,omitempty" xml:"MinBufferTime,omitempty"`
	// The minimum update interval. Unit: seconds. Valid values: 1 to 3600. Default value: 2 segment durations. This parameter applies to DASH.
	//
	// >  Set this parameter to a value less than the minimum buffer time. An excessively large value may cause DASH playback stuttering.
	//
	// example:
	//
	// 8
	MinUpdatePeriod *int32 `json:"MinUpdatePeriod,omitempty" xml:"MinUpdatePeriod,omitempty"`
	// The minimum input bitrate threshold (unit: bit/s). A video track must have a bitrate greater than or equal to this threshold to be played from this endpoint. Valid values: integers greater than 0. By default, this parameter is empty and no minimum bitrate is set.
	//
	// example:
	//
	// 1000000
	MinVideoBitrate *int32 `json:"MinVideoBitrate,omitempty" xml:"MinVideoBitrate,omitempty"`
	PartHoldBackMs  *int32 `json:"PartHoldBackMs,omitempty" xml:"PartHoldBackMs,omitempty"`
	// The suggested presentation delay. Unit: seconds. Valid values: 1 to 60. Default value: 3 segment durations.
	//
	// example:
	//
	// 12
	PresentationDelay *int32 `json:"PresentationDelay,omitempty" xml:"PresentationDelay,omitempty"`
	// The number of segments. This parameter applies to the HLS/HLS_CMAF protocol. By default, the channel configuration is used. Valid values: 2 to 100.
	//
	// example:
	//
	// 3
	SegmentCount *int32 `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	// The segment template. Currently, only NUMBER_TIMELINE (default) is supported. This parameter applies to DASH.
	//
	// example:
	//
	// NUMBER_TIMELINE
	SegmentTemplateFormat *string `json:"SegmentTemplateFormat,omitempty" xml:"SegmentTemplateFormat,omitempty"`
	// The stream sorting rule. Valid values:
	//
	// - ORIGINAL: retains the original order of the input sub-manifest.
	//
	// - VIDEO_BITRATE_ASCENDING: sorts by video stream bitrate in ascending order.
	//
	// - VIDEO_BITRATE_DESCENDING: sorts by video stream bitrate in descending order.
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

func (s *LiveManifestConfig) GetPartHoldBackMs() *int32 {
	return s.PartHoldBackMs
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

func (s *LiveManifestConfig) SetPartHoldBackMs(v int32) *LiveManifestConfig {
	s.PartHoldBackMs = &v
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
