// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizableSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody
	GetCode() *string
	SetData(v *ListGatewayAuthorizableSecurityGroupsResponseBodyData) *ListGatewayAuthorizableSecurityGroupsResponseBody
	GetData() *ListGatewayAuthorizableSecurityGroupsResponseBodyData
	SetMessage(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody
	GetRequestId() *string
}

type ListGatewayAuthorizableSecurityGroupsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListGatewayAuthorizableSecurityGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) GetData() *ListGatewayAuthorizableSecurityGroupsResponseBodyData {
	return s.Data
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) SetCode(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) SetData(v *ListGatewayAuthorizableSecurityGroupsResponseBodyData) *ListGatewayAuthorizableSecurityGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) SetMessage(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) SetRequestId(v string) *ListGatewayAuthorizableSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthorizableSecurityGroupsResponseBodyData struct {
	// The security groups.
	Items []*ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyData) GetItems() []*ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyData) SetItems(v []*ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) *ListGatewayAuthorizableSecurityGroupsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyData) Validate() error {
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

type ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems struct {
	// The security group name.
	//
	// example:
	//
	// 商品中心集群安全组。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1ftp5sm9os***
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The type of the security group. Valid values:
	//
	// 	- Normal: general security group
	//
	// 	- Enterprise: enterprise security group
	//
	// example:
	//
	// Normal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf61resqa9am***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) GetType() *string {
	return s.Type
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) SetName(v string) *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) SetSecurityGroupId(v string) *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems {
	s.SecurityGroupId = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) SetType(v string) *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) SetVpcId(v string) *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems {
	s.VpcId = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
