// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizedSecurityGroupRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody
	GetCode() *string
	SetData(v *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) *ListGatewayAuthorizedSecurityGroupRulesResponseBody
	GetData() *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData
	SetMessage(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody
	GetRequestId() *string
}

type ListGatewayAuthorizedSecurityGroupRulesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) GetData() *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData {
	return s.Data
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) SetCode(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) SetData(v *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) *ListGatewayAuthorizedSecurityGroupRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) SetMessage(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) SetRequestId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthorizedSecurityGroupRulesResponseBodyData struct {
	// The security group rules.
	Items []*ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) GetItems() []*ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) SetItems(v []*ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems struct {
	// The list of authorized CIDR blocks.
	AuthCidrs []*string `json:"authCidrs,omitempty" xml:"authCidrs,omitempty" type:"Repeated"`
	// The rule description.
	//
	// example:
	//
	// 商品中心预发网关授权安全组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The protocol. Valid values:
	//
	// 	- TCP
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"ipProtocol,omitempty" xml:"ipProtocol,omitempty"`
	// The port range.
	//
	// example:
	//
	// 8080/8089
	PortRange *string `json:"portRange,omitempty" xml:"portRange,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1ftp5sm9os***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The security group name.
	//
	// example:
	//
	// 商品中心集群安全组
	SecurityGroupName *string `json:"securityGroupName,omitempty" xml:"securityGroupName,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// gsgr-cqadis5lhtgmv***
	SecurityGroupRuleId *string `json:"securityGroupRuleId,omitempty" xml:"securityGroupRuleId,omitempty"`
	// The ID of the source security group.
	//
	// example:
	//
	// sg-bp19akuepfe***
	SourceSecurityGroupId *string `json:"sourceSecurityGroupId,omitempty" xml:"sourceSecurityGroupId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp1g63b5q2q29***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetAuthCidrs() []*string {
	return s.AuthCidrs
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetPortRange() *string {
	return s.PortRange
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetSourceSecurityGroupId() *string {
	return s.SourceSecurityGroupId
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetAuthCidrs(v []*string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.AuthCidrs = v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetDescription(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetIpProtocol(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.IpProtocol = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetPortRange(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.PortRange = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetSecurityGroupId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.SecurityGroupId = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetSecurityGroupName(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.SecurityGroupName = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetSecurityGroupRuleId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetSourceSecurityGroupId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.SourceSecurityGroupId = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) SetVpcId(v string) *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems {
	s.VpcId = &v
	return s
}

func (s *ListGatewayAuthorizedSecurityGroupRulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
