// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcGatewayEndpointAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetCreationTime() *string
	SetEndpointDescription(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetEndpointDescription() *string
	SetEndpointId(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetEndpointName() *string
	SetEndpointStatus(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetEndpointStatus() *string
	SetPolicyDocument(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetPolicyDocument() *string
	SetRequestId(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetResourceGroupId() *string
	SetRouteTables(v []*string) *GetVpcGatewayEndpointAttributeResponseBody
	GetRouteTables() []*string
	SetServiceName(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetServiceName() *string
	SetTags(v []*GetVpcGatewayEndpointAttributeResponseBodyTags) *GetVpcGatewayEndpointAttributeResponseBody
	GetTags() []*GetVpcGatewayEndpointAttributeResponseBodyTags
	SetVpcId(v string) *GetVpcGatewayEndpointAttributeResponseBody
	GetVpcId() *string
}

type GetVpcGatewayEndpointAttributeResponseBody struct {
	// The time when the endpoint was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-27T01:58:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the gateway endpoint.
	//
	// example:
	//
	// test
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The ID of the gateway endpoint.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the gateway endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The status of the gateway endpoint. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Associating**
	//
	// 	- **Dissociating**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Created
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The access policy for the cloud service.
	//
	// example:
	//
	// {"Version" : "1",   "Statement" : [ {     "Effect" : "Allow",     "Resource" : [ "*" ],     "Action" : [ "*" ],     "Principal" : [ "*" ]   } ] }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1122D0F-7B3B-5445-BB19-17F27F97FE1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the gateway endpoint belongs.
	//
	// example:
	//
	// rg-acfmxvfvazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the route table associated with the gateway endpoint.
	RouteTables []*string `json:"RouteTables,omitempty" xml:"RouteTables,omitempty" type:"Repeated"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tag list.
	Tags []*GetVpcGatewayEndpointAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the gateway endpoint belongs.
	//
	// example:
	//
	// vpc-bp1nh86rugg01zol0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetVpcGatewayEndpointAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcGatewayEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetRouteTables() []*string {
	return s.RouteTables
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetTags() []*GetVpcGatewayEndpointAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetCreationTime(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetEndpointDescription(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.EndpointDescription = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetEndpointId(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.EndpointId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetEndpointName(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.EndpointName = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetEndpointStatus(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.EndpointStatus = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetPolicyDocument(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.PolicyDocument = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetRequestId(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetResourceGroupId(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetRouteTables(v []*string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.RouteTables = v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetServiceName(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetTags(v []*GetVpcGatewayEndpointAttributeResponseBodyTags) *GetVpcGatewayEndpointAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) SetVpcId(v string) *GetVpcGatewayEndpointAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBody) Validate() error {
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

type GetVpcGatewayEndpointAttributeResponseBodyTags struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetVpcGatewayEndpointAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetVpcGatewayEndpointAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetVpcGatewayEndpointAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetVpcGatewayEndpointAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetVpcGatewayEndpointAttributeResponseBodyTags) SetKey(v string) *GetVpcGatewayEndpointAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBodyTags) SetValue(v string) *GetVpcGatewayEndpointAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetVpcGatewayEndpointAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
