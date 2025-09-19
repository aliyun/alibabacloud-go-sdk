// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotNodeResponseBody
	GetCode() *string
	SetHoneypotNode(v *GetHoneypotNodeResponseBodyHoneypotNode) *GetHoneypotNodeResponseBody
	GetHoneypotNode() *GetHoneypotNodeResponseBodyHoneypotNode
	SetHttpStatusCode(v int32) *GetHoneypotNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotNodeResponseBody
	GetSuccess() *bool
}

type GetHoneypotNodeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the management node.
	HoneypotNode *GetHoneypotNodeResponseBodyHoneypotNode `json:"HoneypotNode,omitempty" xml:"HoneypotNode,omitempty" type:"Struct"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0A453658-070B-5554-B46C-867425BE4FD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneypotNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotNodeResponseBody) GetHoneypotNode() *GetHoneypotNodeResponseBodyHoneypotNode {
	return s.HoneypotNode
}

func (s *GetHoneypotNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotNodeResponseBody) SetCode(v string) *GetHoneypotNodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotNodeResponseBody) SetHoneypotNode(v *GetHoneypotNodeResponseBodyHoneypotNode) *GetHoneypotNodeResponseBody {
	s.HoneypotNode = v
	return s
}

func (s *GetHoneypotNodeResponseBody) SetHttpStatusCode(v int32) *GetHoneypotNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotNodeResponseBody) SetMessage(v string) *GetHoneypotNodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotNodeResponseBody) SetRequestId(v string) *GetHoneypotNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotNodeResponseBody) SetSuccess(v bool) *GetHoneypotNodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotNodeResponseBodyHoneypotNode struct {
	// Indicates whether a honeypot is allowed to access the Internet. Valid values:
	//
	// 	- **true**: The honeypot is allowed to access the Internet.
	//
	// 	- **false**: The honeypot is not allowed to access the Internet.
	//
	// example:
	//
	// true
	AllowHoneypotAccessInternet *bool `json:"AllowHoneypotAccessInternet,omitempty" xml:"AllowHoneypotAccessInternet,omitempty"`
	// The time when the management node was created.
	//
	// example:
	//
	// 2022-12-02 17:13:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-2vccskxjunf1ag6w****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The maximum number of honeypots that can be deployed to the management node.
	//
	// example:
	//
	// 10
	HoneypotTotalCount *int32 `json:"HoneypotTotalCount,omitempty" xml:"HoneypotTotalCount,omitempty"`
	// The number of honeypots that are deployed to the management node.
	//
	// example:
	//
	// 5
	HoneypotUsedCount *int32 `json:"HoneypotUsedCount,omitempty" xml:"HoneypotUsedCount,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// a7409a58-bc60-41af-9d36-080d58ae****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The IP address of the management node.
	//
	// example:
	//
	// 101.37.XX.XX
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// gmmc
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The maximum number of probes that can be deployed for the management node.
	//
	// example:
	//
	// 20
	ProbeTotalCount *int32 `json:"ProbeTotalCount,omitempty" xml:"ProbeTotalCount,omitempty"`
	// The number of probes that are deployed for the management node.
	//
	// example:
	//
	// 15
	ProbeUsedCount *int32 `json:"ProbeUsedCount,omitempty" xml:"ProbeUsedCount,omitempty"`
	// An array consisting of the CIDR blocks that are allowed to access the management node.
	SecurityGroupProbeIpList []*string `json:"SecurityGroupProbeIpList,omitempty" xml:"SecurityGroupProbeIpList,omitempty" type:"Repeated"`
	// The status of the management node. Valid values:
	//
	// 	- **0**: preparing
	//
	// 	- **1**: normal
	//
	// 	- **2**: abnormal
	//
	// 	- **4**: starting
	//
	// 	- **5**: upgrading
	//
	// example:
	//
	// 2
	TotalStatus *int32 `json:"TotalStatus,omitempty" xml:"TotalStatus,omitempty"`
	// Indicates whether the management node can be upgraded. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	UpgradeAvailable *bool `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
}

func (s GetHoneypotNodeResponseBodyHoneypotNode) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeResponseBodyHoneypotNode) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetAllowHoneypotAccessInternet() *bool {
	return s.AllowHoneypotAccessInternet
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetHoneypotTotalCount() *int32 {
	return s.HoneypotTotalCount
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetHoneypotUsedCount() *int32 {
	return s.HoneypotUsedCount
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetNodeIp() *string {
	return s.NodeIp
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetProbeTotalCount() *int32 {
	return s.ProbeTotalCount
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetProbeUsedCount() *int32 {
	return s.ProbeUsedCount
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetSecurityGroupProbeIpList() []*string {
	return s.SecurityGroupProbeIpList
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetTotalStatus() *int32 {
	return s.TotalStatus
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) GetUpgradeAvailable() *bool {
	return s.UpgradeAvailable
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetAllowHoneypotAccessInternet(v bool) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.AllowHoneypotAccessInternet = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetCreateTime(v string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.CreateTime = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetEcsInstanceId(v string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.EcsInstanceId = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetHoneypotTotalCount(v int32) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.HoneypotTotalCount = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetHoneypotUsedCount(v int32) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.HoneypotUsedCount = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetNodeId(v string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.NodeId = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetNodeIp(v string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.NodeIp = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetNodeName(v string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.NodeName = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetProbeTotalCount(v int32) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.ProbeTotalCount = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetProbeUsedCount(v int32) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.ProbeUsedCount = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetSecurityGroupProbeIpList(v []*string) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.SecurityGroupProbeIpList = v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetTotalStatus(v int32) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.TotalStatus = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) SetUpgradeAvailable(v bool) *GetHoneypotNodeResponseBodyHoneypotNode {
	s.UpgradeAvailable = &v
	return s
}

func (s *GetHoneypotNodeResponseBodyHoneypotNode) Validate() error {
	return dara.Validate(s)
}
