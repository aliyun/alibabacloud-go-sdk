// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVbrFlowTopNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVbrFlowTopNResponseBody
	GetRequestId() *string
	SetVirtualBorderRouterFlowlogTopN(v []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) *GetVbrFlowTopNResponseBody
	GetVirtualBorderRouterFlowlogTopN() []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN
}

type GetVbrFlowTopNResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ranking result of hybrid cloud traffic data.
	VirtualBorderRouterFlowlogTopN []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN `json:"VirtualBorderRouterFlowlogTopN,omitempty" xml:"VirtualBorderRouterFlowlogTopN,omitempty" type:"Repeated"`
}

func (s GetVbrFlowTopNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVbrFlowTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVbrFlowTopNResponseBody) GetVirtualBorderRouterFlowlogTopN() []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	return s.VirtualBorderRouterFlowlogTopN
}

func (s *GetVbrFlowTopNResponseBody) SetRequestId(v string) *GetVbrFlowTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBody) SetVirtualBorderRouterFlowlogTopN(v []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) *GetVbrFlowTopNResponseBody {
	s.VirtualBorderRouterFlowlogTopN = v
	return s
}

func (s *GetVbrFlowTopNResponseBody) Validate() error {
	if s.VirtualBorderRouterFlowlogTopN != nil {
		for _, item := range s.VirtualBorderRouterFlowlogTopN {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN struct {
	// The account ID.
	//
	// example:
	//
	// 156237031628****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The CEN connection ID.
	//
	// example:
	//
	// tr-attach-u6v1j3jre0fe9h****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The total volume of traffic in the specified time range.
	//
	// example:
	//
	// 108
	Bytes *float64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 120.24.X.X
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The local region where the local IP address resides.
	//
	// example:
	//
	// cn-shanghai
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 222.85.X.X
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 10965
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The total number of packets in the specified time range.
	//
	// example:
	//
	// 66
	Packets *float64 `json:"Packets,omitempty" xml:"Packets,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the VBR that is associated with the Express Connect circuit.
	//
	// example:
	//
	// vbr-k1atj46citwuek42j****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) String() string {
	return dara.Prettify(s)
}

func (s GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetAccountId() *string {
	return s.AccountId
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetBytes() *float64 {
	return s.Bytes
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetCloudIp() *string {
	return s.CloudIp
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetCloudPort() *string {
	return s.CloudPort
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetCloudRegion() *string {
	return s.CloudRegion
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetOtherIp() *string {
	return s.OtherIp
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetOtherPort() *string {
	return s.OtherPort
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetPackets() *float64 {
	return s.Packets
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetProtocol() *string {
	return s.Protocol
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GetVirtualBorderRouterId() *string {
	return s.VirtualBorderRouterId
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetAccountId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.AccountId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetAttachmentId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetBytes(v float64) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Bytes = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudIp(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudPort(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudRegion(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudRegion = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetOtherIp(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetOtherPort(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetPackets(v float64) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Packets = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetProtocol(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.VirtualBorderRouterId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) Validate() error {
	return dara.Validate(s)
}
