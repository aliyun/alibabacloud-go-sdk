// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructMvBaseTableDetailModel interface {
	dara.Model
	String() string
	GoString() string
	SetDataVolumn(v string) *OpenStructMvBaseTableDetailModel
	GetDataVolumn() *string
	SetEnableBinlog(v bool) *OpenStructMvBaseTableDetailModel
	GetEnableBinlog() *bool
	SetSchemaName(v string) *OpenStructMvBaseTableDetailModel
	GetSchemaName() *string
	SetTableName(v string) *OpenStructMvBaseTableDetailModel
	GetTableName() *string
}

type OpenStructMvBaseTableDetailModel struct {
	DataVolumn   *string `json:"DataVolumn,omitempty" xml:"DataVolumn,omitempty"`
	EnableBinlog *bool   `json:"EnableBinlog,omitempty" xml:"EnableBinlog,omitempty"`
	SchemaName   *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s OpenStructMvBaseTableDetailModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMvBaseTableDetailModel) GoString() string {
	return s.String()
}

func (s *OpenStructMvBaseTableDetailModel) GetDataVolumn() *string {
	return s.DataVolumn
}

func (s *OpenStructMvBaseTableDetailModel) GetEnableBinlog() *bool {
	return s.EnableBinlog
}

func (s *OpenStructMvBaseTableDetailModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *OpenStructMvBaseTableDetailModel) GetTableName() *string {
	return s.TableName
}

func (s *OpenStructMvBaseTableDetailModel) SetDataVolumn(v string) *OpenStructMvBaseTableDetailModel {
	s.DataVolumn = &v
	return s
}

func (s *OpenStructMvBaseTableDetailModel) SetEnableBinlog(v bool) *OpenStructMvBaseTableDetailModel {
	s.EnableBinlog = &v
	return s
}

func (s *OpenStructMvBaseTableDetailModel) SetSchemaName(v string) *OpenStructMvBaseTableDetailModel {
	s.SchemaName = &v
	return s
}

func (s *OpenStructMvBaseTableDetailModel) SetTableName(v string) *OpenStructMvBaseTableDetailModel {
	s.TableName = &v
	return s
}

func (s *OpenStructMvBaseTableDetailModel) Validate() error {
	return dara.Validate(s)
}
