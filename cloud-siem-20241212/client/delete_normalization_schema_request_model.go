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
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the normalization schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROCESS_QUERY_DNS_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The type of the normalization schema. Valid values:
	//
	// - log: Log.
	//
	// - entity: Entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
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
