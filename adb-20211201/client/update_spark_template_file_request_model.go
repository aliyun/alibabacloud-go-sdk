// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSparkTemplateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateSparkTemplateFileRequest
	GetContent() *string
	SetDBClusterId(v string) *UpdateSparkTemplateFileRequest
	GetDBClusterId() *string
	SetId(v int64) *UpdateSparkTemplateFileRequest
	GetId() *int64
	SetResourceGroupName(v string) *UpdateSparkTemplateFileRequest
	GetResourceGroupName() *string
}

type UpdateSparkTemplateFileRequest struct {
	// The template data to be updated.
	//
	// >  If you do not specify this parameter, the application template is not updated. For information about how to configure a Spark application template, see [Configure a Spark application](https://help.aliyun.com/document_detail/452402.html).
	//
	// example:
	//
	// set spark.driver.resourceSpec=medium;set spark.executor.instances=2;set spark.executor.resourceSpec=medium;set spark.app.name=Spark SQL Test;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-pz5vp4585l466****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application template ID.
	//
	// >  You can call the [GetSparkTemplateFullTree](https://help.aliyun.com/document_detail/456205.html) operation to query the application template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 718056
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the job resource group.
	//
	// example:
	//
	// adb
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s UpdateSparkTemplateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSparkTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateSparkTemplateFileRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSparkTemplateFileRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateSparkTemplateFileRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSparkTemplateFileRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UpdateSparkTemplateFileRequest) SetContent(v string) *UpdateSparkTemplateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetDBClusterId(v string) *UpdateSparkTemplateFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetId(v int64) *UpdateSparkTemplateFileRequest {
	s.Id = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) SetResourceGroupName(v string) *UpdateSparkTemplateFileRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *UpdateSparkTemplateFileRequest) Validate() error {
	return dara.Validate(s)
}
