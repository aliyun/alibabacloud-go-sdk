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
	// The configurations of the index.
	//
	// >  You must specify either IndexStatus or IndexConfig.
	//
	// example:
	//
	// {}
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// The state of the index. Valid values:
	//
	// 	- active (default): the index is enabled.
	//
	// 	- Deactive: the index is not enabled.
	//
	// >  You must specify either IndexStatus or IndexConfig.
	//
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// The category of the index. Valid values:
	//
	// 	- mm: large visual model.
	//
	// 	- face: face recognition.
	//
	// 	- aiLabel: smart tagging.
	//
	// This parameter is required.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The name of the search library.
	//
	// 	- If you leave this parameter empty, the search index is created in the default search library of Intelligent Media Service (IMS). Default value: ims-default-search-lib.
	//
	// 	- To query information about an existing search library, call the [QuerySearchLib](https://help.aliyun.com/document_detail/2584455.html) API operation.
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
