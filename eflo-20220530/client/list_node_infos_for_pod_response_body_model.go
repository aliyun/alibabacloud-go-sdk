// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInfosForPodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListNodeInfosForPodResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListNodeInfosForPodResponseBody
	GetCode() *int32
	SetContent(v []*ListNodeInfosForPodResponseBodyContent) *ListNodeInfosForPodResponseBody
	GetContent() []*ListNodeInfosForPodResponseBodyContent
	SetMessage(v string) *ListNodeInfosForPodResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNodeInfosForPodResponseBody
	GetRequestId() *string
}

type ListNodeInfosForPodResponseBody struct {
	// The information about the request denial.
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
	// Response body
	Content []*ListNodeInfosForPodResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// You don\\"t have the permission of this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeInfosForPodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInfosForPodResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListNodeInfosForPodResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListNodeInfosForPodResponseBody) GetContent() []*ListNodeInfosForPodResponseBodyContent {
	return s.Content
}

func (s *ListNodeInfosForPodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNodeInfosForPodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeInfosForPodResponseBody) SetAccessDeniedDetail(v string) *ListNodeInfosForPodResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetCode(v int32) *ListNodeInfosForPodResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetContent(v []*ListNodeInfosForPodResponseBodyContent) *ListNodeInfosForPodResponseBody {
	s.Content = v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetMessage(v string) *ListNodeInfosForPodResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) SetRequestId(v string) *ListNodeInfosForPodResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeInfosForPodResponseBodyContent struct {
	// The cluster ID.
	//
	// example:
	//
	// cluster-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun Gaomi network interface controller quota
	//
	// example:
	//
	// 10
	HdeniQuota *int32 `json:"HdeniQuota,omitempty" xml:"HdeniQuota,omitempty"`
	// Lingjun Elastic Network Interface quota, excluding system type
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
	// List of VSwitches to which IP addresses can be applied for this node
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

func (s ListNodeInfosForPodResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInfosForPodResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListNodeInfosForPodResponseBodyContent) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeInfosForPodResponseBodyContent) GetHdeniQuota() *int32 {
	return s.HdeniQuota
}

func (s *ListNodeInfosForPodResponseBodyContent) GetLeniQuota() *int32 {
	return s.LeniQuota
}

func (s *ListNodeInfosForPodResponseBodyContent) GetLeniSipQuota() *int32 {
	return s.LeniSipQuota
}

func (s *ListNodeInfosForPodResponseBodyContent) GetLniSipQuota() *int32 {
	return s.LniSipQuota
}

func (s *ListNodeInfosForPodResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNodeInfosForPodResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNodeInfosForPodResponseBodyContent) GetVSwitches() []*string {
	return s.VSwitches
}

func (s *ListNodeInfosForPodResponseBodyContent) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNodeInfosForPodResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNodeInfosForPodResponseBodyContent) SetClusterId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.ClusterId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetHdeniQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.HdeniQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLeniQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LeniQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLeniSipQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LeniSipQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetLniSipQuota(v int32) *ListNodeInfosForPodResponseBodyContent {
	s.LniSipQuota = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetNodeId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetRegionId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetVSwitches(v []*string) *ListNodeInfosForPodResponseBodyContent {
	s.VSwitches = v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetVpcId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.VpcId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) SetZoneId(v string) *ListNodeInfosForPodResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *ListNodeInfosForPodResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
