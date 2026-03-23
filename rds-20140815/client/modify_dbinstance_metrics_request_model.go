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
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	MetricsConfig   *string `json:"MetricsConfig,omitempty" xml:"MetricsConfig,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
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
