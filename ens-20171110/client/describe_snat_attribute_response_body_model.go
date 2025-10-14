// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *DescribeSnatAttributeResponseBody
	GetCreationTime() *string
	SetDestCIDR(v string) *DescribeSnatAttributeResponseBody
	GetDestCIDR() *string
	SetEipAffinity(v bool) *DescribeSnatAttributeResponseBody
	GetEipAffinity() *bool
	SetIdleTimeout(v int32) *DescribeSnatAttributeResponseBody
	GetIdleTimeout() *int32
	SetIspAffinity(v bool) *DescribeSnatAttributeResponseBody
	GetIspAffinity() *bool
	SetNatGatewayId(v string) *DescribeSnatAttributeResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *DescribeSnatAttributeResponseBody
	GetRequestId() *string
	SetSnatEntryId(v string) *DescribeSnatAttributeResponseBody
	GetSnatEntryId() *string
	SetSnatEntryName(v string) *DescribeSnatAttributeResponseBody
	GetSnatEntryName() *string
	SetSnatIp(v string) *DescribeSnatAttributeResponseBody
	GetSnatIp() *string
	SetSnatIps(v []*DescribeSnatAttributeResponseBodySnatIps) *DescribeSnatAttributeResponseBody
	GetSnatIps() []*DescribeSnatAttributeResponseBodySnatIps
	SetSourceCIDR(v string) *DescribeSnatAttributeResponseBody
	GetSourceCIDR() *string
	SetStandbySnatIp(v string) *DescribeSnatAttributeResponseBody
	GetStandbySnatIp() *string
	SetStandbyStatus(v string) *DescribeSnatAttributeResponseBody
	GetStandbyStatus() *string
	SetStatus(v string) *DescribeSnatAttributeResponseBody
	GetStatus() *string
	SetType(v string) *DescribeSnatAttributeResponseBody
	GetType() *string
}

