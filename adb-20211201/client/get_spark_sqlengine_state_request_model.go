// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkSQLEngineStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSparkSQLEngineStateRequest
	GetDBClusterId() *string
	SetResourceGroupName(v string) *GetSparkSQLEngineStateRequest
	GetResourceGroupName() *string
}

type GetSparkSQLEngineStateRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the job resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_instance
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s GetSparkSQLEngineStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkSQLEngineStateRequest) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkSQLEngineStateRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetSparkSQLEngineStateRequest) SetDBClusterId(v string) *GetSparkSQLEngineStateRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkSQLEngineStateRequest) SetResourceGroupName(v string) *GetSparkSQLEngineStateRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *GetSparkSQLEngineStateRequest) Validate() error {
	return dara.Validate(s)
}
