// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyGateways(v []*ListEgressProxyGatewaysResponseBodyEgressProxyGateways) *ListEgressProxyGatewaysResponseBody
	GetEgressProxyGateways() []*ListEgressProxyGatewaysResponseBodyEgressProxyGateways
	SetMaxResults(v int32) *ListEgressProxyGatewaysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyGatewaysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEgressProxyGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEgressProxyGatewaysResponseBody
	GetTotalCount() *int32
}

type ListEgressProxyGatewaysResponseBody struct {
	EgressProxyGateways []*ListEgressProxyGatewaysResponseBodyEgressProxyGateways `json:"EgressProxyGateways,omitempty" xml:"EgressProxyGateways,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I4FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 01A05AB6-3B8A-561C-8CC6-D061498BDEBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEgressProxyGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysResponseBody) GetEgressProxyGateways() []*ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	return s.EgressProxyGateways
}

func (s *ListEgressProxyGatewaysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyGatewaysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEgressProxyGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEgressProxyGatewaysResponseBody) SetEgressProxyGateways(v []*ListEgressProxyGatewaysResponseBodyEgressProxyGateways) *ListEgressProxyGatewaysResponseBody {
	s.EgressProxyGateways = v
	return s
}

func (s *ListEgressProxyGatewaysResponseBody) SetMaxResults(v int32) *ListEgressProxyGatewaysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBody) SetNextToken(v string) *ListEgressProxyGatewaysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBody) SetRequestId(v string) *ListEgressProxyGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBody) SetTotalCount(v int32) *ListEgressProxyGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBody) Validate() error {
	if s.EgressProxyGateways != nil {
		for _, item := range s.EgressProxyGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyGatewaysResponseBodyEgressProxyGateways struct {
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	DefaultEgressProxyIpPoolId *string `json:"DefaultEgressProxyIpPoolId,omitempty" xml:"DefaultEgressProxyIpPoolId,omitempty"`
	// example:
	//
	// created by epg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// epc-uf609a6a8fdf9f4****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyGatewayName *string `json:"EgressProxyGatewayName,omitempty" xml:"EgressProxyGatewayName,omitempty"`
	// example:
	//
	// rg-acfmypqurk****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Provisioning
	Status *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-m5e03az9xwn6siqft****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEgressProxyGatewaysResponseBodyEgressProxyGateways) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetDefaultEgressProxyIpPoolId() *string {
	return s.DefaultEgressProxyIpPoolId
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetDescription() *string {
	return s.Description
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetEgressProxyGatewayName() *string {
	return s.EgressProxyGatewayName
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetTags() []*ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags {
	return s.Tags
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetCreateTime(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.CreateTime = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetDefaultEgressProxyIpPoolId(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.DefaultEgressProxyIpPoolId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetDescription(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.Description = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetEgressProxyConfigurationId(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetEgressProxyGatewayId(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetEgressProxyGatewayName(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.EgressProxyGatewayName = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetResourceGroupId(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetStatus(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.Status = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetTags(v []*ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.Tags = v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) SetVpcId(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGateways {
	s.VpcId = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGateways) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) SetKey(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags {
	s.Key = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) SetValue(v string) *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags {
	s.Value = &v
	return s
}

func (s *ListEgressProxyGatewaysResponseBodyEgressProxyGatewaysTags) Validate() error {
	return dara.Validate(s)
}