type DescribeSnatAttributeResponseBody struct {
	// The time when the entry was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-04-26T15:38:27Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The destination CIDR block. The rule takes effect only on requests that access the destination CIDR block.
	//
	// example:
	//
	// 101.10. XX.XX/24
	DestCIDR *string `json:"DestCIDR,omitempty" xml:"DestCIDR,omitempty"`
	// Specifies whether to enable IP affinity. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// >  After you enable IP affinity, if multiple EIPs are associated with an SNAT entry, one client uses the same EIP to for communication. If IP affinity is disabled, the client uses a random EIP for communication.
	//
	// example:
	//
	// false
	EipAffinity *bool `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 10
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Whether to enable operator affinity. Value taking:
	//
	// - false:Do not open.
	//
	// - true:Open.
	//
	// example:
	//
	// true
	IspAffinity *bool `json:"IspAffinity,omitempty" xml:"IspAffinity,omitempty"`
	// The ID of the Network Address Translation (NAT) gateway.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-5tfi6f8gds82mjmlofeym****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// example:
	//
	// test0
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// The EIP specified in the SNAT entry. Multiple EIPs are separated by commas (,).
	//
	// example:
	//
	// 120.72.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The information about the EIP specified in the SNAT entry.
	SnatIps []*DescribeSnatAttributeResponseBodySnatIps `json:"SnatIps,omitempty" xml:"SnatIps,omitempty" type:"Repeated"`
	// The source CIDR block specified in the SNAT entry.
	//
	// example:
	//
	// 10.0.XX.XX/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The secondary EIP specified in the SNAT entry. Multiple secondary EIPs are separated by commas (,).
	//
	// example:
	//
	// 101.23. XX.XX
	StandbySnatIp *string `json:"StandbySnatIp,omitempty" xml:"StandbySnatIp,omitempty"`
	// The status of the secondary EIP.
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Starting
	//
	// example:
	//
	// Stopped
	StandbyStatus *string `json:"StandbyStatus,omitempty" xml:"StandbyStatus,omitempty"`
	// The status of the SNAT entry.
	//
	// 	- Pending: The SNAT entry is being created or modified.
	//
	// 	- Available: The SNAT entry is available.
	//
	// 	- Deleting: The SNAT entry is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the NAT.
	//
	// 	- Empty: symmetric NAT.
	//
	// 	- FullCone: full cone NAT.
	//
	// example:
	//
	// FullCone
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSnatAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnatAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnatAttributeResponseBody) GetDestCIDR() *string {
	return s.DestCIDR
}

func (s *DescribeSnatAttributeResponseBody) GetEipAffinity() *bool {
	return s.EipAffinity
}

func (s *DescribeSnatAttributeResponseBody) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *DescribeSnatAttributeResponseBody) GetIspAffinity() *bool {
	return s.IspAffinity
}

func (s *DescribeSnatAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnatAttributeResponseBody) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatAttributeResponseBody) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatAttributeResponseBody) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatAttributeResponseBody) GetSnatIps() []*DescribeSnatAttributeResponseBodySnatIps {
	return s.SnatIps
}

func (s *DescribeSnatAttributeResponseBody) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatAttributeResponseBody) GetStandbySnatIp() *string {
	return s.StandbySnatIp
}

func (s *DescribeSnatAttributeResponseBody) GetStandbyStatus() *string {
	return s.StandbyStatus
}

func (s *DescribeSnatAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnatAttributeResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeSnatAttributeResponseBody) SetCreationTime(v string) *DescribeSnatAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetDestCIDR(v string) *DescribeSnatAttributeResponseBody {
	s.DestCIDR = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetEipAffinity(v bool) *DescribeSnatAttributeResponseBody {
	s.EipAffinity = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetIdleTimeout(v int32) *DescribeSnatAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetIspAffinity(v bool) *DescribeSnatAttributeResponseBody {
	s.IspAffinity = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetNatGatewayId(v string) *DescribeSnatAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetRequestId(v string) *DescribeSnatAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetSnatEntryId(v string) *DescribeSnatAttributeResponseBody {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetSnatEntryName(v string) *DescribeSnatAttributeResponseBody {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetSnatIp(v string) *DescribeSnatAttributeResponseBody {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetSnatIps(v []*DescribeSnatAttributeResponseBodySnatIps) *DescribeSnatAttributeResponseBody {
	s.SnatIps = v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetSourceCIDR(v string) *DescribeSnatAttributeResponseBody {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetStandbySnatIp(v string) *DescribeSnatAttributeResponseBody {
	s.StandbySnatIp = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetStandbyStatus(v string) *DescribeSnatAttributeResponseBody {
	s.StandbyStatus = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetStatus(v string) *DescribeSnatAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) SetType(v string) *DescribeSnatAttributeResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeSnatAttributeResponseBody) Validate() error {
	if s.SnatIps != nil {
		for _, item := range s.SnatIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnatAttributeResponseBodySnatIps struct {
	// The time when the IP address was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-04-26T15:38:27Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 203.132.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The status of the IP address.
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Starting
	//
	// 	- Releasing
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnatAttributeResponseBodySnatIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatAttributeResponseBodySnatIps) GoString() string {
	return s.String()
}

func (s *DescribeSnatAttributeResponseBodySnatIps) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnatAttributeResponseBodySnatIps) GetIp() *string {
	return s.Ip
}

func (s *DescribeSnatAttributeResponseBodySnatIps) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnatAttributeResponseBodySnatIps) SetCreationTime(v string) *DescribeSnatAttributeResponseBodySnatIps {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnatAttributeResponseBodySnatIps) SetIp(v string) *DescribeSnatAttributeResponseBodySnatIps {
	s.Ip = &v
	return s
}

func (s *DescribeSnatAttributeResponseBodySnatIps) SetStatus(v string) *DescribeSnatAttributeResponseBodySnatIps {
	s.Status = &v
	return s
}

func (s *DescribeSnatAttributeResponseBodySnatIps) Validate() error {
	return dara.Validate(s)
}
