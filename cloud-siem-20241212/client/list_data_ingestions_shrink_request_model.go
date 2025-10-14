// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionIdsShrink(v string) *ListDataIngestionsShrinkRequest
	GetDataIngestionIdsShrink() *string
	SetDataIngestionStatus(v string) *ListDataIngestionsShrinkRequest
	GetDataIngestionStatus() *string
	SetDataIngestionTemplateIdsShrink(v string) *ListDataIngestionsShrinkRequest
	GetDataIngestionTemplateIdsShrink() *string
	SetLang(v string) *ListDataIngestionsShrinkRequest
	GetLang() *string
	SetProductId(v string) *ListDataIngestionsShrinkRequest
	GetProductId() *string
	SetRegionId(v string) *ListDataIngestionsShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataIngestionsShrinkRequest
	GetRoleFor() *int64
}

type ListDataIngestionsShrinkRequest struct {
	DataIngestionIdsShrink *string `json:"DataIngestionIds,omitempty" xml:"DataIngestionIds,omitempty"`
	// example:
	//
	// enabled。
	DataIngestionStatus            *string `json:"DataIngestionStatus,omitempty" xml:"DataIngestionStatus,omitempty"`
	DataIngestionTemplateIdsShrink *string `json:"DataIngestionTemplateIds,omitempty" xml:"DataIngestionTemplateIds,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataIngestionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataIngestionsShrinkRequest) GetDataIngestionIdsShrink() *string {
	return s.DataIngestionIdsShrink
}

func (s *ListDataIngestionsShrinkRequest) GetDataIngestionStatus() *string {
	return s.DataIngestionStatus
}

func (s *ListDataIngestionsShrinkRequest) GetDataIngestionTemplateIdsShrink() *string {
	return s.DataIngestionTemplateIdsShrink
}

func (s *ListDataIngestionsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataIngestionsShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListDataIngestionsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataIngestionsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataIngestionsShrinkRequest) SetDataIngestionIdsShrink(v string) *ListDataIngestionsShrinkRequest {
	s.DataIngestionIdsShrink = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetDataIngestionStatus(v string) *ListDataIngestionsShrinkRequest {
	s.DataIngestionStatus = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetDataIngestionTemplateIdsShrink(v string) *ListDataIngestionsShrinkRequest {
	s.DataIngestionTemplateIdsShrink = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetLang(v string) *ListDataIngestionsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetProductId(v string) *ListDataIngestionsShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetRegionId(v string) *ListDataIngestionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) SetRoleFor(v int64) *ListDataIngestionsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataIngestionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
