// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamMergeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveStreamMergeRequest
	GetAppName() *string
	SetDomainName(v string) *AddLiveStreamMergeRequest
	GetDomainName() *string
	SetEndTime(v string) *AddLiveStreamMergeRequest
	GetEndTime() *string
	SetInAppName1(v string) *AddLiveStreamMergeRequest
	GetInAppName1() *string
	SetInAppName2(v string) *AddLiveStreamMergeRequest
	GetInAppName2() *string
	SetInStreamName1(v string) *AddLiveStreamMergeRequest
	GetInStreamName1() *string
	SetInStreamName2(v string) *AddLiveStreamMergeRequest
	GetInStreamName2() *string
	SetLiveMerger(v string) *AddLiveStreamMergeRequest
	GetLiveMerger() *string
	SetMergeParameters(v string) *AddLiveStreamMergeRequest
	GetMergeParameters() *string
	SetOwnerId(v int64) *AddLiveStreamMergeRequest
	GetOwnerId() *int64
	SetProtocol(v string) *AddLiveStreamMergeRequest
	GetProtocol() *string
	SetRegionId(v string) *AddLiveStreamMergeRequest
	GetRegionId() *string
	SetSelectAppName(v string) *AddLiveStreamMergeRequest
	GetSelectAppName() *string
	SetSelectStreamName(v string) *AddLiveStreamMergeRequest
	GetSelectStreamName() *string
	SetStartTime(v string) *AddLiveStreamMergeRequest
	GetStartTime() *string
	SetStreamName(v string) *AddLiveStreamMergeRequest
	GetStreamName() *string
	SetSwitchMode(v string) *AddLiveStreamMergeRequest
	GetSwitchMode() *string
}

