// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLivePackageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLivePackageConfigRequest
	GetAppName() *string
	SetDomainName(v string) *AddLivePackageConfigRequest
	GetDomainName() *string
	SetIgnoreTranscode(v bool) *AddLivePackageConfigRequest
	GetIgnoreTranscode() *bool
	SetOwnerId(v int64) *AddLivePackageConfigRequest
	GetOwnerId() *int64
	SetPartDuration(v int32) *AddLivePackageConfigRequest
	GetPartDuration() *int32
	SetProtocol(v string) *AddLivePackageConfigRequest
	GetProtocol() *string
	SetRegionId(v string) *AddLivePackageConfigRequest
	GetRegionId() *string
	SetSegmentDuration(v int32) *AddLivePackageConfigRequest
	GetSegmentDuration() *int32
	SetSegmentNum(v int32) *AddLivePackageConfigRequest
	GetSegmentNum() *int32
	SetStreamName(v string) *AddLivePackageConfigRequest
	GetStreamName() *string
}

type AddLivePackageConfigRequest struct {
	// The AppName must match the AppName in the ingest URL for the template to take effect. The AppName can be up to 255 characters in length and can contain digits, uppercase letters, lowercase letters, hyphens (-), and underscores (_). A hyphen or an underscore cannot be the first character. You can also set this parameter to a single asterisk (\\*) to match all AppNames.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The live streaming domain name. This is the primary playback domain name.
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
	// The length of the part segment, in milliseconds.
	//
	// - If SegmentDuration is 1 s, the valid values are 100 to 500. The default value is 350.
	//
	// - If SegmentDuration is 2 s, the valid values are 100 to 1000. The default value is 700.
	//
	// - This parameter is valid only when you set Protocol to LLHLS_\\*.
	//
	// example:
	//
	// 350
	PartDuration *int32 `json:"PartDuration,omitempty" xml:"PartDuration,omitempty"`
	// The live streaming protocol and container format. Valid values:
	//
	// - **HLS_CMAF**
	//
	// - **LLHLS_TS**
	//
	//   (low latency)
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
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The segment length, in seconds.
	//
	// - If you set Protocol to HLS_CMAF, the valid values are 1 to 10. The default value is 5.
	//
	// - If you set Protocol to LLHLS_\\*, the valid values are 1 to 2. The default value is 1.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The number of M3U8 segments for live streaming.
	//
	// - Valid values: 3 to 10.
	//
	// - Default value: 3.
	//
	// example:
	//
	// 3
	SegmentNum *int32 `json:"SegmentNum,omitempty" xml:"SegmentNum,omitempty"`
	// The StreamName must match the StreamName in the ingest URL for the template to take effect. The StreamName can be up to 255 characters in length and can contain digits, uppercase letters, lowercase letters, hyphens (-), and underscores (_). A hyphen or an underscore cannot be the first character. You can also set this parameter to a single asterisk (\\*) to match all StreamNames.
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddLivePackageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLivePackageConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLivePackageConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLivePackageConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLivePackageConfigRequest) GetIgnoreTranscode() *bool {
	return s.IgnoreTranscode
}

func (s *AddLivePackageConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLivePackageConfigRequest) GetPartDuration() *int32 {
	return s.PartDuration
}

func (s *AddLivePackageConfigRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddLivePackageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLivePackageConfigRequest) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *AddLivePackageConfigRequest) GetSegmentNum() *int32 {
	return s.SegmentNum
}

func (s *AddLivePackageConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLivePackageConfigRequest) SetAppName(v string) *AddLivePackageConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetDomainName(v string) *AddLivePackageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetIgnoreTranscode(v bool) *AddLivePackageConfigRequest {
	s.IgnoreTranscode = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetOwnerId(v int64) *AddLivePackageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetPartDuration(v int32) *AddLivePackageConfigRequest {
	s.PartDuration = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetProtocol(v string) *AddLivePackageConfigRequest {
	s.Protocol = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetRegionId(v string) *AddLivePackageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetSegmentDuration(v int32) *AddLivePackageConfigRequest {
	s.SegmentDuration = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetSegmentNum(v int32) *AddLivePackageConfigRequest {
	s.SegmentNum = &v
	return s
}

func (s *AddLivePackageConfigRequest) SetStreamName(v string) *AddLivePackageConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLivePackageConfigRequest) Validate() error {
	return dara.Validate(s)
}
