// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetStandardRequest
	GetOpTenantId() *int64
	SetStandardGetQuery(v *GetStandardRequestStandardGetQuery) *GetStandardRequest
	GetStandardGetQuery() *GetStandardRequestStandardGetQuery
}

type GetStandardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StandardGetQuery *GetStandardRequestStandardGetQuery `json:"StandardGetQuery,omitempty" xml:"StandardGetQuery,omitempty" type:"Struct"`
}

func (s GetStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardRequest) GoString() string {
	return s.String()
}

func (s *GetStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardRequest) GetStandardGetQuery() *GetStandardRequestStandardGetQuery {
	return s.StandardGetQuery
}

func (s *GetStandardRequest) SetOpTenantId(v int64) *GetStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardRequest) SetStandardGetQuery(v *GetStandardRequestStandardGetQuery) *GetStandardRequest {
	s.StandardGetQuery = v
	return s
}

func (s *GetStandardRequest) Validate() error {
	if s.StandardGetQuery != nil {
		if err := s.StandardGetQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardRequestStandardGetQuery struct {
	NeedRelation *bool `json:"NeedRelation,omitempty" xml:"NeedRelation,omitempty"`
	Nullable     *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// example:
	//
	// dev
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardRequestStandardGetQuery) String() string {
	return dara.Prettify(s)
}

func (s GetStandardRequestStandardGetQuery) GoString() string {
	return s.String()
}

func (s *GetStandardRequestStandardGetQuery) GetNeedRelation() *bool {
	return s.NeedRelation
}

func (s *GetStandardRequestStandardGetQuery) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardRequestStandardGetQuery) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetStandardRequestStandardGetQuery) GetStandardStage() *string {
	return s.StandardStage
}

func (s *GetStandardRequestStandardGetQuery) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardRequestStandardGetQuery) SetNeedRelation(v bool) *GetStandardRequestStandardGetQuery {
	s.NeedRelation = &v
	return s
}

func (s *GetStandardRequestStandardGetQuery) SetNullable(v bool) *GetStandardRequestStandardGetQuery {
	s.Nullable = &v
	return s
}

func (s *GetStandardRequestStandardGetQuery) SetStandardId(v int64) *GetStandardRequestStandardGetQuery {
	s.StandardId = &v
	return s
}

func (s *GetStandardRequestStandardGetQuery) SetStandardStage(v string) *GetStandardRequestStandardGetQuery {
	s.StandardStage = &v
	return s
}

func (s *GetStandardRequestStandardGetQuery) SetVersion(v int32) *GetStandardRequestStandardGetQuery {
	s.Version = &v
	return s
}

func (s *GetStandardRequestStandardGetQuery) Validate() error {
	return dara.Validate(s)
}
