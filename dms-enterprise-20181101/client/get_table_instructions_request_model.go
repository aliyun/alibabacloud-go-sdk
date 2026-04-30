// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInstructionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *GetTableInstructionsRequest
	GetDbId() *int32
	SetTableName(v string) *GetTableInstructionsRequest
	GetTableName() *string
}

type GetTableInstructionsRequest struct {
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

func (s GetTableInstructionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableInstructionsRequest) GoString() string {
	return s.String()
}

func (s *GetTableInstructionsRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetTableInstructionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableInstructionsRequest) SetDbId(v int32) *GetTableInstructionsRequest {
	s.DbId = &v
	return s
}

func (s *GetTableInstructionsRequest) SetTableName(v string) *GetTableInstructionsRequest {
	s.TableName = &v
	return s
}

func (s *GetTableInstructionsRequest) Validate() error {
	return dara.Validate(s)
}
