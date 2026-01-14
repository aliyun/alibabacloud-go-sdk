// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIds(v []*int64) *RetrieveRequest
	GetDocumentIds() []*int64
	SetFolderIds(v []*string) *RetrieveRequest
	GetFolderIds() []*string
	SetPreRetrieveTopK(v int32) *RetrieveRequest
	GetPreRetrieveTopK() *int32
	SetQuery(v string) *RetrieveRequest
	GetQuery() *string
	SetRerankerThreshold(v float32) *RetrieveRequest
	GetRerankerThreshold() *float32
	SetTopK(v int32) *RetrieveRequest
	GetTopK() *int32
	SetUseReranker(v bool) *RetrieveRequest
	GetUseReranker() *bool
}

type RetrieveRequest struct {
	DocumentIds []*int64  `json:"documentIds,omitempty" xml:"documentIds,omitempty" type:"Repeated"`
	FolderIds   []*string `json:"folderIds,omitempty" xml:"folderIds,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	PreRetrieveTopK *int32 `json:"preRetrieveTopK,omitempty" xml:"preRetrieveTopK,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "What\\"s the meaning of the file?"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// 0.0
	RerankerThreshold *float32 `json:"rerankerThreshold,omitempty" xml:"rerankerThreshold,omitempty"`
	// example:
	//
	// 10
	TopK        *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
	UseReranker *bool  `json:"useReranker,omitempty" xml:"useReranker,omitempty"`
}

func (s RetrieveRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRequest) GetDocumentIds() []*int64 {
	return s.DocumentIds
}

func (s *RetrieveRequest) GetFolderIds() []*string {
	return s.FolderIds
}

func (s *RetrieveRequest) GetPreRetrieveTopK() *int32 {
	return s.PreRetrieveTopK
}

func (s *RetrieveRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveRequest) GetRerankerThreshold() *float32 {
	return s.RerankerThreshold
}

func (s *RetrieveRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RetrieveRequest) GetUseReranker() *bool {
	return s.UseReranker
}

func (s *RetrieveRequest) SetDocumentIds(v []*int64) *RetrieveRequest {
	s.DocumentIds = v
	return s
}

func (s *RetrieveRequest) SetFolderIds(v []*string) *RetrieveRequest {
	s.FolderIds = v
	return s
}

func (s *RetrieveRequest) SetPreRetrieveTopK(v int32) *RetrieveRequest {
	s.PreRetrieveTopK = &v
	return s
}

func (s *RetrieveRequest) SetQuery(v string) *RetrieveRequest {
	s.Query = &v
	return s
}

func (s *RetrieveRequest) SetRerankerThreshold(v float32) *RetrieveRequest {
	s.RerankerThreshold = &v
	return s
}

func (s *RetrieveRequest) SetTopK(v int32) *RetrieveRequest {
	s.TopK = &v
	return s
}

func (s *RetrieveRequest) SetUseReranker(v bool) *RetrieveRequest {
	s.UseReranker = &v
	return s
}

func (s *RetrieveRequest) Validate() error {
	return dara.Validate(s)
}
