// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMergeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamMergeList(v *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) *DescribeLiveStreamMergeResponseBody
	GetLiveStreamMergeList() *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList
	SetRequestId(v string) *DescribeLiveStreamMergeResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamMergeResponseBody struct {
	LiveStreamMergeList *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList `json:"LiveStreamMergeList,omitempty" xml:"LiveStreamMergeList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamMergeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMergeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMergeResponseBody) GetLiveStreamMergeList() *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList {
	return s.LiveStreamMergeList
}

func (s *DescribeLiveStreamMergeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamMergeResponseBody) SetLiveStreamMergeList(v *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) *DescribeLiveStreamMergeResponseBody {
	s.LiveStreamMergeList = v
	return s
}

func (s *DescribeLiveStreamMergeResponseBody) SetRequestId(v string) *DescribeLiveStreamMergeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBody) Validate() error {
	if s.LiveStreamMergeList != nil {
		if err := s.LiveStreamMergeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamMergeResponseBodyLiveStreamMergeList struct {
	LiveStreamMerge []*DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge `json:"LiveStreamMerge,omitempty" xml:"LiveStreamMerge,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) GetLiveStreamMerge() []*DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	return s.LiveStreamMerge
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) SetLiveStreamMerge(v []*DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList {
	s.LiveStreamMerge = v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeList) Validate() error {
	if s.LiveStreamMerge != nil {
		for _, item := range s.LiveStreamMerge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge struct {
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppUsing          *string `json:"AppUsing,omitempty" xml:"AppUsing,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtraInAppStreams *string `json:"ExtraInAppStreams,omitempty" xml:"ExtraInAppStreams,omitempty"`
	InAppName1        *string `json:"InAppName1,omitempty" xml:"InAppName1,omitempty"`
	InAppName2        *string `json:"InAppName2,omitempty" xml:"InAppName2,omitempty"`
	InStreamName1     *string `json:"InStreamName1,omitempty" xml:"InStreamName1,omitempty"`
	InStreamName2     *string `json:"InStreamName2,omitempty" xml:"InStreamName2,omitempty"`
	LiveMerger        *string `json:"LiveMerger,omitempty" xml:"LiveMerger,omitempty"`
	MergeParameters   *string `json:"MergeParameters,omitempty" xml:"MergeParameters,omitempty"`
	Protocol          *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamUsing       *string `json:"StreamUsing,omitempty" xml:"StreamUsing,omitempty"`
}

func (s DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetAppUsing() *string {
	return s.AppUsing
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetExtraInAppStreams() *string {
	return s.ExtraInAppStreams
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetInAppName1() *string {
	return s.InAppName1
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetInAppName2() *string {
	return s.InAppName2
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetInStreamName1() *string {
	return s.InStreamName1
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetInStreamName2() *string {
	return s.InStreamName2
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetLiveMerger() *string {
	return s.LiveMerger
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetMergeParameters() *string {
	return s.MergeParameters
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) GetStreamUsing() *string {
	return s.StreamUsing
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetAppName(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetAppUsing(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.AppUsing = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetDomainName(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetEndTime(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetExtraInAppStreams(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.ExtraInAppStreams = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetInAppName1(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.InAppName1 = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetInAppName2(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.InAppName2 = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetInStreamName1(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.InStreamName1 = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetInStreamName2(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.InStreamName2 = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetLiveMerger(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.LiveMerger = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetMergeParameters(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.MergeParameters = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetProtocol(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.Protocol = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetStartTime(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetStreamName(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) SetStreamUsing(v string) *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge {
	s.StreamUsing = &v
	return s
}

func (s *DescribeLiveStreamMergeResponseBodyLiveStreamMergeListLiveStreamMerge) Validate() error {
	return dara.Validate(s)
}
