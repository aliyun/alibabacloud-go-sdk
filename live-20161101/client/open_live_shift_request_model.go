// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLiveShiftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OpenLiveShiftRequest
	GetAppName() *string
	SetDomainName(v string) *OpenLiveShiftRequest
	GetDomainName() *string
	SetDuration(v int32) *OpenLiveShiftRequest
	GetDuration() *int32
	SetIgnoreTranscode(v bool) *OpenLiveShiftRequest
	GetIgnoreTranscode() *bool
	SetOwnerId(v int64) *OpenLiveShiftRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenLiveShiftRequest
	GetRegionId() *string
	SetStreamName(v string) *OpenLiveShiftRequest
	GetStreamName() *string
	SetVision(v int32) *OpenLiveShiftRequest
	GetVision() *int32
}

type OpenLiveShiftRequest struct {
	// The name of the application. The wildcard character (\\*) is supported. An asterisk (\\*) represents all applications under the specified domain name. For more information, see [Stream management](https://help.aliyun.com/document_detail/197397.html).
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The duration of an HTTP Live Streaming (HLS) transport stream (TS) segment. Unit: seconds.
	//
	// example:
	//
	// 3
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable time shifting for transcoded streams. Valid values:
	//
	// - **true**: Time shifting is disabled for transcoded streams.
	//
	// - **false**: Time shifting is enabled for transcoded streams.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	IgnoreTranscode *bool  `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stream. The wildcard character (\\*) is supported. An asterisk (\\*) represents all streams under the specified application. For more information, see [Stream management](https://help.aliyun.com/document_detail/197397.html).
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The data retention period. The default value is 7. Unit: days.
	//
	// example:
	//
	// 10
	Vision *int32 `json:"Vision,omitempty" xml:"Vision,omitempty"`
}

func (s OpenLiveShiftRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenLiveShiftRequest) GoString() string {
	return s.String()
}

func (s *OpenLiveShiftRequest) GetAppName() *string {
	return s.AppName
}

func (s *OpenLiveShiftRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *OpenLiveShiftRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *OpenLiveShiftRequest) GetIgnoreTranscode() *bool {
	return s.IgnoreTranscode
}

func (s *OpenLiveShiftRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenLiveShiftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenLiveShiftRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *OpenLiveShiftRequest) GetVision() *int32 {
	return s.Vision
}

func (s *OpenLiveShiftRequest) SetAppName(v string) *OpenLiveShiftRequest {
	s.AppName = &v
	return s
}

func (s *OpenLiveShiftRequest) SetDomainName(v string) *OpenLiveShiftRequest {
	s.DomainName = &v
	return s
}

func (s *OpenLiveShiftRequest) SetDuration(v int32) *OpenLiveShiftRequest {
	s.Duration = &v
	return s
}

func (s *OpenLiveShiftRequest) SetIgnoreTranscode(v bool) *OpenLiveShiftRequest {
	s.IgnoreTranscode = &v
	return s
}

func (s *OpenLiveShiftRequest) SetOwnerId(v int64) *OpenLiveShiftRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenLiveShiftRequest) SetRegionId(v string) *OpenLiveShiftRequest {
	s.RegionId = &v
	return s
}

func (s *OpenLiveShiftRequest) SetStreamName(v string) *OpenLiveShiftRequest {
	s.StreamName = &v
	return s
}

func (s *OpenLiveShiftRequest) SetVision(v int32) *OpenLiveShiftRequest {
	s.Vision = &v
	return s
}

func (s *OpenLiveShiftRequest) Validate() error {
	return dara.Validate(s)
}
