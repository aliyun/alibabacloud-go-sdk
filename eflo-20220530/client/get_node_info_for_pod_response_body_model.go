// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInfoForPodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetNodeInfoForPodResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetNodeInfoForPodResponseBody
	GetCode() *int32
	SetContent(v *GetNodeInfoForPodResponseBodyContent) *GetNodeInfoForPodResponseBody
	GetContent() *GetNodeInfoForPodResponseBodyContent
	SetMessage(v string) *GetNodeInfoForPodResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNodeInfoForPodResponseBody
	GetRequestId() *string
}

type GetNodeInfoForPodResponseBody struct {
	// The details about the failed permission verification.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetNodeInfoForPodResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// You don\\"t have the permission of this operation, action=eflo:GetNodeInfoForPod, arn=acs:eflo:cn-wulanchabu:1111156667137893:networkinterface/*, resourceGroup=null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeInfoForPodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInfoForPodResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetNodeInfoForPodResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetNodeInfoForPodResponseBody) GetContent() *GetNodeInfoForPodResponseBodyContent {
	return s.Content
}

func (s *GetNodeInfoForPodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNodeInfoForPodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeInfoForPodResponseBody) SetAccessDeniedDetail(v string) *GetNodeInfoForPodResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetCode(v int32) *GetNodeInfoForPodResponseBody {
	s.Code = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetContent(v *GetNodeInfoForPodResponseBodyContent) *GetNodeInfoForPodResponseBody {
	s.Content = v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetMessage(v string) *GetNodeInfoForPodResponseBody {
	s.Message = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) SetRequestId(v string) *GetNodeInfoForPodResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeInfoForPodResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun Hdeni Network Interface IPV6 address Quota
	//
	// example:
	//
	// 10
	HdeniIpv6SipQuota *int32 `json:"HdeniIpv6SipQuota,omitempty" xml:"HdeniIpv6SipQuota,omitempty"`
	// Lingjun Gaomi network interface controller quota
	//
	// example:
	//
	// 10
	HdeniQuota *int32 `json:"HdeniQuota,omitempty" xml:"HdeniQuota,omitempty"`
	// Lingjun Hdeni Network Interface IPV4 address Quota
	//
	// example:
	//
	// 10
	HdeniSipQuota *int32 `json:"HdeniSipQuota,omitempty" xml:"HdeniSipQuota,omitempty"`
	// Lingjun Elastic Network Interface IPV6 address Quota
	//
	// example:
	//
	// 10
	LeniIpv6SipQuota *int32 `json:"LeniIpv6SipQuota,omitempty" xml:"LeniIpv6SipQuota,omitempty"`
	// Lingjun Elastic Network Interface quota, including system type
	//
	// example:
	//
	// 10
	LeniQuota *int32 `json:"LeniQuota,omitempty" xml:"LeniQuota,omitempty"`
	// Lingjun Elastic Network Interface Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LeniSipQuota *int32 `json:"LeniSipQuota,omitempty" xml:"LeniSipQuota,omitempty"`
	// Lingjun network interface controller Secondary Private IP Quota
	//
	// example:
	//
	// 10
	LniSipQuota *int32 `json:"LniSipQuota,omitempty" xml:"LniSipQuota,omitempty"`
	// The ID of the node for this operation.
	//
	// example:
	//
	// node-be70****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of VSwitches that can apply for IP addresses on this node
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud to which the current node belongs.
	//
	// example:
	//
	// vpc-j6ctp4n75306****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNodeInfoForPodResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInfoForPodResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetNodeInfoForPodResponseBodyContent) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetNodeInfoForPodResponseBodyContent) GetHdeniIpv6SipQuota() *int32 {
	return s.HdeniIpv6SipQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetHdeniQuota() *int32 {
	return s.HdeniQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetHdeniSipQuota() *int32 {
	return s.HdeniSipQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetLeniIpv6SipQuota() *int32 {
	return s.LeniIpv6SipQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetLeniQuota() *int32 {
	return s.LeniQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetLeniSipQuota() *int32 {
	return s.LeniSipQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetLniSipQuota() *int32 {
	return s.LniSipQuota
}

func (s *GetNodeInfoForPodResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeInfoForPodResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNodeInfoForPodResponseBodyContent) GetVSwitches() []*string {
	return s.VSwitches
}

func (s *GetNodeInfoForPodResponseBodyContent) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNodeInfoForPodResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetNodeInfoForPodResponseBodyContent) SetClusterId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetHdeniIpv6SipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.HdeniIpv6SipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetHdeniQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.HdeniQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetHdeniSipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.HdeniSipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLeniIpv6SipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LeniIpv6SipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLeniQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LeniQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLeniSipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LeniSipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetLniSipQuota(v int32) *GetNodeInfoForPodResponseBodyContent {
	s.LniSipQuota = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetNodeId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetRegionId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetVSwitches(v []*string) *GetNodeInfoForPodResponseBodyContent {
	s.VSwitches = v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetVpcId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) SetZoneId(v string) *GetNodeInfoForPodResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *GetNodeInfoForPodResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
