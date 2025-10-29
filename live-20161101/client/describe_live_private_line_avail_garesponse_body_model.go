// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAvailGAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePrivateLineAvailGAs(v *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) *DescribeLivePrivateLineAvailGAResponseBody
	GetLivePrivateLineAvailGAs() *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs
	SetRequestId(v string) *DescribeLivePrivateLineAvailGAResponseBody
	GetRequestId() *string
}

type DescribeLivePrivateLineAvailGAResponseBody struct {
	// The GA instance configuration details.
	LivePrivateLineAvailGAs *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs `json:"LivePrivateLineAvailGAs,omitempty" xml:"LivePrivateLineAvailGAs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C4865B85-664B-19D3-BB16-C62FB83C8226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLivePrivateLineAvailGAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAvailGAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAvailGAResponseBody) GetLivePrivateLineAvailGAs() *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs {
	return s.LivePrivateLineAvailGAs
}

func (s *DescribeLivePrivateLineAvailGAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePrivateLineAvailGAResponseBody) SetLivePrivateLineAvailGAs(v *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) *DescribeLivePrivateLineAvailGAResponseBody {
	s.LivePrivateLineAvailGAs = v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBody) SetRequestId(v string) *DescribeLivePrivateLineAvailGAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBody) Validate() error {
	if s.LivePrivateLineAvailGAs != nil {
		if err := s.LivePrivateLineAvailGAs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs struct {
	LivePrivateLineAvailGA []*DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA `json:"LivePrivateLineAvailGA,omitempty" xml:"LivePrivateLineAvailGA,omitempty" type:"Repeated"`
}

func (s DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) GetLivePrivateLineAvailGA() []*DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	return s.LivePrivateLineAvailGA
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) SetLivePrivateLineAvailGA(v []*DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs {
	s.LivePrivateLineAvailGA = v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAs) Validate() error {
	if s.LivePrivateLineAvailGA != nil {
		for _, item := range s.LivePrivateLineAvailGA {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA struct {
	// The acceleration channel.
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
	// example:
	//
	// play
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// live
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Indicates whether the GA instance is bound to an acceleration circuit. Valid values:
	//
	// 	- yes
	//
	// 	- no
	//
	// example:
	//
	// yes
	BindingStatus *string `json:"BindingStatus,omitempty" xml:"BindingStatus,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The accelerated IP address.
	//
	// example:
	//
	// 127.0.0.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1iovsdpf01ym9su****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the GA instance. Valid values:
	//
	// 	- active: The GA instance is available.
	//
	// 	- inactive: The GA instance is unavailable.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// testStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The user ID (UID).
	//
	// example:
	//
	// 1833220971116****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The live center.
	//
	// example:
	//
	// cn-shanghai
	VideoCenter *string `json:"VideoCenter,omitempty" xml:"VideoCenter,omitempty"`
}

func (s DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetAccelerationArea() *string {
	return s.AccelerationArea
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetBindingStatus() *string {
	return s.BindingStatus
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetIP() *string {
	return s.IP
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetStatus() *string {
	return s.Status
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetUid() *string {
	return s.Uid
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) GetVideoCenter() *string {
	return s.VideoCenter
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetAccelerationArea(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.AccelerationArea = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetAccelerationType(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.AccelerationType = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetAppName(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.AppName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetBindingStatus(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.BindingStatus = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetDomainName(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetIP(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.IP = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetInstanceId(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.InstanceId = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetStatus(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.Status = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetStreamName(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetUid(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.Uid = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) SetVideoCenter(v string) *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA {
	s.VideoCenter = &v
	return s
}

func (s *DescribeLivePrivateLineAvailGAResponseBodyLivePrivateLineAvailGAsLivePrivateLineAvailGA) Validate() error {
	return dara.Validate(s)
}
