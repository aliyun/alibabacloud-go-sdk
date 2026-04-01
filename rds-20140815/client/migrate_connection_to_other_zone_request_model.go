// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateConnectionToOtherZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *MigrateConnectionToOtherZoneRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *MigrateConnectionToOtherZoneRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *MigrateConnectionToOtherZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerId(v int64) *MigrateConnectionToOtherZoneRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *MigrateConnectionToOtherZoneRequest
	GetZoneId() *string
}

type MigrateConnectionToOtherZoneRequest struct {
	// The endpoint of the instance. The endpoint is specified when you create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1oypo6ky19y****
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1oypo6ky19y****
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s MigrateConnectionToOtherZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateConnectionToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *MigrateConnectionToOtherZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateConnectionToOtherZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateConnectionToOtherZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateConnectionToOtherZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateConnectionToOtherZoneRequest) SetConnectionString(v string) *MigrateConnectionToOtherZoneRequest {
	s.ConnectionString = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetDBInstanceId(v string) *MigrateConnectionToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetOwnerId(v int64) *MigrateConnectionToOtherZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetResourceOwnerId(v int64) *MigrateConnectionToOtherZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) SetZoneId(v string) *MigrateConnectionToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateConnectionToOtherZoneRequest) Validate() error {
	return dara.Validate(s)
}
