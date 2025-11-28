// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPrivateRAGServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeployPrivateRAGServiceRequest
	GetDBInstanceId() *string
	SetVSwitchId(v string) *DeployPrivateRAGServiceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *DeployPrivateRAGServiceRequest
	GetZoneId() *string
}

type DeployPrivateRAGServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeployPrivateRAGServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployPrivateRAGServiceRequest) GoString() string {
	return s.String()
}

func (s *DeployPrivateRAGServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeployPrivateRAGServiceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeployPrivateRAGServiceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeployPrivateRAGServiceRequest) SetDBInstanceId(v string) *DeployPrivateRAGServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) SetVSwitchId(v string) *DeployPrivateRAGServiceRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) SetZoneId(v string) *DeployPrivateRAGServiceRequest {
	s.ZoneId = &v
	return s
}

func (s *DeployPrivateRAGServiceRequest) Validate() error {
	return dara.Validate(s)
}
