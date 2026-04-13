// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *AddTableShrinkRequest
	GetConnectorId() *string
	SetTableColumnsShrink(v string) *AddTableShrinkRequest
	GetTableColumnsShrink() *string
	SetTableName(v string) *AddTableShrinkRequest
	GetTableName() *string
}

type AddTableShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// connector-d51861492df64257
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// This parameter is required.
	TableColumnsShrink *string `json:"TableColumns,omitempty" xml:"TableColumns,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ads_wjjr_basc_enterprise_di
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s AddTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTableShrinkRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AddTableShrinkRequest) GetTableColumnsShrink() *string {
	return s.TableColumnsShrink
}

func (s *AddTableShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *AddTableShrinkRequest) SetConnectorId(v string) *AddTableShrinkRequest {
	s.ConnectorId = &v
	return s
}

func (s *AddTableShrinkRequest) SetTableColumnsShrink(v string) *AddTableShrinkRequest {
	s.TableColumnsShrink = &v
	return s
}

func (s *AddTableShrinkRequest) SetTableName(v string) *AddTableShrinkRequest {
	s.TableName = &v
	return s
}

func (s *AddTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
