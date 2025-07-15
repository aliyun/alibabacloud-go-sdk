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
	SetOwnerId(v int64) *AddLiveStreamMergeRequest
	GetOwnerId() *int64
	SetProtocol(v string) *AddLiveStreamMergeRequest
	GetProtocol() *string
	SetRegionId(v string) *AddLiveStreamMergeRequest
	GetRegionId() *string
	SetStartTime(v string) *AddLiveStreamMergeRequest
	GetStartTime() *string
	SetStreamName(v string) *AddLiveStreamMergeRequest
	GetStreamName() *string
}

type AddLiveStreamMergeRequest struct {
	// The name of the application that generates the output stream. The value must be the same as the application name in the ingest URL of the output stream. Otherwise, the configuration does not take effect. You cannot set the value to an asterisk (\\*).
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
	// The end time of the stream mixing.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The interval between the start time and the end time must be within 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-29T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the application that generates the input primary stream. The value must be the same as the application name that is specified in the ingest URL of the primary stream. Otherwise, the configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// app1
	InAppName1 *string `json:"InAppName1,omitempty" xml:"InAppName1,omitempty"`
	// The name of the application that generates the input secondary stream. The value must be the same as the application name that is specified in the ingest URL of the secondary stream. Otherwise, the configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// app2
	InAppName2 *string `json:"InAppName2,omitempty" xml:"InAppName2,omitempty"`
	// The name of the input primary stream. The value must be the same as the stream name that is specified in the ingest URL of the primary stream. Otherwise, the configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// InStream1
	InStreamName1 *string `json:"InStreamName1,omitempty" xml:"InStreamName1,omitempty"`
	// The name of the input secondary stream. The value must be the same as the stream name that is specified in the ingest URL of the secondary stream. Otherwise, the configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// stream2
	InStreamName2 *string `json:"InStreamName2,omitempty" xml:"InStreamName2,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The streaming protocol. Valid values:
	//
	// 	- **rtmp**: This is the default value.
	//
	// 	- **rtc**
	//
	// example:
	//
	// rtmp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the stream mixing.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the output stream. The value must be the same as the stream name in the ingest URL of the output stream. Otherwise, the configuration does not take effect. You cannot set the value to an asterisk (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
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

func (s *AddLiveStreamMergeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveStreamMergeRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddLiveStreamMergeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveStreamMergeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddLiveStreamMergeRequest) GetStreamName() *string {
	return s.StreamName
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

func (s *AddLiveStreamMergeRequest) SetStartTime(v string) *AddLiveStreamMergeRequest {
	s.StartTime = &v
	return s
}

func (s *AddLiveStreamMergeRequest) SetStreamName(v string) *AddLiveStreamMergeRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveStreamMergeRequest) Validate() error {
	return dara.Validate(s)
}
