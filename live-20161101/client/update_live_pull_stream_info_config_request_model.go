// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLivePullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *UpdateLivePullStreamInfoConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *UpdateLivePullStreamInfoConfigRequest
	GetEndTime() *string
	SetOwnerId(v int64) *UpdateLivePullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLivePullStreamInfoConfigRequest
	GetRegionId() *string
	SetSourceUrl(v string) *UpdateLivePullStreamInfoConfigRequest
	GetSourceUrl() *string
	SetStartTime(v string) *UpdateLivePullStreamInfoConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *UpdateLivePullStreamInfoConfigRequest
	GetStreamName() *string
}

type UpdateLivePullStreamInfoConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain used as the streamer\\"s stream domain.
	//
	// > - When you specify DomainName, make sure that the domain is a live streaming domain and that the user calling this operation has the permission to operate on the specified domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of stream pulling.
	//
	// The interval between StartTime and EndTime must be within 7 days, and EndTime must be later than the current time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The full URL of the origin server where the live stream resides. You can specify multiple origin URLs separated by semicolons (;).
	//
	// This parameter is required.
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// The start time of stream pulling.
	//
	// The interval between StartTime and EndTime must be within 7 days. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s UpdateLivePullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLivePullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetAppName(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetDomainName(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetEndTime(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetOwnerId(v int64) *UpdateLivePullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetRegionId(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetSourceUrl(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetStartTime(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) SetStreamName(v string) *UpdateLivePullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
