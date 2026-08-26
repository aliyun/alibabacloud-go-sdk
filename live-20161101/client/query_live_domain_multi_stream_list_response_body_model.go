// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveDomainMultiStreamListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineStreams(v []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) *QueryLiveDomainMultiStreamListResponseBody
	GetOnlineStreams() []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreams
	SetPageNumber(v int32) *QueryLiveDomainMultiStreamListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryLiveDomainMultiStreamListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryLiveDomainMultiStreamListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryLiveDomainMultiStreamListResponseBody
	GetTotalCount() *int32
}

type QueryLiveDomainMultiStreamListResponseBody struct {
	// The number of online records.
	OnlineStreams []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreams `json:"OnlineStreams,omitempty" xml:"OnlineStreams,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryLiveDomainMultiStreamListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListResponseBody) GetOnlineStreams() []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	return s.OnlineStreams
}

func (s *QueryLiveDomainMultiStreamListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryLiveDomainMultiStreamListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryLiveDomainMultiStreamListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLiveDomainMultiStreamListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryLiveDomainMultiStreamListResponseBody) SetOnlineStreams(v []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) *QueryLiveDomainMultiStreamListResponseBody {
	s.OnlineStreams = v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBody) SetPageNumber(v int32) *QueryLiveDomainMultiStreamListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBody) SetPageSize(v int32) *QueryLiveDomainMultiStreamListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBody) SetRequestId(v string) *QueryLiveDomainMultiStreamListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBody) SetTotalCount(v int32) *QueryLiveDomainMultiStreamListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBody) Validate() error {
	if s.OnlineStreams != nil {
		for _, item := range s.OnlineStreams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryLiveDomainMultiStreamListResponseBodyOnlineStreams struct {
	// The application name.
	//
	// example:
	//
	// apptest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The stream switching records.
	ChangeLogs []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Repeated"`
	// The streaming domain of the streamer.
	//
	// example:
	//
	// play.***.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The feature switch. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	OptimalMode *string `json:"OptimalMode,omitempty" xml:"OptimalMode,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The list of all candidate streams.
	UpstreamList []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList `json:"UpstreamList,omitempty" xml:"UpstreamList,omitempty" type:"Repeated"`
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetAppName() *string {
	return s.AppName
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetChangeLogs() []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	return s.ChangeLogs
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetDomain() *string {
	return s.Domain
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetOptimalMode() *string {
	return s.OptimalMode
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetStreamName() *string {
	return s.StreamName
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) GetUpstreamList() []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList {
	return s.UpstreamList
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetAppName(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.AppName = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetChangeLogs(v []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.ChangeLogs = v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetDomain(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.Domain = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetOptimalMode(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.OptimalMode = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetStreamName(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.StreamName = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) SetUpstreamList(v []*QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams {
	s.UpstreamList = v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreams) Validate() error {
	if s.ChangeLogs != nil {
		for _, item := range s.ChangeLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpstreamList != nil {
		for _, item := range s.UpstreamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs struct {
	// The reason for stream switching.
	//
	// 	- merge cut manually: The user manually switched the stream.
	//
	// 	- master stream no data: The primary stream has no data.
	//
	// 	- master stream low quality: The primary stream quality degraded.
	//
	// example:
	//
	// merge cut manually
	ChangeReason *string `json:"ChangeReason,omitempty" xml:"ChangeReason,omitempty"`
	// The stream switching time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format (UTC+0).
	//
	// example:
	//
	// 2024-11-13T09:20:47Z
	ChangeTime *string `json:"ChangeTime,omitempty" xml:"ChangeTime,omitempty"`
	// The stream that is actually used after the switch.
	//
	// example:
	//
	// rtmp://118.178.168.35:1936/wwMultitest/pull.livetest2.aliyunlive.com_wwMultitest428_AliRewrite_2?vhost=pull.livetest2.aliyunlive.com&live_rtmp_test=on
	MasterUpstream *string `json:"MasterUpstream,omitempty" xml:"MasterUpstream,omitempty"`
	// The IP address used after the stream switch.
	//
	// example:
	//
	// 1.1.1.1
	UpstreamIp *string `json:"UpstreamIp,omitempty" xml:"UpstreamIp,omitempty"`
	// The stream identifier after the switch.
	//
	// example:
	//
	// ***test_AliRewrite_2
	UpstreamSequence *string `json:"UpstreamSequence,omitempty" xml:"UpstreamSequence,omitempty"`
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GetChangeReason() *string {
	return s.ChangeReason
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GetChangeTime() *string {
	return s.ChangeTime
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GetMasterUpstream() *string {
	return s.MasterUpstream
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GetUpstreamIp() *string {
	return s.UpstreamIp
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) GetUpstreamSequence() *string {
	return s.UpstreamSequence
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) SetChangeReason(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	s.ChangeReason = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) SetChangeTime(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	s.ChangeTime = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) SetMasterUpstream(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	s.MasterUpstream = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) SetUpstreamIp(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	s.UpstreamIp = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) SetUpstreamSequence(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs {
	s.UpstreamSequence = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsChangeLogs) Validate() error {
	return dara.Validate(s)
}

type QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList struct {
	// The primary/secondary flag.
	//
	// > Indicates which stream is currently being used for merged distribution.
	//
	// example:
	//
	// false
	MasterFlag *bool `json:"MasterFlag,omitempty" xml:"MasterFlag,omitempty"`
	// The IP address of the ingest client.
	//
	// example:
	//
	// 1.1.1.1
	UpstreamIp *string `json:"UpstreamIp,omitempty" xml:"UpstreamIp,omitempty"`
	// The unique identifier of the ingest stream.
	//
	// example:
	//
	// ***test_Alirewrite1
	UpstreamSequence *string `json:"UpstreamSequence,omitempty" xml:"UpstreamSequence,omitempty"`
	// The stream ingest time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format (UTC+0).
	//
	// example:
	//
	// 2024-11-13T09:20:47Z
	UpstreamTime *string `json:"UpstreamTime,omitempty" xml:"UpstreamTime,omitempty"`
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) GoString() string {
	return s.String()
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) GetMasterFlag() *bool {
	return s.MasterFlag
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) GetUpstreamIp() *string {
	return s.UpstreamIp
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) GetUpstreamSequence() *string {
	return s.UpstreamSequence
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) GetUpstreamTime() *string {
	return s.UpstreamTime
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) SetMasterFlag(v bool) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList {
	s.MasterFlag = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) SetUpstreamIp(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList {
	s.UpstreamIp = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) SetUpstreamSequence(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList {
	s.UpstreamSequence = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) SetUpstreamTime(v string) *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList {
	s.UpstreamTime = &v
	return s
}

func (s *QueryLiveDomainMultiStreamListResponseBodyOnlineStreamsUpstreamList) Validate() error {
	return dara.Validate(s)
}
