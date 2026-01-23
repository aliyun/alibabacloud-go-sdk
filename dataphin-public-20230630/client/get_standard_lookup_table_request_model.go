// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardLookupTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetStandardLookupTableRequest
	GetId() *int64
	SetNullable(v bool) *GetStandardLookupTableRequest
	GetNullable() *bool
	SetOpTenantId(v int64) *GetStandardLookupTableRequest
	GetOpTenantId() *int64
}

type GetStandardLookupTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id       *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	Nullable *bool  `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetStandardLookupTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableRequest) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableRequest) GetId() *int64 {
	return s.Id
}

func (s *GetStandardLookupTableRequest) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardLookupTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardLookupTableRequest) SetId(v int64) *GetStandardLookupTableRequest {
	s.Id = &v
	return s
}

func (s *GetStandardLookupTableRequest) SetNullable(v bool) *GetStandardLookupTableRequest {
	s.Nullable = &v
	return s
}

func (s *GetStandardLookupTableRequest) SetOpTenantId(v int64) *GetStandardLookupTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardLookupTableRequest) Validate() error {
	return dara.Validate(s)
}
