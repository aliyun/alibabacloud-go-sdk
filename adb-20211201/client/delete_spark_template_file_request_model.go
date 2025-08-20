// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteSparkTemplateFileRequest
	GetDBClusterId() *string
	SetId(v int64) *DeleteSparkTemplateFileRequest
	GetId() *int64
}

type DeleteSparkTemplateFileRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1y769u11748****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the template file to be deleted.
	//
	// >  You can call the [GetSparkTemplateFullTree](https://help.aliyun.com/document_detail/456205.html) operation to query all template file IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 284
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSparkTemplateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateFileRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteSparkTemplateFileRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSparkTemplateFileRequest) SetDBClusterId(v string) *DeleteSparkTemplateFileRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSparkTemplateFileRequest) SetId(v int64) *DeleteSparkTemplateFileRequest {
	s.Id = &v
	return s
}

func (s *DeleteSparkTemplateFileRequest) Validate() error {
	return dara.Validate(s)
}
