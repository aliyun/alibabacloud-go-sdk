// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetStandardSetRequest
	GetId() *int64
	SetNullable(v bool) *GetStandardSetRequest
	GetNullable() *bool
	SetOpTenantId(v int64) *GetStandardSetRequest
	GetOpTenantId() *int64
}

type GetStandardSetRequest struct {
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

func (s GetStandardSetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetRequest) GoString() string {
	return s.String()
}

func (s *GetStandardSetRequest) GetId() *int64 {
	return s.Id
}

func (s *GetStandardSetRequest) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardSetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardSetRequest) SetId(v int64) *GetStandardSetRequest {
	s.Id = &v
	return s
}

func (s *GetStandardSetRequest) SetNullable(v bool) *GetStandardSetRequest {
	s.Nullable = &v
	return s
}

func (s *GetStandardSetRequest) SetOpTenantId(v int64) *GetStandardSetRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardSetRequest) Validate() error {
	return dara.Validate(s)
}
