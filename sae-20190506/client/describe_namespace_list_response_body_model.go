// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeNamespaceListResponseBody
	GetCode() *string
	SetData(v []*DescribeNamespaceListResponseBodyData) *DescribeNamespaceListResponseBody
	GetData() []*DescribeNamespaceListResponseBodyData
	SetErrorCode(v string) *DescribeNamespaceListResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeNamespaceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNamespaceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNamespaceListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeNamespaceListResponseBody
	GetTraceId() *string
}

type DescribeNamespaceListResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of namespaces.
	Data []*DescribeNamespaceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// NULL
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30375C38-F4ED-4135-A0AE-5C75DC7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of namespaces was queried. Valid values:
	//
	// 	- **true**: The list was queried.
	//
	// 	- **false**: The list failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNamespaceListResponseBody) GetData() []*DescribeNamespaceListResponseBodyData {
	return s.Data
}

func (s *DescribeNamespaceListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeNamespaceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNamespaceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespaceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNamespaceListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeNamespaceListResponseBody) SetCode(v string) *DescribeNamespaceListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetData(v []*DescribeNamespaceListResponseBodyData) *DescribeNamespaceListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetErrorCode(v string) *DescribeNamespaceListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetMessage(v string) *DescribeNamespaceListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetRequestId(v string) *DescribeNamespaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetSuccess(v bool) *DescribeNamespaceListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetTraceId(v string) *DescribeNamespaceListResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespaceListResponseBodyData struct {
	// The command that was run to install the agent.
	//
	// example:
	//
	// http://edas-bj.oss-cn-beijing-internal.aliyuncs.com/test/install.sh
	AgentInstall *string `json:"AgentInstall,omitempty" xml:"AgentInstall,omitempty"`
	// This parameter is no longer valid.
	//
	// example:
	//
	// false
	Current *bool `json:"Current,omitempty" xml:"Current,omitempty"`
	// Indicates whether custom namespaces are returned. Valid values:
	//
	// 	- **true**: Custom namespaces are returned.
	//
	// 	- **false**: Custom namespaces are not returned.
	//
	// example:
	//
	// true
	Custom *bool `json:"Custom,omitempty" xml:"Custom,omitempty"`
	// Indicates whether hybrid cloud namespaces are excluded. Valid values:
	//
	// 	- **true**: Hybrid cloud namespaces are excluded.
	//
	// 	- **false**: Hybrid cloud namespaces are included.
	//
	// example:
	//
	// false
	HybridCloudEnable *bool `json:"HybridCloudEnable,omitempty" xml:"HybridCloudEnable,omitempty"`
	// The short ID of the namespace.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The region to which the namespace belongs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNamespaceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponseBodyData) GetAgentInstall() *string {
	return s.AgentInstall
}

func (s *DescribeNamespaceListResponseBodyData) GetCurrent() *bool {
	return s.Current
}

func (s *DescribeNamespaceListResponseBodyData) GetCustom() *bool {
	return s.Custom
}

func (s *DescribeNamespaceListResponseBodyData) GetHybridCloudEnable() *bool {
	return s.HybridCloudEnable
}

func (s *DescribeNamespaceListResponseBodyData) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespaceListResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespaceListResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DescribeNamespaceListResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNamespaceListResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeNamespaceListResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNamespaceListResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNamespaceListResponseBodyData) SetAgentInstall(v string) *DescribeNamespaceListResponseBodyData {
	s.AgentInstall = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetCurrent(v bool) *DescribeNamespaceListResponseBodyData {
	s.Current = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetCustom(v bool) *DescribeNamespaceListResponseBodyData {
	s.Custom = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetHybridCloudEnable(v bool) *DescribeNamespaceListResponseBodyData {
	s.HybridCloudEnable = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetNameSpaceShortId(v string) *DescribeNamespaceListResponseBodyData {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceListResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceListResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetRegionId(v string) *DescribeNamespaceListResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetSecurityGroupId(v string) *DescribeNamespaceListResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetVSwitchId(v string) *DescribeNamespaceListResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetVpcId(v string) *DescribeNamespaceListResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
