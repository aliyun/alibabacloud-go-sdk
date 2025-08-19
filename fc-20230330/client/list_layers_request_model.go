// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListLayersRequest
	GetLimit() *int32
	SetNextToken(v string) *ListLayersRequest
	GetNextToken() *string
	SetOfficial(v string) *ListLayersRequest
	GetOfficial() *string
	SetPrefix(v string) *ListLayersRequest
	GetPrefix() *string
	SetPublic(v string) *ListLayersRequest
	GetPublic() *string
}

type ListLayersRequest struct {
	// The number of layers that are returned
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Specifies whether the layer is official. Valid values: true and false.
	//
	// example:
	//
	// true
	Official *string `json:"official,omitempty" xml:"official,omitempty"`
	// The name prefix of the layer.
	//
	// example:
	//
	// my-layer
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// Specifies whether the layer is public. Valid values: true and false.
	//
	// example:
	//
	// true
	Public *string `json:"public,omitempty" xml:"public,omitempty"`
}

func (s ListLayersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLayersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLayersRequest) GetOfficial() *string {
	return s.Official
}

func (s *ListLayersRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListLayersRequest) GetPublic() *string {
	return s.Public
}

func (s *ListLayersRequest) SetLimit(v int32) *ListLayersRequest {
	s.Limit = &v
	return s
}

func (s *ListLayersRequest) SetNextToken(v string) *ListLayersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLayersRequest) SetOfficial(v string) *ListLayersRequest {
	s.Official = &v
	return s
}

func (s *ListLayersRequest) SetPrefix(v string) *ListLayersRequest {
	s.Prefix = &v
	return s
}

func (s *ListLayersRequest) SetPublic(v string) *ListLayersRequest {
	s.Public = &v
	return s
}

func (s *ListLayersRequest) Validate() error {
	return dara.Validate(s)
}
