// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteSparkTemplateRequest
	GetDBClusterId() *string
	SetId(v int64) *DeleteSparkTemplateRequest
	GetId() *int64
}

type DeleteSparkTemplateRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The directory ID or application ID of the template files that you want to delete.
	//
	// >
	//
	// 	- You can call the [GetSparkTemplateFullTree](https://help.aliyun.com/document_detail/612467.html) operation to query the directory ID or application ID.
	//
	// 	- When you specify a directory ID, the directory and all template files that are included in the directory are deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 725204
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSparkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteSparkTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSparkTemplateRequest) SetDBClusterId(v string) *DeleteSparkTemplateRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSparkTemplateRequest) SetId(v int64) *DeleteSparkTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteSparkTemplateRequest) Validate() error {
	return dara.Validate(s)
}
