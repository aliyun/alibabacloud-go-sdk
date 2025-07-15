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
	// >  This parameter is required if Protocol is set to LLHLS_TS or LLHLS_CMAF.
	//
	// 	- If the value of SegmentDuration is 1, the valid values of this parameter are 100 to 500.
	//
	// 	- If the value of SegmentDuration is 2, the valid values of this parameter are 100 to 1000.
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
	// 	- If Protocol is set to HLS_CMAF, the valid values of this parameter are 1 to 10.
	//
	// 	- If Protocol is set to LLHLS_TS or LLHLS_CMAF, the valid values of this parameter are 1 to 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The number of segments. Valid values: 3 to 10.
	//
	// This parameter is required.
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
