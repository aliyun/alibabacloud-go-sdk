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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1s1j103lo6****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The keys of the Enhanced Monitoring metrics that you want to display for the instance. You can enter a maximum of 30 metric keys. If you enter multiple metric keys, you must separate the metric keys with commas (,).
	//
	// You can call the DescribeAvailableMetrics operation to query the keys of metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// os.cpu_usage.sys.avg,os.cpu_usage.user.avg
	MetricsConfig   *string `json:"MetricsConfig,omitempty" xml:"MetricsConfig,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The application scope of this modification. Valid values:
	//
	// 	- **instance**: This modification is applied only to the current instance.
	//
	// 	- **region**: This modification is applied to all ApsaraDB RDS for PostgreSQL instances that are equipped with the same type of storage media as the current instance in the region to which the current instance belongs. For example, if the current instance is equipped with cloud disks, this modification is applied to all ApsaraDB RDS for PostgreSQL instances that are equipped with cloud disks in the region to which the current instance belongs.
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
