// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMergeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamMergeRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamMergeRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamMergeRequest
	GetOwnerId() *int64
	SetProtocol(v string) *DescribeLiveStreamMergeRequest
	GetProtocol() *string
	SetRegionId(v string) *DescribeLiveStreamMergeRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveStreamMergeRequest
	GetStreamName() *string
}

type DescribeLiveStreamMergeRequest struct {
	// Merged output App name. You can view this on the [Primary/Backup Stream Merge Configuration](https://help.aliyun.com/document_detail/606583.html) page.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Streaming protocol. Valid values:
	//
	// - **rtmp*	- (default)
	//
	// - **rtc**
	//
	// example:
	//
	// rtmp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Merged output Stream name. You can view this on the [Primary/Backup Stream Merge Configuration](https://help.aliyun.com/document_detail/606583.html) page.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamMergeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMergeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMergeRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamMergeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamMergeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamMergeRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLiveStreamMergeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamMergeRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamMergeRequest) SetAppName(v string) *DescribeLiveStreamMergeRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) SetDomainName(v string) *DescribeLiveStreamMergeRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) SetOwnerId(v int64) *DescribeLiveStreamMergeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) SetProtocol(v string) *DescribeLiveStreamMergeRequest {
	s.Protocol = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) SetRegionId(v string) *DescribeLiveStreamMergeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) SetStreamName(v string) *DescribeLiveStreamMergeRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamMergeRequest) Validate() error {
	return dara.Validate(s)
}
