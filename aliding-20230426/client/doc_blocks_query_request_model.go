// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockType(v string) *DocBlocksQueryRequest
	GetBlockType() *string
	SetDocKey(v string) *DocBlocksQueryRequest
	GetDocKey() *string
	SetEndIndex(v int32) *DocBlocksQueryRequest
	GetEndIndex() *int32
	SetStartIndex(v int32) *DocBlocksQueryRequest
	GetStartIndex() *int32
	SetTenantContext(v *DocBlocksQueryRequestTenantContext) *DocBlocksQueryRequest
	GetTenantContext() *DocBlocksQueryRequestTenantContext
}

type DocBlocksQueryRequest struct {
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
	StartIndex    *int32                              `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
	TenantContext *DocBlocksQueryRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DocBlocksQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryRequest) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *DocBlocksQueryRequest) GetDocKey() *string {
	return s.DocKey
}

func (s *DocBlocksQueryRequest) GetEndIndex() *int32 {
	return s.EndIndex
}

func (s *DocBlocksQueryRequest) GetStartIndex() *int32 {
	return s.StartIndex
}

func (s *DocBlocksQueryRequest) GetTenantContext() *DocBlocksQueryRequestTenantContext {
	return s.TenantContext
}

func (s *DocBlocksQueryRequest) SetBlockType(v string) *DocBlocksQueryRequest {
	s.BlockType = &v
	return s
}

func (s *DocBlocksQueryRequest) SetDocKey(v string) *DocBlocksQueryRequest {
	s.DocKey = &v
	return s
}

func (s *DocBlocksQueryRequest) SetEndIndex(v int32) *DocBlocksQueryRequest {
	s.EndIndex = &v
	return s
}

func (s *DocBlocksQueryRequest) SetStartIndex(v int32) *DocBlocksQueryRequest {
	s.StartIndex = &v
	return s
}

func (s *DocBlocksQueryRequest) SetTenantContext(v *DocBlocksQueryRequestTenantContext) *DocBlocksQueryRequest {
	s.TenantContext = v
	return s
}

func (s *DocBlocksQueryRequest) Validate() error {
	return dara.Validate(s)
}

type DocBlocksQueryRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DocBlocksQueryRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DocBlocksQueryRequestTenantContext) SetTenantId(v string) *DocBlocksQueryRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DocBlocksQueryRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
