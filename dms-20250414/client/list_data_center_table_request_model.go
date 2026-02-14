// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *ListDataCenterTableRequest
	GetCallFrom() *string
	SetDatabaseName(v string) *ListDataCenterTableRequest
	GetDatabaseName() *string
	SetDmsUnit(v string) *ListDataCenterTableRequest
	GetDmsUnit() *string
	SetImportType(v string) *ListDataCenterTableRequest
	GetImportType() *string
	SetInstanceName(v string) *ListDataCenterTableRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ListDataCenterTableRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataCenterTableRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListDataCenterTableRequest
	GetSearchKey() *string
	SetTableName(v string) *ListDataCenterTableRequest
	GetTableName() *string
}

type ListDataCenterTableRequest struct {
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// diamonds.csv
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// example:
	//
	// FILE
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// example:
	//
	// f-ean8u5881qk4*********xh5y
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// testdb
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// diamonds
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDataCenterTableRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterTableRequest) GoString() string {
	return s.String()
}

func (s *ListDataCenterTableRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *ListDataCenterTableRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDataCenterTableRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *ListDataCenterTableRequest) GetImportType() *string {
	return s.ImportType
}

func (s *ListDataCenterTableRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDataCenterTableRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataCenterTableRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCenterTableRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataCenterTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDataCenterTableRequest) SetCallFrom(v string) *ListDataCenterTableRequest {
	s.CallFrom = &v
	return s
}

func (s *ListDataCenterTableRequest) SetDatabaseName(v string) *ListDataCenterTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListDataCenterTableRequest) SetDmsUnit(v string) *ListDataCenterTableRequest {
	s.DmsUnit = &v
	return s
}

func (s *ListDataCenterTableRequest) SetImportType(v string) *ListDataCenterTableRequest {
	s.ImportType = &v
	return s
}

func (s *ListDataCenterTableRequest) SetInstanceName(v string) *ListDataCenterTableRequest {
	s.InstanceName = &v
	return s
}

func (s *ListDataCenterTableRequest) SetPageNumber(v int32) *ListDataCenterTableRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCenterTableRequest) SetPageSize(v int32) *ListDataCenterTableRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCenterTableRequest) SetSearchKey(v string) *ListDataCenterTableRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataCenterTableRequest) SetTableName(v string) *ListDataCenterTableRequest {
	s.TableName = &v
	return s
}

func (s *ListDataCenterTableRequest) Validate() error {
	return dara.Validate(s)
}
