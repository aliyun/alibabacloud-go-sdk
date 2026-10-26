// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryDataRangeRequest
	GetKeyword() *string
	SetType(v string) *QueryDataRangeRequest
	GetType() *string
}

type QueryDataRangeRequest struct {
	// Name, for fuzzy search.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Data range type:
	//
	// - llmCube: LlmCube resource.
	//
	// - llmCubeTheme: Analysis theme.
	//
	// This parameter is required.
	//
	// example:
	//
	// llmCube
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRangeRequest) GoString() string {
	return s.String()
}

func (s *QueryDataRangeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryDataRangeRequest) GetType() *string {
	return s.Type
}

func (s *QueryDataRangeRequest) SetKeyword(v string) *QueryDataRangeRequest {
	s.Keyword = &v
	return s
}

func (s *QueryDataRangeRequest) SetType(v string) *QueryDataRangeRequest {
	s.Type = &v
	return s
}

func (s *QueryDataRangeRequest) Validate() error {
	return dara.Validate(s)
}
