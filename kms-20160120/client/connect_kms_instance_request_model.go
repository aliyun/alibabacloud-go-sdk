// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectKmsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKMProvider(v string) *ConnectKmsInstanceRequest
	GetKMProvider() *string
	SetKmsInstanceId(v string) *ConnectKmsInstanceRequest
	GetKmsInstanceId() *string
	SetVSwitchIds(v string) *ConnectKmsInstanceRequest
	GetVSwitchIds() *string
	SetVpcId(v string) *ConnectKmsInstanceRequest
	GetVpcId() *string
	SetZoneIds(v string) *ConnectKmsInstanceRequest
	GetZoneIds() *string
}

type ConnectKmsInstanceRequest struct {
	// The provider of the KMS instance. Currently, only Aliyun is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	KMProvider *string `json:"KMProvider,omitempty" xml:"KMProvider,omitempty"`
	// The ID of the KMS instance to enable.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-phzz64f722a1buamw0****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
	// Sets a vSwitch for the dual-zone deployment. The vSwitch must have at least one available IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1i512amda6d10a0****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) for the KMS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp19z7cwmltad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Sets two zones for the KMS instance. Dual-zone load balancing improves service availability and disaster recovery capabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-k,cn-hangzhou-j
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s ConnectKmsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConnectKmsInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConnectKmsInstanceRequest) GetKMProvider() *string {
	return s.KMProvider
}

func (s *ConnectKmsInstanceRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *ConnectKmsInstanceRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ConnectKmsInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ConnectKmsInstanceRequest) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *ConnectKmsInstanceRequest) SetKMProvider(v string) *ConnectKmsInstanceRequest {
	s.KMProvider = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetKmsInstanceId(v string) *ConnectKmsInstanceRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetVSwitchIds(v string) *ConnectKmsInstanceRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetVpcId(v string) *ConnectKmsInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *ConnectKmsInstanceRequest) SetZoneIds(v string) *ConnectKmsInstanceRequest {
	s.ZoneIds = &v
	return s
}

func (s *ConnectKmsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
