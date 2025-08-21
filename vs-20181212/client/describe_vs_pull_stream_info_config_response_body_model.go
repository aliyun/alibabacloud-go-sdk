// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsPullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAppRecordList(v *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) *DescribeVsPullStreamInfoConfigResponseBody
	GetLiveAppRecordList() *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList
	SetRequestId(v string) *DescribeVsPullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type DescribeVsPullStreamInfoConfigResponseBody struct {
	LiveAppRecordList *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList `json:"LiveAppRecordList,omitempty" xml:"LiveAppRecordList,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) GetLiveAppRecordList() *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList {
	return s.LiveAppRecordList
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) SetLiveAppRecordList(v *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) *DescribeVsPullStreamInfoConfigResponseBody {
	s.LiveAppRecordList = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *DescribeVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList struct {
	LiveAppRecord []*DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord `json:"LiveAppRecord,omitempty" xml:"LiveAppRecord,omitempty" type:"Repeated"`
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) GetLiveAppRecord() []*DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	return s.LiveAppRecord
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) SetLiveAppRecord(v []*DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList {
	s.LiveAppRecord = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordList) Validate() error {
	return dara.Validate(s)
}

type DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord struct {
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2016-05-15T07:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// http://test
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// example:
	//
	// 2016-05-15T01:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetAppName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.AppName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetDomainName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.DomainName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetEndTime(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetSourceUrl(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.SourceUrl = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStartTime(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) SetStreamName(v string) *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord {
	s.StreamName = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponseBodyLiveAppRecordListLiveAppRecord) Validate() error {
	return dara.Validate(s)
}
