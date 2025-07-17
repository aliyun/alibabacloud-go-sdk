// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDataKeyVO interface {
	dara.Model
	String() string
	GoString() string
	SetColName(v string) *GetTableDataKeyVO
	GetColName() *string
	SetDbName(v string) *GetTableDataKeyVO
	GetDbName() *string
	SetMekId(v int64) *GetTableDataKeyVO
	GetMekId() *int64
	SetSchemaName(v string) *GetTableDataKeyVO
	GetSchemaName() *string
	SetTblName(v string) *GetTableDataKeyVO
	GetTblName() *string
	SetUserName(v string) *GetTableDataKeyVO
	GetUserName() *string
}

type GetTableDataKeyVO struct {
	ColName    *string `json:"ColName,omitempty" xml:"ColName,omitempty"`
	DbName     *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	MekId      *int64  `json:"MekId,omitempty" xml:"MekId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TblName    *string `json:"TblName,omitempty" xml:"TblName,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetTableDataKeyVO) String() string {
	return dara.Prettify(s)
}

func (s GetTableDataKeyVO) GoString() string {
	return s.String()
}

func (s *GetTableDataKeyVO) GetColName() *string {
	return s.ColName
}

func (s *GetTableDataKeyVO) GetDbName() *string {
	return s.DbName
}

func (s *GetTableDataKeyVO) GetMekId() *int64 {
	return s.MekId
}

func (s *GetTableDataKeyVO) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableDataKeyVO) GetTblName() *string {
	return s.TblName
}

func (s *GetTableDataKeyVO) GetUserName() *string {
	return s.UserName
}

func (s *GetTableDataKeyVO) SetColName(v string) *GetTableDataKeyVO {
	s.ColName = &v
	return s
}

func (s *GetTableDataKeyVO) SetDbName(v string) *GetTableDataKeyVO {
	s.DbName = &v
	return s
}

func (s *GetTableDataKeyVO) SetMekId(v int64) *GetTableDataKeyVO {
	s.MekId = &v
	return s
}

func (s *GetTableDataKeyVO) SetSchemaName(v string) *GetTableDataKeyVO {
	s.SchemaName = &v
	return s
}

func (s *GetTableDataKeyVO) SetTblName(v string) *GetTableDataKeyVO {
	s.TblName = &v
	return s
}

func (s *GetTableDataKeyVO) SetUserName(v string) *GetTableDataKeyVO {
	s.UserName = &v
	return s
}

func (s *GetTableDataKeyVO) Validate() error {
	return dara.Validate(s)
}
