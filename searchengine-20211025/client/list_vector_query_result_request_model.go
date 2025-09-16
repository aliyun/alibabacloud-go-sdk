// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *ListVectorQueryResultRequest
	GetBody() map[string]interface{}
	SetPath(v string) *ListVectorQueryResultRequest
	GetPath() *string
	SetQueryType(v string) *ListVectorQueryResultRequest
	GetQueryType() *string
	SetVectorQueryType(v string) *ListVectorQueryResultRequest
	GetVectorQueryType() *string
}

type ListVectorQueryResultRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	Path *string                `json:"path,omitempty" xml:"path,omitempty"`
	// The query type. Valid values: vector, primary_key, and vector_text.
	//
	// example:
	//
	// primary_key
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// The vector query type. Valid values: vector, image, and text.
	//
	// example:
	//
	// image
	VectorQueryType *string `json:"vectorQueryType,omitempty" xml:"vectorQueryType,omitempty"`
}

func (s ListVectorQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVectorQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListVectorQueryResultRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *ListVectorQueryResultRequest) GetPath() *string {
	return s.Path
}

func (s *ListVectorQueryResultRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListVectorQueryResultRequest) GetVectorQueryType() *string {
	return s.VectorQueryType
}

func (s *ListVectorQueryResultRequest) SetBody(v map[string]interface{}) *ListVectorQueryResultRequest {
	s.Body = v
	return s
}

func (s *ListVectorQueryResultRequest) SetPath(v string) *ListVectorQueryResultRequest {
	s.Path = &v
	return s
}

func (s *ListVectorQueryResultRequest) SetQueryType(v string) *ListVectorQueryResultRequest {
	s.QueryType = &v
	return s
}

func (s *ListVectorQueryResultRequest) SetVectorQueryType(v string) *ListVectorQueryResultRequest {
	s.VectorQueryType = &v
	return s
}

func (s *ListVectorQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
