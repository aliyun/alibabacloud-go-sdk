// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchDBInstanceVpcRequest
	GetDBInstanceId() *string
	SetPrivateIpAddress(v string) *SwitchDBInstanceVpcRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *SwitchDBInstanceVpcRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *SwitchDBInstanceVpcRequest
	GetVPCId() *string
	SetVSwitchId(v string) *SwitchDBInstanceVpcRequest
	GetVSwitchId() *string
}

type SwitchDBInstanceVpcRequest struct {
	// This parameter is required.
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s SwitchDBInstanceVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceVpcRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceVpcRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceVpcRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *SwitchDBInstanceVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchDBInstanceVpcRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *SwitchDBInstanceVpcRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SwitchDBInstanceVpcRequest) SetDBInstanceId(v string) *SwitchDBInstanceVpcRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetPrivateIpAddress(v string) *SwitchDBInstanceVpcRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetResourceOwnerId(v int64) *SwitchDBInstanceVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetVPCId(v string) *SwitchDBInstanceVpcRequest {
	s.VPCId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) SetVSwitchId(v string) *SwitchDBInstanceVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *SwitchDBInstanceVpcRequest) Validate() error {
	return dara.Validate(s)
}
