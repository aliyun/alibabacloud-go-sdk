// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetNormalizationSchemaRequest
	GetLang() *string
	SetNormalizationSchemaId(v string) *GetNormalizationSchemaRequest
	GetNormalizationSchemaId() *string
	SetNormalizationSchemaType(v string) *GetNormalizationSchemaRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *GetNormalizationSchemaRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetNormalizationSchemaRequest
	GetRoleFor() *int64
}

type GetNormalizationSchemaRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// HTTP_ACTIVITY。
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// example:
	//
	// log。
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetNormalizationSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNormalizationSchemaRequest) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationSchemaRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *GetNormalizationSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNormalizationSchemaRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetNormalizationSchemaRequest) SetLang(v string) *GetNormalizationSchemaRequest {
	s.Lang = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetNormalizationSchemaId(v string) *GetNormalizationSchemaRequest {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetNormalizationSchemaType(v string) *GetNormalizationSchemaRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetRegionId(v string) *GetNormalizationSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetNormalizationSchemaRequest) SetRoleFor(v int64) *GetNormalizationSchemaRequest {
	s.RoleFor = &v
	return s
}

func (s *GetNormalizationSchemaRequest) Validate() error {
	return dara.Validate(s)
}
