// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCenterTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveCenterTransferRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveCenterTransferRequest
	GetDomainName() *string
	SetDstUrl(v string) *DescribeLiveCenterTransferRequest
	GetDstUrl() *string
	SetOwnerId(v int64) *DescribeLiveCenterTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveCenterTransferRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveCenterTransferRequest
	GetStreamName() *string
}

type DescribeLiveCenterTransferRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name for the live stream that you want to relay. Otherwise, the configuration does not take effect. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The third-party URL to which the live stream is relayed.
	//
	// example:
	//
	// rtmp://push.example2.aliyunlive.com/testapp1/teststream2
	DstUrl   *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveCenterTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCenterTransferRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCenterTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveCenterTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveCenterTransferRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DescribeLiveCenterTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveCenterTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveCenterTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveCenterTransferRequest) SetAppName(v string) *DescribeLiveCenterTransferRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) SetDomainName(v string) *DescribeLiveCenterTransferRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) SetDstUrl(v string) *DescribeLiveCenterTransferRequest {
	s.DstUrl = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) SetOwnerId(v int64) *DescribeLiveCenterTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) SetRegionId(v string) *DescribeLiveCenterTransferRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) SetStreamName(v string) *DescribeLiveCenterTransferRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveCenterTransferRequest) Validate() error {
	return dara.Validate(s)
}
