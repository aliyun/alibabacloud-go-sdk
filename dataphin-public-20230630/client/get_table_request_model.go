// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTableRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetTableRequest
	GetOpUserId() *string
	SetTableGuid(v string) *GetTableRequest
	GetTableGuid() *string
}

type GetTableRequest struct {
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
	// example:
	//
	// dp_ds_table.123.12344.db.table
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetTableRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableRequest) SetOpTenantId(v int64) *GetTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableRequest) SetOpUserId(v string) *GetTableRequest {
	s.OpUserId = &v
	return s
}

func (s *GetTableRequest) SetTableGuid(v string) *GetTableRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableRequest) Validate() error {
	return dara.Validate(s)
}
