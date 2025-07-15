// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveCenterTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveCenterTransferRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveCenterTransferRequest
	GetDomainName() *string
	SetDstUrl(v string) *DeleteLiveCenterTransferRequest
	GetDstUrl() *string
	SetOwnerId(v int64) *DeleteLiveCenterTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveCenterTransferRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLiveCenterTransferRequest
	GetStreamName() *string
}

type DeleteLiveCenterTransferRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name for the live stream that you want to relay. Otherwise, the configuration does not take effect. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
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
	// rtmp://push.example2.aliyundoc.com/testapp1/teststream2
	DstUrl   *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveCenterTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveCenterTransferRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveCenterTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveCenterTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveCenterTransferRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *DeleteLiveCenterTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveCenterTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveCenterTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveCenterTransferRequest) SetAppName(v string) *DeleteLiveCenterTransferRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) SetDomainName(v string) *DeleteLiveCenterTransferRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) SetDstUrl(v string) *DeleteLiveCenterTransferRequest {
	s.DstUrl = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) SetOwnerId(v int64) *DeleteLiveCenterTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) SetRegionId(v string) *DeleteLiveCenterTransferRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) SetStreamName(v string) *DeleteLiveCenterTransferRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveCenterTransferRequest) Validate() error {
	return dara.Validate(s)
}
