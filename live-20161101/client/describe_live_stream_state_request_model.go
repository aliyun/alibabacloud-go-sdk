// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamStateRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamStateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamStateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamStateRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveStreamStateRequest
	GetStreamName() *string
}

type DescribeLiveStreamStateRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain or ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamStateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamStateRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamStateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamStateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamStateRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamStateRequest) SetAppName(v string) *DescribeLiveStreamStateRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamStateRequest) SetDomainName(v string) *DescribeLiveStreamStateRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamStateRequest) SetOwnerId(v int64) *DescribeLiveStreamStateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamStateRequest) SetRegionId(v string) *DescribeLiveStreamStateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamStateRequest) SetStreamName(v string) *DescribeLiveStreamStateRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamStateRequest) Validate() error {
	return dara.Validate(s)
}
