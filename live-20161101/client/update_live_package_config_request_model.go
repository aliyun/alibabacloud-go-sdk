// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLivePackageConfigRequest
	GetAppName() *string
	SetDomainName(v string) *UpdateLivePackageConfigRequest
	GetDomainName() *string
	SetIgnoreTranscode(v bool) *UpdateLivePackageConfigRequest
	GetIgnoreTranscode() *bool
	SetOwnerId(v int64) *UpdateLivePackageConfigRequest
	GetOwnerId() *int64
	SetPartDuration(v int32) *UpdateLivePackageConfigRequest
	GetPartDuration() *int32
	SetProtocol(v string) *UpdateLivePackageConfigRequest
	GetProtocol() *string
	SetRegionId(v string) *UpdateLivePackageConfigRequest
	GetRegionId() *string
	SetSegmentDuration(v int32) *UpdateLivePackageConfigRequest
	GetSegmentDuration() *int32
	SetSegmentNum(v int32) *UpdateLivePackageConfigRequest
	GetSegmentNum() *int32
	SetStreamName(v string) *UpdateLivePackageConfigRequest
	GetStreamName() *string
}

type UpdateLivePackageConfigRequest struct {
	// The application name. The template applies only when this AppName matches the application name in the ingest URL. The AppName can be up to 255 characters and can contain digits, letters, hyphens (-), and underscores (_). It cannot start with a hyphen or an underscore. Set this parameter to an asterisk (\\*) to match all application names.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The primary domain name for live streaming playback.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to ignore transcoded streams. Valid values:
	//
	// - **true*	- (default): Ignore transcoded streams.
	//
	// - **false**: Do not ignore transcoded streams.
	//
	// example:
	//
	// true
	IgnoreTranscode *bool  `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The duration of a part segment in milliseconds.
	//
	// > This parameter is required if you set \\`Protocol\\` to \\`LLHLS_\\*\\`.
	//
	// - If SegmentDuration is 1 s, the value can range from 100 to 500 ms.
	//
	// - If SegmentDuration is 2 s, the value can range from 100 to 1000 ms.
	//
	// example:
	//
	// 350
	PartDuration *int32 `json:"PartDuration,omitempty" xml:"PartDuration,omitempty"`
	// The protocol and container format for live streaming. Valid values:
	//
	// - **HLS_CMAF**
	//
	// - **LLHLS_TS*	- (low latency)
	//
	// - **LLHLS_CMAF*	- (low latency)
	//
	// - **DASH_CMAF**
	//
	// - **HLSDASH_CMAF**
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS_CMAF
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The segment duration in seconds.
	//
	// - If you set Protocol to HLS_CMAF, the value can range from 1 to 10 s.
	//
	// - If you set Protocol to LLHLS_\\*, the value can range from 1 to 2 s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The number of M3U8 segments. The value must be an integer from 3 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SegmentNum *int32 `json:"SegmentNum,omitempty" xml:"SegmentNum,omitempty"`
	// The stream name. The template applies only when this StreamName matches the stream name in the ingest URL. The StreamName can be up to 255 characters and can contain digits, letters, hyphens (-), and underscores (_). It cannot start with a hyphen or an underscore. Set this parameter to an asterisk (\\*) to match all stream names.
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s UpdateLivePackageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLivePackageConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLivePackageConfigRequest) GetIgnoreTranscode() *bool {
	return s.IgnoreTranscode
}

func (s *UpdateLivePackageConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLivePackageConfigRequest) GetPartDuration() *int32 {
	return s.PartDuration
}

func (s *UpdateLivePackageConfigRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLivePackageConfigRequest) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *UpdateLivePackageConfigRequest) GetSegmentNum() *int32 {
	return s.SegmentNum
}

func (s *UpdateLivePackageConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLivePackageConfigRequest) SetAppName(v string) *UpdateLivePackageConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetDomainName(v string) *UpdateLivePackageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetIgnoreTranscode(v bool) *UpdateLivePackageConfigRequest {
	s.IgnoreTranscode = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetOwnerId(v int64) *UpdateLivePackageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetPartDuration(v int32) *UpdateLivePackageConfigRequest {
	s.PartDuration = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetProtocol(v string) *UpdateLivePackageConfigRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetRegionId(v string) *UpdateLivePackageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetSegmentDuration(v int32) *UpdateLivePackageConfigRequest {
	s.SegmentDuration = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetSegmentNum(v int32) *UpdateLivePackageConfigRequest {
	s.SegmentNum = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) SetStreamName(v string) *UpdateLivePackageConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLivePackageConfigRequest) Validate() error {
	return dara.Validate(s)
}
