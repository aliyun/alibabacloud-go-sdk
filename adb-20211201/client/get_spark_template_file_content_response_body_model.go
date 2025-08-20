// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFileContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkTemplateFileContentResponseBodyData) *GetSparkTemplateFileContentResponseBody
	GetData() *GetSparkTemplateFileContentResponseBodyData
	SetRequestId(v string) *GetSparkTemplateFileContentResponseBody
	GetRequestId() *string
}

type GetSparkTemplateFileContentResponseBody struct {
	// The returned data.
	Data *GetSparkTemplateFileContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFileContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFileContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponseBody) GetData() *GetSparkTemplateFileContentResponseBodyData {
	return s.Data
}

func (s *GetSparkTemplateFileContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkTemplateFileContentResponseBody) SetData(v *GetSparkTemplateFileContentResponseBodyData) *GetSparkTemplateFileContentResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkTemplateFileContentResponseBody) SetRequestId(v string) *GetSparkTemplateFileContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSparkTemplateFileContentResponseBodyData struct {
	// The application type. Valid values:
	//
	// 	- **SQL**
	//
	// 	- **STREAMING**
	//
	// 	- **BATCH**
	//
	// example:
	//
	// SQL
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The content of the application template.
	//
	// example:
	//
	// set spark.driver.resourceSpec=medium;set spark.executor.instances=2;set spark.executor.resourceSpec=medium;set spark.app.name=Spark SQL Test;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The application template ID.
	//
	// example:
	//
	// 725204
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The file type. Valid values:
	//
	// 	- **folder**
	//
	// 	- **file**
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSparkTemplateFileContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFileContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetSparkTemplateFileContentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetSparkTemplateFileContentResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetSparkTemplateFileContentResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetSparkTemplateFileContentResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetAppType(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetContent(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetId(v int64) *GetSparkTemplateFileContentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetResourceGroupName(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) SetType(v string) *GetSparkTemplateFileContentResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSparkTemplateFileContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
