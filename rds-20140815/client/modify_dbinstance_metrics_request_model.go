// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyDBInstanceMetricsRequest
	GetDBInstanceName() *string
	SetMetricsConfig(v string) *ModifyDBInstanceMetricsRequest
	GetMetricsConfig() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceMetricsRequest
	GetResourceOwnerId() *int64
	SetScope(v string) *ModifyDBInstanceMetricsRequest
	GetScope() *string
}

type ModifyDBInstanceMetricsRequest struct {
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1s1j103lo6****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The monitoring metrics to configure for the instance. You can specify multiple metric keys separated by commas (,). A maximum of 30 metric keys can be specified.
	//
	// You can call the DescribeAvailableMetrics operation to obtain the enhanced monitoring metric keys.
	//
	// This parameter is required.
	//
	// example:
	//
	// os.cpu_usage.sys.avg,os.cpu_usage.user.avg
	MetricsConfig   *string `json:"MetricsConfig,omitempty" xml:"MetricsConfig,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scope of the modification. Valid values:
	//
	// 	- **instance**: instance level. The modification is applied only to cloud disk instance.
	//
	// 	- **region**: region level. The modification is applied to all ApsaraDB RDS for PostgreSQL instances that use the same storage type as cloud disk instance in the current region. For example, if cloud disk instance uses cloud disks, the modification is applied to all ApsaraDB RDS for PostgreSQL instances with cloud disks in the current region.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ModifyDBInstanceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMetricsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceMetricsRequest) GetMetricsConfig() *string {
	return s.MetricsConfig
}

func (s *ModifyDBInstanceMetricsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceMetricsRequest) GetScope() *string {
	return s.Scope
}

func (s *ModifyDBInstanceMetricsRequest) SetDBInstanceName(v string) *ModifyDBInstanceMetricsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceMetricsRequest) SetMetricsConfig(v string) *ModifyDBInstanceMetricsRequest {
	s.MetricsConfig = &v
	return s
}

func (s *ModifyDBInstanceMetricsRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMetricsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMetricsRequest) SetScope(v string) *ModifyDBInstanceMetricsRequest {
	s.Scope = &v
	return s
}

func (s *ModifyDBInstanceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
