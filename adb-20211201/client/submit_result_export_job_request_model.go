// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitResultExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SubmitResultExportJobRequest
	GetDBClusterId() *string
	SetEngine(v string) *SubmitResultExportJobRequest
	GetEngine() *string
	SetExportType(v string) *SubmitResultExportJobRequest
	GetExportType() *string
	SetRegionId(v string) *SubmitResultExportJobRequest
	GetRegionId() *string
	SetResourceGroup(v string) *SubmitResultExportJobRequest
	GetResourceGroup() *string
	SetSQL(v string) *SubmitResultExportJobRequest
	GetSQL() *string
	SetSchema(v string) *SubmitResultExportJobRequest
	GetSchema() *string
}

type SubmitResultExportJobRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp10a0ng21t5****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The engine that is used to run the result set export job. Set the value to XIHE.
	//
	// example:
	//
	// XIHE
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of the result set export job.
	//
	// example:
	//
	// SLS
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource group that runs the result set export job.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The SQL statement that is used in the result set export job. You can specify only SELECT statements. If you specify other SQL statements, the request fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM `fotor_com_datastore_resource`.`fotor_ai_create_task` where `date` = \\"2023-05-13\\" LIMIT 10;
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpch_oss
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s SubmitResultExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitResultExportJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitResultExportJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SubmitResultExportJobRequest) GetEngine() *string {
	return s.Engine
}

func (s *SubmitResultExportJobRequest) GetExportType() *string {
	return s.ExportType
}

func (s *SubmitResultExportJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SubmitResultExportJobRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *SubmitResultExportJobRequest) GetSQL() *string {
	return s.SQL
}

func (s *SubmitResultExportJobRequest) GetSchema() *string {
	return s.Schema
}

func (s *SubmitResultExportJobRequest) SetDBClusterId(v string) *SubmitResultExportJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetEngine(v string) *SubmitResultExportJobRequest {
	s.Engine = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetExportType(v string) *SubmitResultExportJobRequest {
	s.ExportType = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetRegionId(v string) *SubmitResultExportJobRequest {
	s.RegionId = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetResourceGroup(v string) *SubmitResultExportJobRequest {
	s.ResourceGroup = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetSQL(v string) *SubmitResultExportJobRequest {
	s.SQL = &v
	return s
}

func (s *SubmitResultExportJobRequest) SetSchema(v string) *SubmitResultExportJobRequest {
	s.Schema = &v
	return s
}

func (s *SubmitResultExportJobRequest) Validate() error {
	return dara.Validate(s)
}
