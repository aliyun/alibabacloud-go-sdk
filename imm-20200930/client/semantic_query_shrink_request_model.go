// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSemanticQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *SemanticQueryShrinkRequest
	GetDatasetName() *string
	SetMaxResults(v int32) *SemanticQueryShrinkRequest
	GetMaxResults() *int32
	SetMediaTypesShrink(v string) *SemanticQueryShrinkRequest
	GetMediaTypesShrink() *string
	SetNextToken(v string) *SemanticQueryShrinkRequest
	GetNextToken() *string
	SetProjectName(v string) *SemanticQueryShrinkRequest
	GetProjectName() *string
	SetQuery(v string) *SemanticQueryShrinkRequest
	GetQuery() *string
	SetSourceURI(v string) *SemanticQueryShrinkRequest
	GetSourceURI() *string
	SetWithFieldsShrink(v string) *SemanticQueryShrinkRequest
	GetWithFieldsShrink() *string
}

type SemanticQueryShrinkRequest struct {
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
	MediaTypesShrink *string `json:"MediaTypes,omitempty" xml:"MediaTypes,omitempty"`
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
	WithFieldsShrink *string `json:"WithFields,omitempty" xml:"WithFields,omitempty"`
}

func (s SemanticQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SemanticQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *SemanticQueryShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SemanticQueryShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SemanticQueryShrinkRequest) GetMediaTypesShrink() *string {
	return s.MediaTypesShrink
}

func (s *SemanticQueryShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SemanticQueryShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SemanticQueryShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *SemanticQueryShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *SemanticQueryShrinkRequest) GetWithFieldsShrink() *string {
	return s.WithFieldsShrink
}

func (s *SemanticQueryShrinkRequest) SetDatasetName(v string) *SemanticQueryShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetMaxResults(v int32) *SemanticQueryShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetMediaTypesShrink(v string) *SemanticQueryShrinkRequest {
	s.MediaTypesShrink = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetNextToken(v string) *SemanticQueryShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetProjectName(v string) *SemanticQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetQuery(v string) *SemanticQueryShrinkRequest {
	s.Query = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetSourceURI(v string) *SemanticQueryShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *SemanticQueryShrinkRequest) SetWithFieldsShrink(v string) *SemanticQueryShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *SemanticQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
