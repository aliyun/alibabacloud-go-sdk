// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAvailGARequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationArea(v string) *DescribeLivePrivateLineAvailGARequest
	GetAccelerationArea() *string
	SetAppName(v string) *DescribeLivePrivateLineAvailGARequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLivePrivateLineAvailGARequest
	GetDomainName() *string
	SetIsGaInstance(v string) *DescribeLivePrivateLineAvailGARequest
	GetIsGaInstance() *string
	SetOwnerId(v int64) *DescribeLivePrivateLineAvailGARequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLivePrivateLineAvailGARequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLivePrivateLineAvailGARequest
	GetStreamName() *string
	SetVideoCenter(v string) *DescribeLivePrivateLineAvailGARequest
	GetVideoCenter() *string
}

type DescribeLivePrivateLineAvailGARequest struct {
	// The acceleration channel.
	//
	// example:
	//
	// ap-southeast-1
	AccelerationArea *string `json:"AccelerationArea,omitempty" xml:"AccelerationArea,omitempty"`
	// The name of the application.
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
	// Specifies whether to query Global Accelerator (GA) instances. Valid values:
	//
	// 	- yes: queries the status of GA instances.
	//
	// 	- no: queries the binding information between GA instances and acceleration circuits.
	//
	// This parameter is required.
	//
	// example:
	//
	// no
	IsGaInstance *string `json:"IsGaInstance,omitempty" xml:"IsGaInstance,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// testStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The live center. Valid values: cn-beijing, cn-shanghai, cn-shenzhen, cn-qingdao, ap-northeast-1, ap-southeast-5, eu-central-1, ap-southeast-1, and ap-south-1. cn-beijing indicates China (Beijing). cn-shanghai indicates China (Shanghai). cn-shenzhen indicates China (Shenzhen). cn-qingdao indicates China (Qingdao). ap-northeast-1 indicates Japan (Tokyo). ap-southeast-5 indicates Indonesia (Jakarta). eu-central-1 indicates Germany (Frankfurt). ap-southeast-1 indicates Singapore.
	//
	// example:
	//
	// cn-shanghai
	VideoCenter *string `json:"VideoCenter,omitempty" xml:"VideoCenter,omitempty"`
}

func (s DescribeLivePrivateLineAvailGARequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAvailGARequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAvailGARequest) GetAccelerationArea() *string {
	return s.AccelerationArea
}

func (s *DescribeLivePrivateLineAvailGARequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLivePrivateLineAvailGARequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePrivateLineAvailGARequest) GetIsGaInstance() *string {
	return s.IsGaInstance
}

func (s *DescribeLivePrivateLineAvailGARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePrivateLineAvailGARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePrivateLineAvailGARequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLivePrivateLineAvailGARequest) GetVideoCenter() *string {
	return s.VideoCenter
}

func (s *DescribeLivePrivateLineAvailGARequest) SetAccelerationArea(v string) *DescribeLivePrivateLineAvailGARequest {
	s.AccelerationArea = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetAppName(v string) *DescribeLivePrivateLineAvailGARequest {
	s.AppName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetDomainName(v string) *DescribeLivePrivateLineAvailGARequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetIsGaInstance(v string) *DescribeLivePrivateLineAvailGARequest {
	s.IsGaInstance = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetOwnerId(v int64) *DescribeLivePrivateLineAvailGARequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetRegionId(v string) *DescribeLivePrivateLineAvailGARequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetStreamName(v string) *DescribeLivePrivateLineAvailGARequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) SetVideoCenter(v string) *DescribeLivePrivateLineAvailGARequest {
	s.VideoCenter = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGARequest) Validate() error {
	return dara.Validate(s)
}
