// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegionId(v string) *GetBasicIpSetResponseBody
	GetAccelerateRegionId() *string
	SetAcceleratorId(v string) *GetBasicIpSetResponseBody
	GetAcceleratorId() *string
	SetBandwidth(v int64) *GetBasicIpSetResponseBody
	GetBandwidth() *int64
	SetIpAddress(v string) *GetBasicIpSetResponseBody
	GetIpAddress() *string
	SetIpSetId(v string) *GetBasicIpSetResponseBody
	GetIpSetId() *string
	SetIpVersion(v string) *GetBasicIpSetResponseBody
	GetIpVersion() *string
	SetIspType(v string) *GetBasicIpSetResponseBody
	GetIspType() *string
	SetRequestId(v string) *GetBasicIpSetResponseBody
	GetRequestId() *string
	SetState(v string) *GetBasicIpSetResponseBody
	GetState() *string
}

type GetBasicIpSetResponseBody struct {
	// The ID of the region where the basic GA instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth of the acceleration region of the basic GA instance. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The accelerated IP address.
	//
	// example:
	//
	// 118.31.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The ID of the acceleration region of the basic GA instance.
	//
	// example:
	//
	// ips-bp11ilwqjdkjeg9r7****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The Internet protocol version. Only **IPv4*	- may be returned.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) lines.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines.
	//
	// Valid values if you are allowed to use single-ISP bandwidth:
	//
	// 	- **ChinaTelecom**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaTelecom_L2**
	//
	// 	- **ChinaUnicom_L2**
	//
	// 	- **ChinaMobile_L2**
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6D2BFF54-6AF2-4679-88C4-2F2E187F16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the acceleration region of the basic GA instance. Valid values:
	//
	// 	- **init**: The acceleration region is being initialized.
	//
	// 	- **active**: The acceleration region is in the running state.
	//
	// 	- **updating**: The acceleration region is being configured.
	//
	// 	- **Deleting**: The acceleration region is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetResponseBody) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *GetBasicIpSetResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicIpSetResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *GetBasicIpSetResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetBasicIpSetResponseBody) GetIpSetId() *string {
	return s.IpSetId
}

func (s *GetBasicIpSetResponseBody) GetIpVersion() *string {
	return s.IpVersion
}

func (s *GetBasicIpSetResponseBody) GetIspType() *string {
	return s.IspType
}

func (s *GetBasicIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicIpSetResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicIpSetResponseBody) SetAccelerateRegionId(v string) *GetBasicIpSetResponseBody {
	s.AccelerateRegionId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetAcceleratorId(v string) *GetBasicIpSetResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetBandwidth(v int64) *GetBasicIpSetResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpAddress(v string) *GetBasicIpSetResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpSetId(v string) *GetBasicIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpVersion(v string) *GetBasicIpSetResponseBody {
	s.IpVersion = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIspType(v string) *GetBasicIpSetResponseBody {
	s.IspType = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetRequestId(v string) *GetBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetState(v string) *GetBasicIpSetResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
