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
	MediaTypes []*string `json:"MediaTypes,omitempty" xml:"MediaTypes,omitempty" type:"Repeated"`
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
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The fields that you want to include in the response. Including only necessary metadata fields can help reduce the size of the response.
	//
	// If you do not specify this parameter or set the value to null, all existing metadata fields are returned.
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