type AddLiveStreamMergeRequest struct {
	// The AppName of the output stream. For the configuration to take effect, this AppName must match the one in the ingest URL. Wildcards (`*`) are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the stream merge.
	//
	// The time must be in UTC and specified in the ISO 8601 standard format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// > The interval between `StartTime` and `EndTime` cannot exceed 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-29T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The AppName of the primary input stream. This value must match the AppName in the ingest URL for the primary stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// app1
	InAppName1 *string `json:"InAppName1,omitempty" xml:"InAppName1,omitempty"`
	// The AppName of the backup input stream. This value must match the AppName in the ingest URL for the backup stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// app2
	InAppName2 *string `json:"InAppName2,omitempty" xml:"InAppName2,omitempty"`
	// The StreamName of the primary input stream. This value must match the StreamName in the ingest URL for the primary stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// InStream1
	InStreamName1 *string `json:"InStreamName1,omitempty" xml:"InStreamName1,omitempty"`
	// The StreamName of the backup input stream. This value must match the StreamName in the ingest URL for the backup stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// stream2
	InStreamName2 *string `json:"InStreamName2,omitempty" xml:"InStreamName2,omitempty"`
	// The engine to use for stream merging.
	//
	// - `on`: The new liveswitch engine.
	//
	// - `off`: A legacy engine (such as rtmpr). This is the default.
	//
	// example:
	//
	// off
	LiveMerger *string `json:"LiveMerger,omitempty" xml:"LiveMerger,omitempty"`
	// Parameters that define the failover conditions. A failover is triggered when one of the following conditions is met:
	//
	// 1. An explicit stream disconnection occurs, such as an end-of-file (EOF) or network error.
	//
	// 2. The stutter rate exceeds 60% in the last 5 seconds.
	//
	// 3. A stream pulling timeout occurs if no frame data is received for 2 consecutive seconds.
	//
	// 4. The average frame rate over the period specified by `ali_max_no_frame_timeout` drops below `ali_low_frame_rate_threshold`. This condition applies even if there is no stream disconnection or stuttering. If you set `ali_max_no_frame_timeout`, the timeout for Condition 3 is also updated to this value.
	//
	// 5. If `block_all_jitter` is set to `1`, Conditions 2, 3, and 4 do not apply.
	//
	// - `ali_max_no_frame_timeout`: an integer from 2 to 10.<br>`ali_low_frame_rate_threshold`: an integer from 1 to 200.<br>`block_all_jitter`: `0` or `1`.<br><br>
	//
	// example:
	//
	// ali_low_frame_rate_threshold=10&ali_max_no_frame_timeout=5&block_all_jitter=0
	MergeParameters *string `json:"MergeParameters,omitempty" xml:"MergeParameters,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The live stream protocol for the input streams. Valid values:
	//
	// - **rtmp*	- (Default)
	//
	// - **rtc**
	//
	// example:
	//
	// rtmp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SelectAppName    *string `json:"SelectAppName,omitempty" xml:"SelectAppName,omitempty"`
	SelectStreamName *string `json:"SelectStreamName,omitempty" xml:"SelectStreamName,omitempty"`
	// The start time of the stream merge.
	//
	// The time must be in UTC and specified in the ISO 8601 standard format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The StreamName of the output stream. For the configuration to take effect, this StreamName must match the one in the ingest URL. Wildcards (`*`) are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s AddLiveStreamMergeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamMergeRequest) GoString() string {
	return s.String()
}

func (s *AddLiveStreamMergeRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveStreamMergeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveStreamMergeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddLiveStreamMergeRequest) GetInAppName1() *string {
	return s.InAppName1
}

func (s *AddLiveStreamMergeRequest) GetInAppName2() *string {
	return s.InAppName2
}

func (s *AddLiveStreamMergeRequest) GetInStreamName1() *string {
	return s.InStreamName1
}

func (s *AddLiveStreamMergeRequest) GetInStreamName2() *string {
	return s.InStreamName2
}

func (s *AddLiveStreamMergeRequest) GetLiveMerger() *string {
	return s.LiveMerger
}

func (s *AddLiveStreamMergeRequest) GetMergeParameters() *string {
	return s.MergeParameters
}

func (s *AddLiveStreamMergeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveStreamMergeRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddLiveStreamMergeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveStreamMergeRequest) GetSelectAppName() *string {
	return s.SelectAppName
}

func (s *AddLiveStreamMergeRequest) GetSelectStreamName() *string {
	return s.SelectStreamName
}

func (s *AddLiveStreamMergeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddLiveStreamMergeRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveStreamMergeRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *AddLiveStreamMergeRequest) SetAppName(v string) *AddLiveStreamMergeRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetDomainName(v string) *AddLiveStreamMergeRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetEndTime(v string) *AddLiveStreamMergeRequest {
	s.EndTime = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetInAppName1(v string) *AddLiveStreamMergeRequest {
	s.InAppName1 = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetInAppName2(v string) *AddLiveStreamMergeRequest {
	s.InAppName2 = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetInStreamName1(v string) *AddLiveStreamMergeRequest {
	s.InStreamName1 = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetInStreamName2(v string) *AddLiveStreamMergeRequest {
	s.InStreamName2 = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetLiveMerger(v string) *AddLiveStreamMergeRequest {
	s.LiveMerger = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetMergeParameters(v string) *AddLiveStreamMergeRequest {
	s.MergeParameters = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetOwnerId(v int64) *AddLiveStreamMergeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetProtocol(v string) *AddLiveStreamMergeRequest {
	s.Protocol = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetRegionId(v string) *AddLiveStreamMergeRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetSelectAppName(v string) *AddLiveStreamMergeRequest {
	s.SelectAppName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetSelectStreamName(v string) *AddLiveStreamMergeRequest {
	s.SelectStreamName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetStartTime(v string) *AddLiveStreamMergeRequest {
	s.StartTime = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetStreamName(v string) *AddLiveStreamMergeRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetSwitchMode(v string) *AddLiveStreamMergeRequest {
	s.SwitchMode = &v
	return s
}

func (s *AddLiveStreamMergeRequest) Validate() error {
	return dara.Validate(s)
}
