// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomEndpointId(v string) *DeleteCustomEndpointRequest
	GetCustomEndpointId() *string
	SetDBInstanceName(v string) *DeleteCustomEndpointRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteCustomEndpointRequest
	GetRegionId() *string
}

type DeleteCustomEndpointRequest struct {
	// example:
	//
	// pxe-8if3zrfsu****hgw
	CustomEndpointId *string `json:"CustomEndpointId,omitempty" xml:"CustomEndpointId,omitempty"`
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomEndpointRequest) GetCustomEndpointId() *string {
	return s.CustomEndpointId
}

func (s *DeleteCustomEndpointRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteCustomEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomEndpointRequest) SetCustomEndpointId(v string) *DeleteCustomEndpointRequest {
	s.CustomEndpointId = &v
	return s
}

func (s *DeleteCustomEndpointRequest) SetDBInstanceName(v string) *DeleteCustomEndpointRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteCustomEndpointRequest) SetRegionId(v string) *DeleteCustomEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomEndpointRequest) Validate() error {
	return dara.Validate(s)
}
