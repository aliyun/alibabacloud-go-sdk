// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransitRouterFlowTopNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTransitRouterFlowTopNResponseBody
	GetRequestId() *string
	SetTransitRouterFlowTopN(v []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) *GetTransitRouterFlowTopNResponseBody
	GetTransitRouterFlowTopN() []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN
}

type GetTransitRouterFlowTopNResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ranking result of inter-region traffic data.
	TransitRouterFlowTopN []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN `json:"TransitRouterFlowTopN,omitempty" xml:"TransitRouterFlowTopN,omitempty" type:"Repeated"`
}

func (s GetTransitRouterFlowTopNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTransitRouterFlowTopNResponseBody) GetTransitRouterFlowTopN() []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	return s.TransitRouterFlowTopN
}

func (s *GetTransitRouterFlowTopNResponseBody) SetRequestId(v string) *GetTransitRouterFlowTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBody) SetTransitRouterFlowTopN(v []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) *GetTransitRouterFlowTopNResponseBody {
	s.TransitRouterFlowTopN = v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBody) Validate() error {
	if s.TransitRouterFlowTopN != nil {
		for _, item := range s.TransitRouterFlowTopN {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN struct {
	// The account ID.
	//
	// example:
	//
	// 118639953821xxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the CEN bandwidth plan.
	//
	// example:
	//
	// cenbwp-ia8kw1zjv4hyal****
	BandwithPackageId *string `json:"BandwithPackageId,omitempty" xml:"BandwithPackageId,omitempty"`
	// The total volume of traffic in the specified time range.
	//
	// example:
	//
	// 188
	Bytes *float64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The end of the time range that you queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-31T06:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 47.216.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 53470
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The remote region where the **remote IP address*	- resides.
	//
	// example:
	//
	// ap-southeast-1
	OtherRegion *string `json:"OtherRegion,omitempty" xml:"OtherRegion,omitempty"`
	// The total number of packets in the specified time range.
	//
	// example:
	//
	// 88
	Packets *float64 `json:"Packets,omitempty" xml:"Packets,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The beginning of the time range that you queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-31T05:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 1.8.XX.XX
	ThisIp *string `json:"ThisIp,omitempty" xml:"ThisIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	ThisPort *string `json:"ThisPort,omitempty" xml:"ThisPort,omitempty"`
	// The local region where the **local IP address*	- resides.
	//
	// example:
	//
	// cn-shanghai
	ThisRegion *string `json:"ThisRegion,omitempty" xml:"ThisRegion,omitempty"`
}

func (s GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) String() string {
	return dara.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetBandwithPackageId() *string {
	return s.BandwithPackageId
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetBytes() *float64 {
	return s.Bytes
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetCenId() *string {
	return s.CenId
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetOtherRegion() *string {
	return s.OtherRegion
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetPackets() *float64 {
	return s.Packets
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetProtocol() *string {
	return s.Protocol
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetThisIp() *string {
	return s.ThisIp
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetThisPort() *string {
	return s.ThisPort
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GetThisRegion() *string {
	return s.ThisRegion
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetAccountId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.AccountId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetBytes(v float64) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Bytes = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetCenId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetEndTime(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherIp(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherPort(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherRegion(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetPackets(v float64) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Packets = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetProtocol(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetStartTime(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.StartTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisIp(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisPort(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisRegion(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) Validate() error {
	return dara.Validate(s)
}
