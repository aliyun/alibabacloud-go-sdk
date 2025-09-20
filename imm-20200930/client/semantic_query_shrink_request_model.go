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
	// immDatatest
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The types of the media that you want to query. Default value:
	//
	// ["image"]
	MediaTypesShrink *string `json:"MediaTypes,omitempty" xml:"MediaTypes,omitempty"`
	// This parameter is no longer available.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The content of the query that you input.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The fields that you want to include in the response. Including only necessary metadata fields can help reduce the size of the response.
	//
	// If you do not specify this parameter or set the value to null, all existing metadata fields are returned.
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

func (s *SemanticQueryShrinkRequest) SetWithFieldsShrink(v string) *SemanticQueryShrinkRequest {
	s.WithFieldsShrink = &v
	return s
}

func (s *SemanticQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
