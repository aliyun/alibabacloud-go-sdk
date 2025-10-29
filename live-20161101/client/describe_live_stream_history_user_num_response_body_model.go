// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamHistoryUserNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamUserNumInfos(v *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) *DescribeLiveStreamHistoryUserNumResponseBody
	GetLiveStreamUserNumInfos() *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos
	SetRequestId(v string) *DescribeLiveStreamHistoryUserNumResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamHistoryUserNumResponseBody struct {
	// The number of historical online users for the live stream.
	LiveStreamUserNumInfos *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos `json:"LiveStreamUserNumInfos,omitempty" xml:"LiveStreamUserNumInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F5FF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamHistoryUserNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponseBody) GetLiveStreamUserNumInfos() *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos {
	return s.LiveStreamUserNumInfos
}

func (s *DescribeLiveStreamHistoryUserNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamHistoryUserNumResponseBody) SetLiveStreamUserNumInfos(v *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) *DescribeLiveStreamHistoryUserNumResponseBody {
	s.LiveStreamUserNumInfos = v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseBody) SetRequestId(v string) *DescribeLiveStreamHistoryUserNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseBody) Validate() error {
	if s.LiveStreamUserNumInfos != nil {
		if err := s.LiveStreamUserNumInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos struct {
	LiveStreamUserNumInfo []*DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo `json:"LiveStreamUserNumInfo,omitempty" xml:"LiveStreamUserNumInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) GetLiveStreamUserNumInfo() []*DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo {
	return s.LiveStreamUserNumInfo
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) SetLiveStreamUserNumInfo(v []*DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos {
	s.LiveStreamUserNumInfo = v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfos) Validate() error {
	if s.LiveStreamUserNumInfo != nil {
		for _, item := range s.LiveStreamUserNumInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo struct {
	// The time when the stream started. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-10-20T06:20:00Z
	StreamTime *string `json:"StreamTime,omitempty" xml:"StreamTime,omitempty"`
	// The number of users at the current point in time.
	//
	// example:
	//
	// 1
	UserNum *string `json:"UserNum,omitempty" xml:"UserNum,omitempty"`
}

func (s DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) GetStreamTime() *string {
	return s.StreamTime
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) GetUserNum() *string {
	return s.UserNum
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) SetStreamTime(v string) *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo {
	s.StreamTime = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) SetUserNum(v string) *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo {
	s.UserNum = &v
	return s
}

func (s *DescribeLiveStreamHistoryUserNumResponseBodyLiveStreamUserNumInfosLiveStreamUserNumInfo) Validate() error {
	return dara.Validate(s)
}
