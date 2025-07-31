// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceDependenciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataSourceDependenciesRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetDataSourceDependenciesRequest
	GetOpTenantId() *int64
}

type GetDataSourceDependenciesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1023211
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetDataSourceDependenciesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDependenciesRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceDependenciesRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataSourceDependenciesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataSourceDependenciesRequest) SetId(v int64) *GetDataSourceDependenciesRequest {
	s.Id = &v
	return s
}

func (s *GetDataSourceDependenciesRequest) SetOpTenantId(v int64) *GetDataSourceDependenciesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataSourceDependenciesRequest) Validate() error {
	return dara.Validate(s)
}
