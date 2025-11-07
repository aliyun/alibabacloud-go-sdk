// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAnalyzeNetworkPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkReachableAnalysisId(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetNetworkReachableAnalysisId() *string
	SetProtocol(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetProtocol() *string
	SetRequestId(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetRequestId() *string
	SetSourceId(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetSourceId() *string
	SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetSourceIpAddress() *string
	SetSourcePort(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetSourcePort() *string
	SetSourceType(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetSourceType() *string
	SetTargetId(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetTargetId() *string
	SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetTargetIpAddress() *string
	SetTargetPort(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetTargetPort() *string
	SetTargetType(v string) *CreateAndAnalyzeNetworkPathResponseBody
	GetTargetType() *string
}

type CreateAndAnalyzeNetworkPathResponseBody struct {
	// The ID of the task for analyzing network reachability that you initiated.
	//
	// example:
	//
	// nra-dfe9e53d2b524568****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the source resource.
	//
	// example:
	//
	// i-uf62y8khhbkbdrp6****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 0
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source resource.
	//
	// example:
	//
	// ecs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the destination resource.
	//
	// example:
	//
	// i-m5eactvw7wtpktv5****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 172.50.XX.XX
	TargetIpAddress *string `json:"TargetIpAddress,omitempty" xml:"TargetIpAddress,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the destination resource.
	//
	// example:
	//
	// ecs
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAndAnalyzeNetworkPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetNetworkReachableAnalysisId() *string {
	return s.NetworkReachableAnalysisId
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetSourcePort() *string {
	return s.SourcePort
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetTargetIpAddress() *string {
	return s.TargetIpAddress
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetTargetPort() *string {
	return s.TargetPort
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetNetworkReachableAnalysisId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetProtocol(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.Protocol = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetRequestId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourcePort(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourcePort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceType(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetPort(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetPort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetType(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) Validate() error {
	return dara.Validate(s)
}
