// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetAttachedGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachedGateways(v []*ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) *ListPolicySetAttachedGatewaysResponseBody
	GetAttachedGateways() []*ListPolicySetAttachedGatewaysResponseBodyAttachedGateways
	SetMaxResults(v int32) *ListPolicySetAttachedGatewaysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPolicySetAttachedGatewaysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPolicySetAttachedGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPolicySetAttachedGatewaysResponseBody
	GetTotalCount() *int32
}

type ListPolicySetAttachedGatewaysResponseBody struct {
	AttachedGateways []*ListPolicySetAttachedGatewaysResponseBodyAttachedGateways `json:"AttachedGateways,omitempty" xml:"AttachedGateways,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicySetAttachedGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetAttachedGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicySetAttachedGatewaysResponseBody) GetAttachedGateways() []*ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	return s.AttachedGateways
}

func (s *ListPolicySetAttachedGatewaysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPolicySetAttachedGatewaysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicySetAttachedGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicySetAttachedGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPolicySetAttachedGatewaysResponseBody) SetAttachedGateways(v []*ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) *ListPolicySetAttachedGatewaysResponseBody {
	s.AttachedGateways = v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBody) SetMaxResults(v int32) *ListPolicySetAttachedGatewaysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBody) SetNextToken(v string) *ListPolicySetAttachedGatewaysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBody) SetRequestId(v string) *ListPolicySetAttachedGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBody) SetTotalCount(v int32) *ListPolicySetAttachedGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBody) Validate() error {
	if s.AttachedGateways != nil {
		for _, item := range s.AttachedGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicySetAttachedGatewaysResponseBodyAttachedGateways struct {
	// example:
	//
	// default-policy-set
	AttachedPolicySetName *string `json:"AttachedPolicySetName,omitempty" xml:"AttachedPolicySetName,omitempty"`
	// example:
	//
	// 2026-05-08T06:19:17Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"EnforcementMode,omitempty" xml:"EnforcementMode,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:gateway/my-gateway
	GatewayArn  *string `json:"GatewayArn,omitempty" xml:"GatewayArn,omitempty"`
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
}

func (s ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GoString() string {
	return s.String()
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GetAttachedPolicySetName() *string {
	return s.AttachedPolicySetName
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GetGatewayArn() *string {
	return s.GatewayArn
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) SetAttachedPolicySetName(v string) *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	s.AttachedPolicySetName = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) SetAttachedTime(v string) *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	s.AttachedTime = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) SetEnforcementMode(v string) *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	s.EnforcementMode = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) SetGatewayArn(v string) *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	s.GatewayArn = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) SetGatewayType(v string) *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways {
	s.GatewayType = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysResponseBodyAttachedGateways) Validate() error {
	return dara.Validate(s)
}
