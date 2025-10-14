// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointNetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnPrefix(v string) *ModifyCustomEndpointNetRequest
	GetConnPrefix() *string
	SetCustomEndpointId(v string) *ModifyCustomEndpointNetRequest
	GetCustomEndpointId() *string
	SetDBInstanceName(v string) *ModifyCustomEndpointNetRequest
	GetDBInstanceName() *string
	SetPort(v int32) *ModifyCustomEndpointNetRequest
	GetPort() *int32
	SetRegionId(v string) *ModifyCustomEndpointNetRequest
	GetRegionId() *string
	SetVSwitchId(v string) *ModifyCustomEndpointNetRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyCustomEndpointNetRequest
	GetVpcId() *string
}

type ModifyCustomEndpointNetRequest struct {
	// example:
	//
	// pxc-****
	ConnPrefix *string `json:"ConnPrefix,omitempty" xml:"ConnPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cep-002
	CustomEndpointId *string `json:"CustomEndpointId,omitempty" xml:"CustomEndpointId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1ndoug37dtwoedlmru0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyCustomEndpointNetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointNetRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointNetRequest) GetConnPrefix() *string {
	return s.ConnPrefix
}

func (s *ModifyCustomEndpointNetRequest) GetCustomEndpointId() *string {
	return s.CustomEndpointId
}

func (s *ModifyCustomEndpointNetRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyCustomEndpointNetRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyCustomEndpointNetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCustomEndpointNetRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyCustomEndpointNetRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyCustomEndpointNetRequest) SetConnPrefix(v string) *ModifyCustomEndpointNetRequest {
	s.ConnPrefix = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetCustomEndpointId(v string) *ModifyCustomEndpointNetRequest {
	s.CustomEndpointId = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetDBInstanceName(v string) *ModifyCustomEndpointNetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetPort(v int32) *ModifyCustomEndpointNetRequest {
	s.Port = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetRegionId(v string) *ModifyCustomEndpointNetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetVSwitchId(v string) *ModifyCustomEndpointNetRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) SetVpcId(v string) *ModifyCustomEndpointNetRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyCustomEndpointNetRequest) Validate() error {
	return dara.Validate(s)
}
