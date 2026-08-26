// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamSnapshotInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetEndTime() *string
	SetLimit(v int32) *DescribeLiveStreamSnapshotInfoRequest
	GetLimit() *int32
	SetOrder(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLiveStreamSnapshotInfoRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamSnapshotInfoRequest
	GetStreamName() *string
}

type DescribeLiveStreamSnapshotInfoRequest struct {
	// The name of the application to which the stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain of the streamer.
	//
	// > - When you specify DomainName, make sure that the domain name is a live streaming domain name and that the user who calls this operation has the permissions to operate on the specified domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The interval between EndTime and StartTime cannot exceed **1*	- day. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return in a single call. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc*	- (default): ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the stream. Make sure that the specified stream name exists within the query time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamSnapshotInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamSnapshotInfoRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetAppName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetDomainName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetEndTime(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetLimit(v int32) *DescribeLiveStreamSnapshotInfoRequest {
	s.Limit = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetOrder(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.Order = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetOwnerId(v int64) *DescribeLiveStreamSnapshotInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetSecurityToken(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetStartTime(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) SetStreamName(v string) *DescribeLiveStreamSnapshotInfoRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoRequest) Validate() error {
	return dara.Validate(s)
}
