// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKmsInstanceBindVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindVpcs(v string) *UpdateKmsInstanceBindVpcRequest
	GetBindVpcs() *string
	SetKmsInstanceId(v string) *UpdateKmsInstanceBindVpcRequest
	GetKmsInstanceId() *string
}

type UpdateKmsInstanceBindVpcRequest struct {
	// The VPC configuration. Each VPC configuration contains the following parameters:
	//
	// - VpcId: The ID of the VPC.
	//
	// - VSwitchId: The vSwitch in the VPC.
	//
	// - RegionID: The region where the VPC resides.
	//
	// - VpcOwnerId: The Alibaba Cloud account that owns the VPC.
	//
	// The value is a JSON string in the following format: `[{"VpcId":"${VpcId}","VSwitchId":"${VSwitchId}","RegionId":"${RegionId}","VpcOwnerId":${VpcOwnerId}},...]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"VpcId":"vpc-bp1go9qvmj78j4f4c****","VSwitchId":"vsw-bp16c5pvvcf0fp5b9****","RegionId":"cn-hangzhou","VpcOwnerId":120708975881****},{"VpcId":"vpc-bp14c07ucxg6h1xjm****","VSwitchId":"vsw-bp1wujtnspi1l3gvu****","RegionId":"cn-hangzhou","VpcOwnerId":119285303511****}]
	BindVpcs *string `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty"`
	// The ID of the KMS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-phzz64f722a1buamw0****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s UpdateKmsInstanceBindVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcRequest) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcRequest) GetBindVpcs() *string {
	return s.BindVpcs
}

func (s *UpdateKmsInstanceBindVpcRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *UpdateKmsInstanceBindVpcRequest) SetBindVpcs(v string) *UpdateKmsInstanceBindVpcRequest {
	s.BindVpcs = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcRequest) SetKmsInstanceId(v string) *UpdateKmsInstanceBindVpcRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcRequest) Validate() error {
	return dara.Validate(s)
}
