// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUidOnlineStreamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeUidOnlineStreamsRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeUidOnlineStreamsRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeUidOnlineStreamsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUidOnlineStreamsRequest
	GetRegionId() *string
}

type DescribeUidOnlineStreamsRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUidOnlineStreamsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUidOnlineStreamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUidOnlineStreamsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeUidOnlineStreamsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUidOnlineStreamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUidOnlineStreamsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUidOnlineStreamsRequest) SetAppName(v string) *DescribeUidOnlineStreamsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeUidOnlineStreamsRequest) SetDomainName(v string) *DescribeUidOnlineStreamsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUidOnlineStreamsRequest) SetOwnerId(v int64) *DescribeUidOnlineStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUidOnlineStreamsRequest) SetRegionId(v string) *DescribeUidOnlineStreamsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUidOnlineStreamsRequest) Validate() error {
	return dara.Validate(s)
}
