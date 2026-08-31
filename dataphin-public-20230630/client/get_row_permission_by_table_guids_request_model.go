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
	SetOpUserId(v string) *GetRowPermissionByTableGuidsRequest
	GetOpUserId() *string
}

type GetRowPermissionByTableGuidsRequest struct {
	// The request command.
	//
	// This parameter is required.
	GetRowPermissionByTableGuidsQuery *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery `json:"GetRowPermissionByTableGuidsQuery,omitempty" xml:"GetRowPermissionByTableGuidsQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *GetRowPermissionByTableGuidsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetRowPermissionByTableGuidsRequest) SetGetRowPermissionByTableGuidsQuery(v *GetRowPermissionByTableGuidsRequestGetRowPermissionByTableGuidsQuery) *GetRowPermissionByTableGuidsRequest {
	s.GetRowPermissionByTableGuidsQuery = v
	return s
}

func (s *GetRowPermissionByTableGuidsRequest) SetOpTenantId(v int64) *GetRowPermissionByTableGuidsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetRowPermissionByTableGuidsRequest) SetOpUserId(v string) *GetRowPermissionByTableGuidsRequest {
	s.OpUserId = &v
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
	// The list of table GUIDs.
	//
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
