// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStageModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStageModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStageModelsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeStageModelsRequest
	GetSecurityToken() *string
	SetStageAlias(v string) *DescribeStageModelsRequest
	GetStageAlias() *string
	SetStageName(v string) *DescribeStageModelsRequest
	GetStageName() *string
}

type DescribeStageModelsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageAlias    *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeStageModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStageModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStageModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStageModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStageModelsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeStageModelsRequest) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribeStageModelsRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeStageModelsRequest) SetPageNumber(v int32) *DescribeStageModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStageModelsRequest) SetPageSize(v int32) *DescribeStageModelsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStageModelsRequest) SetSecurityToken(v string) *DescribeStageModelsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeStageModelsRequest) SetStageAlias(v string) *DescribeStageModelsRequest {
	s.StageAlias = &v
	return s
}

func (s *DescribeStageModelsRequest) SetStageName(v string) *DescribeStageModelsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeStageModelsRequest) Validate() error {
	return dara.Validate(s)
}
