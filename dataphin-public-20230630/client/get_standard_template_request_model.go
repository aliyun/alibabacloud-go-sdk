// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQuery(v *GetStandardTemplateRequestFilterQuery) *GetStandardTemplateRequest
	GetFilterQuery() *GetStandardTemplateRequestFilterQuery
	SetId(v int64) *GetStandardTemplateRequest
	GetId() *int64
	SetNullable(v bool) *GetStandardTemplateRequest
	GetNullable() *bool
	SetOpTenantId(v int64) *GetStandardTemplateRequest
	GetOpTenantId() *int64
}

type GetStandardTemplateRequest struct {
	FilterQuery *GetStandardTemplateRequestFilterQuery `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id       *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	Nullable *bool  `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetStandardTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateRequest) GetFilterQuery() *GetStandardTemplateRequestFilterQuery {
	return s.FilterQuery
}

func (s *GetStandardTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetStandardTemplateRequest) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardTemplateRequest) SetFilterQuery(v *GetStandardTemplateRequestFilterQuery) *GetStandardTemplateRequest {
	s.FilterQuery = v
	return s
}

func (s *GetStandardTemplateRequest) SetId(v int64) *GetStandardTemplateRequest {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateRequest) SetNullable(v bool) *GetStandardTemplateRequest {
	s.Nullable = &v
	return s
}

func (s *GetStandardTemplateRequest) SetOpTenantId(v int64) *GetStandardTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardTemplateRequest) Validate() error {
	if s.FilterQuery != nil {
		if err := s.FilterQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateRequestFilterQuery struct {
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardTemplateRequestFilterQuery) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateRequestFilterQuery) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateRequestFilterQuery) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardTemplateRequestFilterQuery) SetVersion(v int32) *GetStandardTemplateRequestFilterQuery {
	s.Version = &v
	return s
}

func (s *GetStandardTemplateRequestFilterQuery) Validate() error {
	return dara.Validate(s)
}
