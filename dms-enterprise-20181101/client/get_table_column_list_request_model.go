// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *GetTableColumnListRequest
	GetDbId() *int32
	SetTableName(v string) *GetTableColumnListRequest
	GetTableName() *string
}

type GetTableColumnListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order_info
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnListRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnListRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetTableColumnListRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableColumnListRequest) SetDbId(v int32) *GetTableColumnListRequest {
	s.DbId = &v
	return s
}

func (s *GetTableColumnListRequest) SetTableName(v string) *GetTableColumnListRequest {
	s.TableName = &v
	return s
}

func (s *GetTableColumnListRequest) Validate() error {
	return dara.Validate(s)
}
