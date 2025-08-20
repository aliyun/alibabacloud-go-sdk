// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkSQLEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *StartSparkSQLEngineRequest
	GetConfig() *string
	SetDBClusterId(v string) *StartSparkSQLEngineRequest
	GetDBClusterId() *string
	SetJars(v string) *StartSparkSQLEngineRequest
	GetJars() *string
	SetMaxExecutor(v int64) *StartSparkSQLEngineRequest
	GetMaxExecutor() *int64
	SetMinExecutor(v int64) *StartSparkSQLEngineRequest
	GetMinExecutor() *int64
	SetResourceGroupName(v string) *StartSparkSQLEngineRequest
	GetResourceGroupName() *string
	SetSlotNum(v int64) *StartSparkSQLEngineRequest
	GetSlotNum() *int64
}

type StartSparkSQLEngineRequest struct {
	// The configuration that is required to start the Spark SQL engine. Specify this value in the JSON format. For more information, see [Conf configuration parameters](https://help.aliyun.com/document_detail/471203.html).
	//
	// example:
	//
	// { "spark.shuffle.timeout": ":0s" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-abcd****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Object Storage Service (OSS) paths of third-party JAR packages that are required to start the Spark SQL engine. Separate multiple OSS paths with commas (,).
	//
	// example:
	//
	// oss://testBuckname/test.jar,oss://testBuckname/test2.jar
	Jars *string `json:"Jars,omitempty" xml:"Jars,omitempty"`
	// The maximum number of executors that are required to execute SQL statements. Valid values: 1 to 2000. If this value exceeds the total number of executes that are supported by the resource group, the Spark SQL engine fails to be started.
	//
	// example:
	//
	// 10
	MaxExecutor *int64 `json:"MaxExecutor,omitempty" xml:"MaxExecutor,omitempty"`
	// The minimum number of executors that are required to execute SQL statements. Valid values: 0 to 2000. A value of 0 indicates that no executors are permanent if no SQL statements are executed. If this value exceeds the total number of executors that are supported by the resource group, the Spark SQL engine fails to be started. The value must be less than the value of MaxExecutor.
	//
	// example:
	//
	// 1
	MinExecutor *int64 `json:"MinExecutor,omitempty" xml:"MinExecutor,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// spark-rg-name
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The maximum number of slots that are required to maintain Spark sessions for executing SQL statements. Valid values: 1 to 500.
	//
	// example:
	//
	// 100
	SlotNum *int64 `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
}

func (s StartSparkSQLEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSparkSQLEngineRequest) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineRequest) GetConfig() *string {
	return s.Config
}

func (s *StartSparkSQLEngineRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StartSparkSQLEngineRequest) GetJars() *string {
	return s.Jars
}

func (s *StartSparkSQLEngineRequest) GetMaxExecutor() *int64 {
	return s.MaxExecutor
}

func (s *StartSparkSQLEngineRequest) GetMinExecutor() *int64 {
	return s.MinExecutor
}

func (s *StartSparkSQLEngineRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *StartSparkSQLEngineRequest) GetSlotNum() *int64 {
	return s.SlotNum
}

func (s *StartSparkSQLEngineRequest) SetConfig(v string) *StartSparkSQLEngineRequest {
	s.Config = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetDBClusterId(v string) *StartSparkSQLEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetJars(v string) *StartSparkSQLEngineRequest {
	s.Jars = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetMaxExecutor(v int64) *StartSparkSQLEngineRequest {
	s.MaxExecutor = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetMinExecutor(v int64) *StartSparkSQLEngineRequest {
	s.MinExecutor = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetResourceGroupName(v string) *StartSparkSQLEngineRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *StartSparkSQLEngineRequest) SetSlotNum(v int64) *StartSparkSQLEngineRequest {
	s.SlotNum = &v
	return s
}

func (s *StartSparkSQLEngineRequest) Validate() error {
	return dara.Validate(s)
}
