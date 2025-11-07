// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatTopNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTopNOpen(v bool) *GetNatTopNResponseBody
	GetIsTopNOpen() *bool
	SetNatGatewayTopN(v []*GetNatTopNResponseBodyNatGatewayTopN) *GetNatTopNResponseBody
	GetNatGatewayTopN() []*GetNatTopNResponseBodyNatGatewayTopN
	SetRequestId(v string) *GetNatTopNResponseBody
	GetRequestId() *string
}

type GetNatTopNResponseBody struct {
	// Indicates whether Network Intelligence Service (NIS) is activated. The NatGatewayTopN parameter returns an empty array when NIS is not activated.
	//
	// 	- **true**: activated
	//
	// 	- **false**: not activated
	//
	// example:
	//
	// true
	IsTopNOpen *bool `json:"IsTopNOpen,omitempty" xml:"IsTopNOpen,omitempty"`
	// An array of statistics about real-time SNAT performance ranking.
	NatGatewayTopN []*GetNatTopNResponseBodyNatGatewayTopN `json:"NatGatewayTopN,omitempty" xml:"NatGatewayTopN,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 77C512B5-12f3-f892-BD94-88A98271C1A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNatTopNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNatTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponseBody) GetIsTopNOpen() *bool {
	return s.IsTopNOpen
}

func (s *GetNatTopNResponseBody) GetNatGatewayTopN() []*GetNatTopNResponseBodyNatGatewayTopN {
	return s.NatGatewayTopN
}

func (s *GetNatTopNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNatTopNResponseBody) SetIsTopNOpen(v bool) *GetNatTopNResponseBody {
	s.IsTopNOpen = &v
	return s
}

func (s *GetNatTopNResponseBody) SetNatGatewayTopN(v []*GetNatTopNResponseBodyNatGatewayTopN) *GetNatTopNResponseBody {
	s.NatGatewayTopN = v
	return s
}

func (s *GetNatTopNResponseBody) SetRequestId(v string) *GetNatTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNatTopNResponseBody) Validate() error {
	if s.NatGatewayTopN != nil {
		for _, item := range s.NatGatewayTopN {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNatTopNResponseBodyNatGatewayTopN struct {
	// The number of concurrent connections. Unit: connections.
	//
	// example:
	//
	// 8
	ActiveSessionCount *float32 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// The inbound data transfer. Unit: bit/s.
	//
	// example:
	//
	// 100
	InBps *float32 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// This field is reserved and not in use.
	//
	// example:
	//
	// 10
	InFlowPerMinute *float32 `json:"InFlowPerMinute,omitempty" xml:"InFlowPerMinute,omitempty"`
	// The inbound packet forwarding rate. Unit: packets per second.
	//
	// example:
	//
	// 10
	InPps *float32 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.156.101
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The new connection creation rate. Unit: connections per second.
	//
	// example:
	//
	// 2
	NewSessionPerSecond *float32 `json:"NewSessionPerSecond,omitempty" xml:"NewSessionPerSecond,omitempty"`
	// The outbound data transfer. Unit: bit/s.
	//
	// example:
	//
	// 200
	OutBps *float32 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// This field is reserved and not in use.
	//
	// example:
	//
	// 10
	OutFlowPerMinute *float32 `json:"OutFlowPerMinute,omitempty" xml:"OutFlowPerMinute,omitempty"`
	// The outbound packet forwarding rate. Unit: packets per second.
	//
	// example:
	//
	// 20
	OutPps *float32 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
}

func (s GetNatTopNResponseBodyNatGatewayTopN) String() string {
	return dara.Prettify(s)
}

func (s GetNatTopNResponseBodyNatGatewayTopN) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetActiveSessionCount() *float32 {
	return s.ActiveSessionCount
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetInBps() *float32 {
	return s.InBps
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetInFlowPerMinute() *float32 {
	return s.InFlowPerMinute
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetInPps() *float32 {
	return s.InPps
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetIp() *string {
	return s.Ip
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetNewSessionPerSecond() *float32 {
	return s.NewSessionPerSecond
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetOutBps() *float32 {
	return s.OutBps
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetOutFlowPerMinute() *float32 {
	return s.OutFlowPerMinute
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) GetOutPps() *float32 {
	return s.OutPps
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetActiveSessionCount(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInBps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InBps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInFlowPerMinute(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InFlowPerMinute = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInPps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InPps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetIp(v string) *GetNatTopNResponseBodyNatGatewayTopN {
	s.Ip = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetNewSessionPerSecond(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.NewSessionPerSecond = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutBps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutBps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutFlowPerMinute(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutFlowPerMinute = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutPps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutPps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) Validate() error {
	return dara.Validate(s)
}
