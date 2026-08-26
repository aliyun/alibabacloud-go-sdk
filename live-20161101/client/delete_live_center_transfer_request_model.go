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
	// The name of the application. The value must be the same as that of the live stream that you want to relay. View AppName on the [Stream Management](t2019924.xdita#).
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
	DstUrl  *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. View StreamName on the [Stream Management](t2019924.xdita#) page.
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
