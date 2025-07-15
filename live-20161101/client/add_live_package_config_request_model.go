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
	// The application name. The value of this parameter must be the same as the application name that is specified in the ingest URL. Otherwise, the configuration does not take effect. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name cannot start with a hyphen (-) or underscore (_). You can also specify an asterisk (\\*) as the value to match all applications.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to ignore the transcoded stream. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IgnoreTranscode *bool  `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The part length. Unit: milliseconds.
	//
	// 	- If the value of SegmentDuration is 1, the valid values of this parameter are 100 to 500 and the default value of this parameter is 350.
	//
	// 	- If the value of SegmentDuration is 2, the valid values of this parameter are 100 to 1000 and the default value of this parameter is 700.
	//
	// 	- This parameter takes effect only if Protocol is set to LLHLS_TS or LLHLS_CMAF.
	//
	// example:
	//
	// 350
	PartDuration *int32 `json:"PartDuration,omitempty" xml:"PartDuration,omitempty"`
	// The streaming protocol and encapsulation format. Valid values:
	//
	// 	- **HLS_CMAF**
	//
	// 	- **LLHLS_TS*	- (low latency)
	//
	// 	- **LLHLS_CMAF*	- (low latency)
	//
	// 	- **DASH_CMAF**
	//
	// 	- **HLSDASH_CMAF**
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS_CMAF
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The segment length. Unit: seconds.
	//
	// 	- If Protocol is set to HLS_CMAF: Valid values: 1 to 10. Default value: 5.
	//
	// 	- If Protocol is set to LLHLS_TS or LLHLS_CMAF: Valid values: 1 to 2. Default value: 1.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The number of segments.
	//
	// 	- Valid values: 3 to 10.
	//
	// 	- Default value: 3.
	//
	// example:
	//
	// 3
	SegmentNum *int32 `json:"SegmentNum,omitempty" xml:"SegmentNum,omitempty"`
	// The stream name. The value of this parameter must be the same as the stream name that is specified in the ingest URL. Otherwise, the configuration does not take effect. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name cannot start with a hyphen (-) or underscore (_). You can also specify an asterisk (\\*) as the value to match all streams.
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
