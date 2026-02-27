// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRowPermissionByTableGuidsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetRowPermissionByTableGuidsQuery(v *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) *GetRowPermissionByTableGuidsRequest
	GetGetRowPermissionByTableGuidsQuery() *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery
	SetOpTenantId(v int64) *GetRowPermissionByTableGuidsRequest
	GetOpTenantId() *int64
}

type GetRowPermissionByTableGuidsRequest struct {
	// This parameter is required.
	GetRowPermissionByTableGuidsQuery *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery `json:"GetRowPermissionByTableGuidsQuery,omitempty" xml:"GetRowPermissionByTableGuidsQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetRowPermissionByTableGuidsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsRequest) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsRequest) GetGetRowPermissionByTableGuidsQuery() *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery {
	return s.GetRowPermissionByTableGuidsQuery
}

func (s *GetRowPermissionByTableGuidsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetRowPermissionByTableGuidsRequest) SetGetRowPermissionByTableGuidsQuery(v *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) *GetRowPermissionByTableGuidsRequest {
	s.GetRowPermissionByTableGuidsQuery = v
	return s
}

func (s *GetRowPermissionByTableGuidsRequest) SetOpTenantId(v int64) *GetRowPermissionByTableGuidsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetRowPermissionByTableGuidsRequest) Validate() error {
	if s.GetRowPermissionByTableGuidsQuery != nil {
		if err := s.GetRowPermissionByTableGuidsQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery struct {
	// This parameter is required.
	TableGuids []*string `json:"TableGuids,omitempty" xml:"TableGuids,omitempty" type:"Repeated"`
}

func (s GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) GetTableGuids() []*string {
	return s.TableGuids
}

func (s *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) SetTableGuids(v []*string) *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery {
	s.TableGuids = v
	return s
}

func (s *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) Validate() error {
	return dara.Validate(s)
}
