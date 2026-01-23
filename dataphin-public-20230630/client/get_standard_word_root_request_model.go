// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardWordRootRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetStandardWordRootRequest
	GetName() *string
	SetNullable(v bool) *GetStandardWordRootRequest
	GetNullable() *bool
	SetOpTenantId(v int64) *GetStandardWordRootRequest
	GetOpTenantId() *int64
}

type GetStandardWordRootRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nullable *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetStandardWordRootRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardWordRootRequest) GoString() string {
	return s.String()
}

func (s *GetStandardWordRootRequest) GetName() *string {
	return s.Name
}

func (s *GetStandardWordRootRequest) GetNullable() *bool {
	return s.Nullable
}

func (s *GetStandardWordRootRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardWordRootRequest) SetName(v string) *GetStandardWordRootRequest {
	s.Name = &v
	return s
}

func (s *GetStandardWordRootRequest) SetNullable(v bool) *GetStandardWordRootRequest {
	s.Nullable = &v
	return s
}

func (s *GetStandardWordRootRequest) SetOpTenantId(v int64) *GetStandardWordRootRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardWordRootRequest) Validate() error {
	return dara.Validate(s)
}
