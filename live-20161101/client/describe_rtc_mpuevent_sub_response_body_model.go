// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcMPUEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRtcMPUEventSubResponseBody
	GetRequestId() *string
	SetSubInfo(v *DescribeRtcMPUEventSubResponseBodySubInfo) *DescribeRtcMPUEventSubResponseBody
	GetSubInfo() *DescribeRtcMPUEventSubResponseBodySubInfo
}

type DescribeRtcMPUEventSubResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the subscription.
	SubInfo *DescribeRtcMPUEventSubResponseBodySubInfo `json:"SubInfo,omitempty" xml:"SubInfo,omitempty" type:"Struct"`
}

func (s DescribeRtcMPUEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcMPUEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcMPUEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcMPUEventSubResponseBody) GetSubInfo() *DescribeRtcMPUEventSubResponseBodySubInfo {
	return s.SubInfo
}

func (s *DescribeRtcMPUEventSubResponseBody) SetRequestId(v string) *DescribeRtcMPUEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBody) SetSubInfo(v *DescribeRtcMPUEventSubResponseBodySubInfo) *DescribeRtcMPUEventSubResponseBody {
	s.SubInfo = v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBody) Validate() error {
	if s.SubInfo != nil {
		if err := s.SubInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcMPUEventSubResponseBodySubInfo struct {
	// The application ID. You can specify only one application ID.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://testcallback***.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the channel to which mixed-stream relay event callbacks are sent. Multiple channel IDs are separated by commas (,). If this parameter is not returned, mixed-stream relay event callbacks are sent to all channels.
	//
	// example:
	//
	// yourCh1,yourCh2
	ChannelIds *string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty"`
	// The time when the event callback was fired. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-09 18:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// Sub-******9799B2C4500******
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s DescribeRtcMPUEventSubResponseBodySubInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcMPUEventSubResponseBodySubInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) GetChannelIds() *string {
	return s.ChannelIds
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) GetSubId() *string {
	return s.SubId
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) SetAppId(v string) *DescribeRtcMPUEventSubResponseBodySubInfo {
	s.AppId = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) SetCallbackUrl(v string) *DescribeRtcMPUEventSubResponseBodySubInfo {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) SetChannelIds(v string) *DescribeRtcMPUEventSubResponseBodySubInfo {
	s.ChannelIds = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) SetCreateTime(v string) *DescribeRtcMPUEventSubResponseBodySubInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) SetSubId(v string) *DescribeRtcMPUEventSubResponseBodySubInfo {
	s.SubId = &v
	return s
}

func (s *DescribeRtcMPUEventSubResponseBodySubInfo) Validate() error {
	return dara.Validate(s)
}
