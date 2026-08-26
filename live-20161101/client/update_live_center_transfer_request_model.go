// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveCenterTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveCenterTransferRequest
	GetAppName() *string
	SetDomainName(v string) *UpdateLiveCenterTransferRequest
	GetDomainName() *string
	SetDstUrl(v string) *UpdateLiveCenterTransferRequest
	GetDstUrl() *string
	SetEndTime(v string) *UpdateLiveCenterTransferRequest
	GetEndTime() *string
	SetOwnerId(v int64) *UpdateLiveCenterTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveCenterTransferRequest
	GetRegionId() *string
	SetStartTime(v string) *UpdateLiveCenterTransferRequest
	GetStartTime() *string
	SetStreamName(v string) *UpdateLiveCenterTransferRequest
	GetStreamName() *string
	SetTransferArgs(v string) *UpdateLiveCenterTransferRequest
	GetTransferArgs() *string
}

type UpdateLiveCenterTransferRequest struct {
	// The name of the live stream application. The AppName you enter must match the AppName of the live stream to be transferred for the configuration to take effect. You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
	// The third-party live streaming address to transfer to. A maximum of one address is supported.
	//
	// >The protocol of the destination address must match the protocol of the live stream being transferred. Only RTMP and SRT protocols are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://push.example2.aliyunlive.com/testapp1/teststream2
	DstUrl *string `json:"DstUrl,omitempty" xml:"DstUrl,omitempty"`
	// The end time of the transfer. The date format follows ISO 8601 and uses UTC+0 time in the format yyyy-MM-ddTHH:mm:ssZ.
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
	// The start time of the transfer. The date format follows ISO 8601 and uses UTC+0 time in the format yyyy-MM-ddTHH:mm:ssZ.
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
	// The transfer validity period. Valid values:
	//
	// - **always**: Permanently effective.
	//
	// - **time**: Effective within a specified time range.
	//
	// >If you set this parameter to **time**, **StartTime*	- and **EndTime*	- are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// always
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s UpdateLiveCenterTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveCenterTransferRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveCenterTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveCenterTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveCenterTransferRequest) GetDstUrl() *string {
	return s.DstUrl
}

func (s *UpdateLiveCenterTransferRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateLiveCenterTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveCenterTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveCenterTransferRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateLiveCenterTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLiveCenterTransferRequest) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *UpdateLiveCenterTransferRequest) SetAppName(v string) *UpdateLiveCenterTransferRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetDomainName(v string) *UpdateLiveCenterTransferRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetDstUrl(v string) *UpdateLiveCenterTransferRequest {
	s.DstUrl = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetEndTime(v string) *UpdateLiveCenterTransferRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetOwnerId(v int64) *UpdateLiveCenterTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetRegionId(v string) *UpdateLiveCenterTransferRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetStartTime(v string) *UpdateLiveCenterTransferRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetStreamName(v string) *UpdateLiveCenterTransferRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) SetTransferArgs(v string) *UpdateLiveCenterTransferRequest {
	s.TransferArgs = &v
	return s
}

func (s *UpdateLiveCenterTransferRequest) Validate() error {
	return dara.Validate(s)
}
