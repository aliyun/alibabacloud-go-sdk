// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePrivateLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationArea(v string) *CreateLivePrivateLineRequest
	GetAccelerationArea() *string
	SetAccelerationType(v string) *CreateLivePrivateLineRequest
	GetAccelerationType() *string
	SetAppName(v string) *CreateLivePrivateLineRequest
	GetAppName() *string
	SetDomainName(v string) *CreateLivePrivateLineRequest
	GetDomainName() *string
	SetInstanceId(v string) *CreateLivePrivateLineRequest
	GetInstanceId() *string
	SetMaxBandwidth(v string) *CreateLivePrivateLineRequest
	GetMaxBandwidth() *string
	SetOwnerId(v int64) *CreateLivePrivateLineRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLivePrivateLineRequest
	GetRegionId() *string
	SetReuse(v string) *CreateLivePrivateLineRequest
	GetReuse() *string
	SetStreamName(v string) *CreateLivePrivateLineRequest
	GetStreamName() *string
	SetVideoCenter(v string) *CreateLivePrivateLineRequest
	GetVideoCenter() *string
}

type CreateLivePrivateLineRequest struct {
	// The acceleration channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	AccelerationArea *string `json:"AccelerationArea,omitempty" xml:"AccelerationArea,omitempty"`
	// The acceleration type. Valid values:
	//
	// 	- play: streaming acceleration
	//
	// 	- publish: stream ingest acceleration
	//
	// This parameter is required.
	//
	// example:
	//
	// play
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// live
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The acceleration channel that you want to reuse. This parameter is required if Reuse is set to yes.
	//
	// example:
	//
	// ga-bp1iovsdpf01ym9su****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The accelerated bandwidth. Unit: Mbit/s. This parameter is required if Reuse is set to no.
	//
	// example:
	//
	// 200
	MaxBandwidth *string `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to reuse an existing acceleration channel. Valid values:
	//
	// 	- yes: reuses an existing acceleration channel.
	//
	// 	- no: creates a new acceleration channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// no
	Reuse *string `json:"Reuse,omitempty" xml:"Reuse,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// testStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The live center. Valid values: cn-beijing, cn-shanghai, cn-shenzhen, cn-qingdao, ap-northeast-1, ap-southeast-5, eu-central-1, and ap-southeast-1, which indicate China (Beijing), China (Shanghai), China (Shenzhen), China (Qingdao), Japan (Tokyo), Indonesia (Jakarta), Germany (Frankfurt), and Singapore, respectively.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	VideoCenter *string `json:"VideoCenter,omitempty" xml:"VideoCenter,omitempty"`
}

func (s CreateLivePrivateLineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePrivateLineRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePrivateLineRequest) GetAccelerationArea() *string {
	return s.AccelerationArea
}

func (s *CreateLivePrivateLineRequest) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *CreateLivePrivateLineRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateLivePrivateLineRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateLivePrivateLineRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLivePrivateLineRequest) GetMaxBandwidth() *string {
	return s.MaxBandwidth
}

func (s *CreateLivePrivateLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLivePrivateLineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLivePrivateLineRequest) GetReuse() *string {
	return s.Reuse
}

func (s *CreateLivePrivateLineRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *CreateLivePrivateLineRequest) GetVideoCenter() *string {
	return s.VideoCenter
}

func (s *CreateLivePrivateLineRequest) SetAccelerationArea(v string) *CreateLivePrivateLineRequest {
	s.AccelerationArea = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetAccelerationType(v string) *CreateLivePrivateLineRequest {
	s.AccelerationType = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetAppName(v string) *CreateLivePrivateLineRequest {
	s.AppName = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetDomainName(v string) *CreateLivePrivateLineRequest {
	s.DomainName = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetInstanceId(v string) *CreateLivePrivateLineRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetMaxBandwidth(v string) *CreateLivePrivateLineRequest {
	s.MaxBandwidth = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetOwnerId(v int64) *CreateLivePrivateLineRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetRegionId(v string) *CreateLivePrivateLineRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetReuse(v string) *CreateLivePrivateLineRequest {
	s.Reuse = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetStreamName(v string) *CreateLivePrivateLineRequest {
	s.StreamName = &v
	return s
}

func (s *CreateLivePrivateLineRequest) SetVideoCenter(v string) *CreateLivePrivateLineRequest {
	s.VideoCenter = &v
	return s
}

func (s *CreateLivePrivateLineRequest) Validate() error {
	return dara.Validate(s)
}
