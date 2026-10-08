// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSemanticQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *SemanticQueryRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *SemanticQueryRequest
	GetMaxResults() *int32
	SetMediaTypes(v []*string) *SemanticQueryRequest
	GetMediaTypes() []*string
	SetNextToken(v string) *SemanticQueryRequest
	GetNextToken() *string
	SetProjectName(v string) *SemanticQueryRequest
	GetProjectName() *string
	SetQuery(v string) *SemanticQueryRequest
	GetQuery() *string
	SetSourceURI(v string) *SemanticQueryRequest
	GetSourceURI() *string
	SetWithFields(v []*string) *SemanticQueryRequest
	GetWithFields() []*string
}

type SemanticQueryRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of data records to return in this request. Value range: (0,100].
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The media types to search. If this parameter is left empty, the default value is:
	MediaTypes []*string `json:"MediaTypes,omitempty" xml:"MediaTypes,omitempty" type:"Repeated"`
	// This parameter is no longer provided.
	//
	// example:
	//
	// Reserved. Not supported yet.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// <notice>Either this parameter or the SourceURI parameter must be specified.</notice>
	//
	// The content for semantic search.
	//
	// example:
	//
	// Scenery of Hangzhou in April 2021
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// <notice>Either this parameter or the Query parameter must be specified. This parameter is currently valid only when the search type is specified as image and the dataset is configured with a workflow template for image-to-image search.</notice>
	//
	// The storage address of the source data used for retrieval. The storage address supports OSS URIs.
	//
	// The OSS address format is oss://${Bucket}/${Object}, where ${Bucket} is the name of the OSS bucket that resides in the same region as the current project, and ${Object} is the full path of the file including the file name extension.
	//
	// If you need to configure the corresponding workflow template, [contact us](https://help.aliyun.com/document_detail/84454.html).
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Specifies the specific fields to return instead of all existing metadata fields. This helps reduce the size of the returned struct.
	//
	// If this parameter is left empty, all fields are returned.
	WithFields []*string `json:"WithFields,omitempty" xml:"WithFields,omitempty" type:"Repeated"`
}

func (s SemanticQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s SemanticQueryRequest) GoString() string {
	return s.String()
}

func (s *SemanticQueryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SemanticQueryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SemanticQueryRequest) GetMediaTypes() []*string {
	return s.MediaTypes
}

func (s *SemanticQueryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SemanticQueryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SemanticQueryRequest) GetQuery() *string {
	return s.Query
}

func (s *SemanticQueryRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *SemanticQueryRequest) GetWithFields() []*string {
	return s.WithFields
}

func (s *SemanticQueryRequest) SetDatasetName(v string) *SemanticQueryRequest {
	s.DatasetName = &v
	return s
}

func (s *SemanticQueryRequest) SetMaxResults(v int32) *SemanticQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *SemanticQueryRequest) SetMediaTypes(v []*string) *SemanticQueryRequest {
	s.MediaTypes = v
	return s
}

func (s *SemanticQueryRequest) SetNextToken(v string) *SemanticQueryRequest {
	s.NextToken = &v
	return s
}

func (s *SemanticQueryRequest) SetProjectName(v string) *SemanticQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *SemanticQueryRequest) SetQuery(v string) *SemanticQueryRequest {
	s.Query = &v
	return s
}

func (s *SemanticQueryRequest) SetSourceURI(v string) *SemanticQueryRequest {
	s.SourceURI = &v
	return s
}

func (s *SemanticQueryRequest) SetWithFields(v []*string) *SemanticQueryRequest {
	s.WithFields = v
	return s
}

func (s *SemanticQueryRequest) Validate() error {
	return dara.Validate(s)
}
