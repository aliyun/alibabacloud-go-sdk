// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFileContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSparkTemplateFileContentRequest
	GetDBClusterId() *string
	SetId(v int64) *GetSparkTemplateFileContentRequest
	GetId() *int64
}

type GetSparkTemplateFileContentRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbn8pq537k8w****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application template ID.
	//
	// >  You can call the [GetSparkTemplateFullTree](https://help.aliyun.com/document_detail/456205.html) operation to query the application template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 725204
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSparkTemplateFileContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFileContentRequest) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFileContentRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkTemplateFileContentRequest) GetId() *int64 {
	return s.Id
}

func (s *GetSparkTemplateFileContentRequest) SetDBClusterId(v string) *GetSparkTemplateFileContentRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkTemplateFileContentRequest) SetId(v int64) *GetSparkTemplateFileContentRequest {
	s.Id = &v
	return s
}

func (s *GetSparkTemplateFileContentRequest) Validate() error {
	return dara.Validate(s)
}
