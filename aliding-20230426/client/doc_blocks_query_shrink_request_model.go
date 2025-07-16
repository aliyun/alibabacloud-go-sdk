// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockType(v string) *DocBlocksQueryShrinkRequest
	GetBlockType() *string
	SetDocKey(v string) *DocBlocksQueryShrinkRequest
	GetDocKey() *string
	SetEndIndex(v int32) *DocBlocksQueryShrinkRequest
	GetEndIndex() *int32
	SetStartIndex(v int32) *DocBlocksQueryShrinkRequest
	GetStartIndex() *int32
	SetTenantContextShrink(v string) *DocBlocksQueryShrinkRequest
	GetTenantContextShrink() *string
}

type DocBlocksQueryShrinkRequest struct {
	// example:
	//
	// heading
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	DocKey *string `json:"DocKey,omitempty" xml:"DocKey,omitempty"`
	// example:
	//
	// 1
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// 0
	StartIndex          *int32  `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DocBlocksQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryShrinkRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *DocBlocksQueryShrinkRequest) GetDocKey() *string {
	return s.DocKey
}

func (s *DocBlocksQueryShrinkRequest) GetEndIndex() *int32 {
	return s.EndIndex
}

func (s *DocBlocksQueryShrinkRequest) GetStartIndex() *int32 {
	return s.StartIndex
}

func (s *DocBlocksQueryShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DocBlocksQueryShrinkRequest) SetBlockType(v string) *DocBlocksQueryShrinkRequest {
	s.BlockType = &v
	return s
}

func (s *DocBlocksQueryShrinkRequest) SetDocKey(v string) *DocBlocksQueryShrinkRequest {
	s.DocKey = &v
	return s
}

func (s *DocBlocksQueryShrinkRequest) SetEndIndex(v int32) *DocBlocksQueryShrinkRequest {
	s.EndIndex = &v
	return s
}

func (s *DocBlocksQueryShrinkRequest) SetStartIndex(v int32) *DocBlocksQueryShrinkRequest {
	s.StartIndex = &v
	return s
}

func (s *DocBlocksQueryShrinkRequest) SetTenantContextShrink(v string) *DocBlocksQueryShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DocBlocksQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
