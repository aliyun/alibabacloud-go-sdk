// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexConfig(v string) *AlterSearchIndexRequest
	GetIndexConfig() *string
	SetIndexStatus(v string) *AlterSearchIndexRequest
	GetIndexStatus() *string
	SetIndexType(v string) *AlterSearchIndexRequest
	GetIndexType() *string
	SetSearchLibName(v string) *AlterSearchIndexRequest
	GetSearchLibName() *string
}

type AlterSearchIndexRequest struct {
	// The index configuration.
	//
	// 	Notice:  You must specify either IndexStatus or IndexConfig.
	//
	// example:
	//
	// {}
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// The index status. Default value: Active. Valid values:
	//
	// - Active: activated.
	//
	// - Deactive: deactivated.
	//
	// 	Notice:  You must specify either IndexStatus or IndexConfig.
	//
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// The index type. Valid values:
	//
	// - mm: large model.
	//
	// - face: face.
	//
	// - aiLabel: intelligent tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The search library name.
	//
	// - If no search library name is specified, the search index is created in the default IMS search library. Default value: ims-default-search-lib.
	//
	// - You can call the [QuerySearchLib](https://help.aliyun.com/document_detail/2584455.html) operation to query existing search library information.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s AlterSearchIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchIndexRequest) GoString() string {
	return s.String()
}

func (s *AlterSearchIndexRequest) GetIndexConfig() *string {
	return s.IndexConfig
}

func (s *AlterSearchIndexRequest) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *AlterSearchIndexRequest) GetIndexType() *string {
	return s.IndexType
}

func (s *AlterSearchIndexRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *AlterSearchIndexRequest) SetIndexConfig(v string) *AlterSearchIndexRequest {
	s.IndexConfig = &v
	return s
}

func (s *AlterSearchIndexRequest) SetIndexStatus(v string) *AlterSearchIndexRequest {
	s.IndexStatus = &v
	return s
}

func (s *AlterSearchIndexRequest) SetIndexType(v string) *AlterSearchIndexRequest {
	s.IndexType = &v
	return s
}

func (s *AlterSearchIndexRequest) SetSearchLibName(v string) *AlterSearchIndexRequest {
	s.SearchLibName = &v
	return s
}

func (s *AlterSearchIndexRequest) Validate() error {
	return dara.Validate(s)
}
