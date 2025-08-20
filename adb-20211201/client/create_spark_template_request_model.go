// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSparkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *CreateSparkTemplateRequest
	GetAppType() *string
	SetDBClusterId(v string) *CreateSparkTemplateRequest
	GetDBClusterId() *string
	SetName(v string) *CreateSparkTemplateRequest
	GetName() *string
	SetParentId(v int64) *CreateSparkTemplateRequest
	GetParentId() *int64
	SetType(v string) *CreateSparkTemplateRequest
	GetType() *string
}

type CreateSparkTemplateRequest struct {
	// The application type. Valid values:
	//
	// 	- **SQL**
	//
	// 	- **STREAMING**
	//
	// 	- **BATCH**
	//
	// >  You do not need to specify this parameter when Type is set to folder.
	//
	// example:
	//
	// SQL
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the application template. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// batchfile
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the directory to which the application template belongs.
	//
	// >  You can call the [GetSparkTemplateFolderTree](https://help.aliyun.com/document_detail/456218.html) operation to query the directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the application template. Valid values:
	//
	// 	- **folder**: directory.
	//
	// 	- **file**: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateSparkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSparkTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSparkTemplateRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateSparkTemplateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateSparkTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateSparkTemplateRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateSparkTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateSparkTemplateRequest) SetAppType(v string) *CreateSparkTemplateRequest {
	s.AppType = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetDBClusterId(v string) *CreateSparkTemplateRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetName(v string) *CreateSparkTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetParentId(v int64) *CreateSparkTemplateRequest {
	s.ParentId = &v
	return s
}

func (s *CreateSparkTemplateRequest) SetType(v string) *CreateSparkTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateSparkTemplateRequest) Validate() error {
	return dara.Validate(s)
}
