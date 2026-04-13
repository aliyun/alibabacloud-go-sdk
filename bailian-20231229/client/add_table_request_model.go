// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *AddTableRequest
	GetConnectorId() *string
	SetTableColumns(v []*AddTableRequestTableColumns) *AddTableRequest
	GetTableColumns() []*AddTableRequestTableColumns
	SetTableName(v string) *AddTableRequest
	GetTableName() *string
}

type AddTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// connector-d51861492df64257
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// This parameter is required.
	TableColumns []*AddTableRequestTableColumns `json:"TableColumns,omitempty" xml:"TableColumns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ads_wjjr_basc_enterprise_di
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s AddTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTableRequest) GoString() string {
	return s.String()
}

func (s *AddTableRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AddTableRequest) GetTableColumns() []*AddTableRequestTableColumns {
	return s.TableColumns
}

func (s *AddTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *AddTableRequest) SetConnectorId(v string) *AddTableRequest {
	s.ConnectorId = &v
	return s
}

func (s *AddTableRequest) SetTableColumns(v []*AddTableRequestTableColumns) *AddTableRequest {
	s.TableColumns = v
	return s
}

func (s *AddTableRequest) SetTableName(v string) *AddTableRequest {
	s.TableName = &v
	return s
}

func (s *AddTableRequest) Validate() error {
	if s.TableColumns != nil {
		for _, item := range s.TableColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddTableRequestTableColumns struct {
	// example:
	//
	// desc
	ColumnDesc *string `json:"ColumnDesc,omitempty" xml:"ColumnDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s AddTableRequestTableColumns) String() string {
	return dara.Prettify(s)
}

func (s AddTableRequestTableColumns) GoString() string {
	return s.String()
}

func (s *AddTableRequestTableColumns) GetColumnDesc() *string {
	return s.ColumnDesc
}

func (s *AddTableRequestTableColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *AddTableRequestTableColumns) GetDataType() *string {
	return s.DataType
}

func (s *AddTableRequestTableColumns) SetColumnDesc(v string) *AddTableRequestTableColumns {
	s.ColumnDesc = &v
	return s
}

func (s *AddTableRequestTableColumns) SetColumnName(v string) *AddTableRequestTableColumns {
	s.ColumnName = &v
	return s
}

func (s *AddTableRequestTableColumns) SetDataType(v string) *AddTableRequestTableColumns {
	s.DataType = &v
	return s
}

func (s *AddTableRequestTableColumns) Validate() error {
	return dara.Validate(s)
}
