// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkSQLEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *KillSparkSQLEngineRequest
	GetDBClusterId() *string
	SetResourceGroupName(v string) *KillSparkSQLEngineRequest
	GetResourceGroupName() *string
}

type KillSparkSQLEngineRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbn8pq537k8w****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// spark_test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s KillSparkSQLEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s KillSparkSQLEngineRequest) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillSparkSQLEngineRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *KillSparkSQLEngineRequest) SetDBClusterId(v string) *KillSparkSQLEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillSparkSQLEngineRequest) SetResourceGroupName(v string) *KillSparkSQLEngineRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *KillSparkSQLEngineRequest) Validate() error {
	return dara.Validate(s)
}
