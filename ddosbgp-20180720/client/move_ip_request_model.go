// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *MoveIpRequest
	GetInstanceId() *string
	SetIpList(v string) *MoveIpRequest
	GetIpList() *string
	SetRegionId(v string) *MoveIpRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *MoveIpRequest
	GetSourceInstanceId() *string
}

type MoveIpRequest struct {
	// The ID of the destination Anti-DDoS Origin instance.
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// example:
	//
	// ddos_originpre_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses to migrate to an Anti-DDoS Origin instance for protection. The list is represented as a character string in JSON array format. Each element in the JSON array is a structure that contains the following fields:
	//
	// - **ip**: String type. Required. The IP address to add.
	//
	// - **member_uid**: String type. Optional. The member accounts to which the asset belongs. This field is required only when you add an asset under member accounts. Example: [{"ip":"121.41.XX.XX","member_uid":"120100811162****"}]
	//
	// example:
	//
	// [{"ip":"121.41.XX.XX"},{"ip":"121.42.XX.XX"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// > You can call the DescribeRegions operation to query all regions supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source Anti-DDoS Origin instance.
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// example:
	//
	// ddos_originpre_public_cn-***
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s MoveIpRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveIpRequest) GoString() string {
	return s.String()
}

func (s *MoveIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *MoveIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveIpRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *MoveIpRequest) SetInstanceId(v string) *MoveIpRequest {
	s.InstanceId = &v
	return s
}

func (s *MoveIpRequest) SetIpList(v string) *MoveIpRequest {
	s.IpList = &v
	return s
}

func (s *MoveIpRequest) SetRegionId(v string) *MoveIpRequest {
	s.RegionId = &v
	return s
}

func (s *MoveIpRequest) SetSourceInstanceId(v string) *MoveIpRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *MoveIpRequest) Validate() error {
	return dara.Validate(s)
}
