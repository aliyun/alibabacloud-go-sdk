// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSecurityGroupRuleResponseBody
	GetCode() *int32
	SetData(v []*ListSecurityGroupRuleResponseBodyData) *ListSecurityGroupRuleResponseBody
	GetData() []*ListSecurityGroupRuleResponseBodyData
	SetHttpStatusCode(v int32) *ListSecurityGroupRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSecurityGroupRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecurityGroupRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecurityGroupRuleResponseBody
	GetSuccess() *bool
}

type ListSecurityGroupRuleResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data []*ListSecurityGroupRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9EC7BDBF-3C38-5C9C-95DD-61E298CD43E8
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

func (s ListSecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSecurityGroupRuleResponseBody) GetData() []*ListSecurityGroupRuleResponseBodyData {
	return s.Data
}

func (s *ListSecurityGroupRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSecurityGroupRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityGroupRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecurityGroupRuleResponseBody) SetCode(v int32) *ListSecurityGroupRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) SetData(v []*ListSecurityGroupRuleResponseBodyData) *ListSecurityGroupRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) SetHttpStatusCode(v int32) *ListSecurityGroupRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) SetMessage(v string) *ListSecurityGroupRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) SetRequestId(v string) *ListSecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) SetSuccess(v bool) *ListSecurityGroupRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecurityGroupRuleResponseBodyData struct {
	AuthCidrs []*string `json:"AuthCidrs,omitempty" xml:"AuthCidrs,omitempty" type:"Repeated"`
	// The rule description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 81
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-12a432a1f5da423997d8880bd32c304d
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the security group authorization record.
	//
	// example:
	//
	// 21
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The port range.
	//
	// example:
	//
	// 8000/8000
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1cg6qlyjepj0y6cf2c
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ListSecurityGroupRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupRuleResponseBodyData) GetAuthCidrs() []*string {
	return s.AuthCidrs
}

func (s *ListSecurityGroupRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListSecurityGroupRuleResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListSecurityGroupRuleResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListSecurityGroupRuleResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListSecurityGroupRuleResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListSecurityGroupRuleResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListSecurityGroupRuleResponseBodyData) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListSecurityGroupRuleResponseBodyData) GetPortRange() *string {
	return s.PortRange
}

func (s *ListSecurityGroupRuleResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListSecurityGroupRuleResponseBodyData) SetAuthCidrs(v []*string) *ListSecurityGroupRuleResponseBodyData {
	s.AuthCidrs = v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetDescription(v string) *ListSecurityGroupRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetGatewayId(v int64) *ListSecurityGroupRuleResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetGatewayUniqueId(v string) *ListSecurityGroupRuleResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetGmtCreate(v string) *ListSecurityGroupRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetGmtModified(v string) *ListSecurityGroupRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetId(v string) *ListSecurityGroupRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetIpProtocol(v string) *ListSecurityGroupRuleResponseBodyData {
	s.IpProtocol = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetPortRange(v string) *ListSecurityGroupRuleResponseBodyData {
	s.PortRange = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) SetSecurityGroupId(v string) *ListSecurityGroupRuleResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListSecurityGroupRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
