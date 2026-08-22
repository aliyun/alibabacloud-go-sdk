// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchAccessProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyOpenSearchAccessProtocolRequest
	GetDBInstanceName() *string
	SetProtocol(v string) *ModifyOpenSearchAccessProtocolRequest
	GetProtocol() *string
	SetRegionId(v string) *ModifyOpenSearchAccessProtocolRequest
	GetRegionId() *string
}

type ModifyOpenSearchAccessProtocolRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The access protocol. Valid values:
	//
	// - **http**: HTTP protocol.
	//
	// - **https**: HTTPS protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOpenSearchAccessProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchAccessProtocolRequest) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchAccessProtocolRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyOpenSearchAccessProtocolRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyOpenSearchAccessProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOpenSearchAccessProtocolRequest) SetDBInstanceName(v string) *ModifyOpenSearchAccessProtocolRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolRequest) SetProtocol(v string) *ModifyOpenSearchAccessProtocolRequest {
	s.Protocol = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolRequest) SetRegionId(v string) *ModifyOpenSearchAccessProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolRequest) Validate() error {
	return dara.Validate(s)
}
