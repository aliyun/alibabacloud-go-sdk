// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteNormalizationSchemaRequest
	GetLang() *string
	SetNormalizationSchemaId(v string) *DeleteNormalizationSchemaRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSchemaType(v string) *DeleteNormalizationSchemaRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *DeleteNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteNormalizationSchemaRequest
	GetRoleFor() *int64
}

type DeleteNormalizationSchemaRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The normalization schema ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_QUERY_DNS_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The normalization schema type. Valid values:
	//
	// - log: log.
	//
	// - entity: entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteNormalizationSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationSchemaRequest) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNormalizationSchemaRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *DeleteNormalizationSchemaRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *DeleteNormalizationSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNormalizationSchemaRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteNormalizationSchemaRequest) SetLang(v string) *DeleteNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNormalizationSchemaRequest) SetNormalizationSchemaId(v string) *DeleteNormalizationSchemaRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *DeleteNormalizationSchemaRequest) SetNormalizationSchemaType(v string) *DeleteNormalizationSchemaRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *DeleteNormalizationSchemaRequest) SetRegionId(v string) *DeleteNormalizationSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNormalizationSchemaRequest) SetRoleFor(v int64) *DeleteNormalizationSchemaRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteNormalizationSchemaRequest) Validate() error {
	return dara.Validate(s)
}
