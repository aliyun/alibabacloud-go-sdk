// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexType(v string) *DropSearchIndexRequest
	GetIndexType() *string
	SetSearchLibName(v string) *DropSearchIndexRequest
	GetSearchLibName() *string
}

type DropSearchIndexRequest struct {
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

func (s DropSearchIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DropSearchIndexRequest) GoString() string {
	return s.String()
}

func (s *DropSearchIndexRequest) GetIndexType() *string {
	return s.IndexType
}

func (s *DropSearchIndexRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *DropSearchIndexRequest) SetIndexType(v string) *DropSearchIndexRequest {
	s.IndexType = &v
	return s
}

func (s *DropSearchIndexRequest) SetSearchLibName(v string) *DropSearchIndexRequest {
	s.SearchLibName = &v
	return s
}

func (s *DropSearchIndexRequest) Validate() error {
	return dara.Validate(s)
}
