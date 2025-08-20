// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSource(v string) *SubmitSparkAppRequest
	GetAgentSource() *string
	SetAgentVersion(v string) *SubmitSparkAppRequest
	GetAgentVersion() *string
	SetAppName(v string) *SubmitSparkAppRequest
	GetAppName() *string
	SetAppType(v string) *SubmitSparkAppRequest
	GetAppType() *string
	SetDBClusterId(v string) *SubmitSparkAppRequest
	GetDBClusterId() *string
	SetData(v string) *SubmitSparkAppRequest
	GetData() *string
	SetResourceGroupName(v string) *SubmitSparkAppRequest
	GetResourceGroupName() *string
	SetTemplateFileId(v int64) *SubmitSparkAppRequest
	GetTemplateFileId() *int64
}

type SubmitSparkAppRequest struct {
	// The type of the client. The value can be up to 64 characters in length.
	//
	// example:
	//
	// CONSOLE
	AgentSource *string `json:"AgentSource,omitempty" xml:"AgentSource,omitempty"`
	// The version of the client. The value can be up to 64 characters in length.
	//
	// example:
	//
	// 1.091
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The name of the application. The value can be up to 64 characters in length.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- **SQL**
	//
	// 	- **STREAMING**
	//
	// 	- **BATCH*	- (default)
	//
	// example:
	//
	// SQL
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data of the application template.
	//
	// > For information about the application template configuration, see [Spark application configuration guide](https://help.aliyun.com/document_detail/452402.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// conf spark.driver.resourceSpec=small; conf spark.executor.instances=1; conf spark.executor.resourceSpec=small; conf spark.app.name=TestApp;
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The name of the job resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/612410.html) operation to query the name of a resource group within a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The ID of the application template.
	//
	// > You can call the [GetSparkTemplateFullTree](https://help.aliyun.com/document_detail/456205.html) operation to query the application template ID.
	//
	// example:
	//
	// 15
	TemplateFileId *int64 `json:"TemplateFileId,omitempty" xml:"TemplateFileId,omitempty"`
}

func (s SubmitSparkAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkAppRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppRequest) GetAgentSource() *string {
	return s.AgentSource
}

func (s *SubmitSparkAppRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *SubmitSparkAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *SubmitSparkAppRequest) GetAppType() *string {
	return s.AppType
}

func (s *SubmitSparkAppRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SubmitSparkAppRequest) GetData() *string {
	return s.Data
}

func (s *SubmitSparkAppRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *SubmitSparkAppRequest) GetTemplateFileId() *int64 {
	return s.TemplateFileId
}

func (s *SubmitSparkAppRequest) SetAgentSource(v string) *SubmitSparkAppRequest {
	s.AgentSource = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAgentVersion(v string) *SubmitSparkAppRequest {
	s.AgentVersion = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAppName(v string) *SubmitSparkAppRequest {
	s.AppName = &v
	return s
}

func (s *SubmitSparkAppRequest) SetAppType(v string) *SubmitSparkAppRequest {
	s.AppType = &v
	return s
}

func (s *SubmitSparkAppRequest) SetDBClusterId(v string) *SubmitSparkAppRequest {
	s.DBClusterId = &v
	return s
}

func (s *SubmitSparkAppRequest) SetData(v string) *SubmitSparkAppRequest {
	s.Data = &v
	return s
}

func (s *SubmitSparkAppRequest) SetResourceGroupName(v string) *SubmitSparkAppRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *SubmitSparkAppRequest) SetTemplateFileId(v int64) *SubmitSparkAppRequest {
	s.TemplateFileId = &v
	return s
}

func (s *SubmitSparkAppRequest) Validate() error {
	return dara.Validate(s)
}
