// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteSecurityGroupRuleResponseBody
	GetCode() *int32
	SetData(v *DeleteSecurityGroupRuleResponseBodyData) *DeleteSecurityGroupRuleResponseBody
	GetData() *DeleteSecurityGroupRuleResponseBodyData
	SetHttpStatusCode(v int32) *DeleteSecurityGroupRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityGroupRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecurityGroupRuleResponseBody
	GetSuccess() *bool
}

type DeleteSecurityGroupRuleResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteSecurityGroupRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 43E50CB7-258E-5AFF-9B93-ECC19928C699
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteSecurityGroupRuleResponseBody) GetData() *DeleteSecurityGroupRuleResponseBodyData {
	return s.Data
}

func (s *DeleteSecurityGroupRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityGroupRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityGroupRuleResponseBody) SetCode(v int32) *DeleteSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) SetData(v *DeleteSecurityGroupRuleResponseBodyData) *DeleteSecurityGroupRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityGroupRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) SetMessage(v string) *DeleteSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) SetRequestId(v string) *DeleteSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) SetSuccess(v bool) *DeleteSecurityGroupRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityGroupRuleResponseBodyData struct {
	// The description.
	//
	// example:
	//
	// auto-description1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 103
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-01-11T14:12:55.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transport layer protocol. The value of this parameter is case-insensitive. Valid values:
	//
	// 	- icmp
	//
	// 	- gre
	//
	// 	- tcp
	//
	// 	- udp
	//
	// 	- all: All protocols are supported.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The range of ports for the transport layer protocol in the destination security group. Valid values:
	//
	// 	- When the IpProtocol parameter is set to tcp or udp, the port number range is 1 to 65535. The start port number and the end port number are separated by a forward slash (/). Example: 1/200.
	//
	// 	- If the IpProtocol parameter is set to icmp, the port number range is -1/-1, which indicates all ports.
	//
	// 	- If the IpProtocol parameter is set to gre, the port number range is -1/-1, which indicates all ports.
	//
	// 	- If the IpProtocol parameter is set to all, the port number range is -1/-1, which indicates all ports.
	//
	// example:
	//
	// 8443/8443
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-uf6hgwe067prhg68agfa
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetPortRange() *string {
	return s.PortRange
}

func (s *DeleteSecurityGroupRuleResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetDescription(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetGatewayId(v int64) *DeleteSecurityGroupRuleResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetGatewayUniqueId(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetGmtCreate(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetGmtModified(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetId(v int64) *DeleteSecurityGroupRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetIpProtocol(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.IpProtocol = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetPortRange(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.PortRange = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) SetSecurityGroupId(v string) *DeleteSecurityGroupRuleResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteSecurityGroupRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
