// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullStreamConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAppRecordList(v *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) *DescribeLivePullStreamConfigResponseBody
	GetLiveAppRecordList() *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList
	SetRequestId(v string) *DescribeLivePullStreamConfigResponseBody
	GetRequestId() *string
}

type DescribeLivePullStreamConfigResponseBody struct {
	LiveAppRecordList *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLivePullStreamConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponseBody) GetLiveAppRecordList() *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList {
	return s.LiveAppRecordList
}

func (s *DescribeLivePullStreamConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePullStreamConfigResponseBody) SetLiveAppRecordList(v *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) *DescribeLivePullStreamConfigResponseBody {
	s.LiveAppRecordList = v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBody) SetRequestId(v string) *DescribeLivePullStreamConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBody) Validate() error {
	if s.LiveAppRecordList != nil {
		if err := s.LiveAppRecordList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLivePullStreamConfigResponseBodyLiveAppRecordList struct {
	LiveAppRecord []*DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" type:"Repeated"`
}

func (s DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) GetLiveAppRecord() []*DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	return s.LiveAppRecord
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) SetLiveAppRecord(v []*DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordList) Validate() error {
	if s.LiveAppRecord != nil {
		for _, item := range s.LiveAppRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceUrl   *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	SourceUsing *string `json:"SourceUsing,omitempty" xml:"SourceUsing,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName  *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetSourceUsing() *string {
	return s.SourceUsing
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetSourceUrl(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.SourceUrl = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetSourceUsing(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.SourceUsing = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponseBodyLiveAppRecordListLiveAppRecord) Validate() error {
	return dara.Validate(s)
}
