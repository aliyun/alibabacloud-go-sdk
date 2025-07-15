// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsControlHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlInfo(v *DescribeLiveStreamsControlHistoryResponseBodyControlInfo) *DescribeLiveStreamsControlHistoryResponseBody
	GetControlInfo() *DescribeLiveStreamsControlHistoryResponseBodyControlInfo
	SetRequestId(v string) *DescribeLiveStreamsControlHistoryResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamsControlHistoryResponseBody struct {
	// The operation records.
	ControlInfo *DescribeLiveStreamsControlHistoryResponseBodyControlInfo `json:"ControlInfo,omitempty" xml:"ControlInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9C31856F-386D-4DB3-BE79-A0AA493D702A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamsControlHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponseBody) GetControlInfo() *DescribeLiveStreamsControlHistoryResponseBodyControlInfo {
	return s.ControlInfo
}

func (s *DescribeLiveStreamsControlHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsControlHistoryResponseBody) SetControlInfo(v *DescribeLiveStreamsControlHistoryResponseBodyControlInfo) *DescribeLiveStreamsControlHistoryResponseBody {
	s.ControlInfo = v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBody) SetRequestId(v string) *DescribeLiveStreamsControlHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsControlHistoryResponseBodyControlInfo struct {
	LiveStreamControlInfo []*DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo `json:"LiveStreamControlInfo,omitempty" xml:"LiveStreamControlInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsControlHistoryResponseBodyControlInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponseBodyControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfo) GetLiveStreamControlInfo() []*DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo {
	return s.LiveStreamControlInfo
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfo) SetLiveStreamControlInfo(v []*DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) *DescribeLiveStreamsControlHistoryResponseBodyControlInfo {
	s.LiveStreamControlInfo = v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo struct {
	// The name of the operation performed.
	//
	// example:
	//
	// DescribeLiveStreamsControlHistory
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The IP address that is used by the client for live streaming.
	//
	// example:
	//
	// 10.207.XX.XX
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The time when the operation was performed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T16:36:18Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) GetAction() *string {
	return s.Action
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) SetAction(v string) *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo {
	s.Action = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) SetClientIP(v string) *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo {
	s.ClientIP = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) SetStreamName(v string) *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) SetTimeStamp(v string) *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponseBodyControlInfoLiveStreamControlInfo) Validate() error {
	return dara.Validate(s)
}
