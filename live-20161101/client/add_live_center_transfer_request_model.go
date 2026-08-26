// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveCenterTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveCenterTransferRequest
	GetAppName() *string
	SetDomainName(v string) *AddLiveCenterTransferRequest
	GetDomainName() *string
	SetDstUrl(v string) *AddLiveCenterTransferRequest
	GetDstUrl() *string
	SetEndTime(v string) *AddLiveCenterTransferRequest
	GetEndTime() *string
	SetOwnerId(v int64) *AddLiveCenterTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveCenterTransferRequest
	GetRegionId() *string
	SetStartTime(v string) *AddLiveCenterTransferRequest
	GetStartTime() *string
	SetStreamName(v string) *AddLiveCenterTransferRequest
	GetStreamName() *string
	SetTransferArgs(v string) *AddLiveCenterTransferRequest
	GetTransferArgs() *string
}

type AddLiveCenterTransferRequest struct {
	// The name of the live stream application. The AppName you enter must be the same as the AppName of the live stream to be relayed for the configuration to take effect. You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
	// The third-party live streaming address for relay. You can add up to one address.
	//
	// >The protocol of the destination address must be the same as the protocol of the live stream to be relayed. Only RTMP and SRT relay addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://push.example2.aliyunlive.com/testapp1/teststream2
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of the relay. The date format follows the ISO 8601 notation and uses UTC+0 time. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// >The end time must be later than the start time.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the relay. The date format follows the ISO 8601 notation and uses UTC+0 time. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2017-12-21T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. You can view the StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The validity period of the relay. Valid values:
	//
	// - **always**: permanently effective.
	//
	// - **time**: effective within the specified time period.
	//
	// >If the value is **time**, **StartTime*	- and **EndTime*	- are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// always
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s AddLiveCenterTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveCenterTransferRequest) GoString() string {
	return s.String()
}

func (s *AddLiveCenterTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveCenterTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveCenterTransferRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *AddLiveCenterTransferRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddLiveCenterTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveCenterTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveCenterTransferRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddLiveCenterTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveCenterTransferRequest) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *AddLiveCenterTransferRequest) SetAppName(v string) *AddLiveCenterTransferRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetDomainName(v string) *AddLiveCenterTransferRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetDstUrl(v string) *AddLiveCenterTransferRequest {
	s.DstUrl = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetEndTime(v string) *AddLiveCenterTransferRequest {
	s.EndTime = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetOwnerId(v int64) *AddLiveCenterTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetRegionId(v string) *AddLiveCenterTransferRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetStartTime(v string) *AddLiveCenterTransferRequest {
	s.StartTime = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetStreamName(v string) *AddLiveCenterTransferRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveCenterTransferRequest) SetTransferArgs(v string) *AddLiveCenterTransferRequest {
	s.TransferArgs = &v
	return s
}

func (s *AddLiveCenterTransferRequest) Validate() error {
	return dara.Validate(s)
}
