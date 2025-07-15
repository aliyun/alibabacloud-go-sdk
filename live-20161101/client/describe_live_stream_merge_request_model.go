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
	// The name of the application that generates the output stream. You can view the application name on the [Primary/Secondary Stream Mixing Settings](https://help.aliyun.com/document_detail/606583.html) page of the ApsaraVideo Live console.
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
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The name of the output stream. You can view the stream name on the [Primary/Secondary Stream Mixing Settings](https://help.aliyun.com/document_detail/606583.html) page of the ApsaraVideo Live console.
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
