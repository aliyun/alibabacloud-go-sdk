// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCenterDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *ListDataCenterDatabaseRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *ListDataCenterDatabaseRequest
	GetDmsUnit() *string
	SetImportType(v string) *ListDataCenterDatabaseRequest
	GetImportType() *string
	SetLanguage(v string) *ListDataCenterDatabaseRequest
	GetLanguage() *string
	SetSearchKey(v string) *ListDataCenterDatabaseRequest
	GetSearchKey() *string
}

type ListDataCenterDatabaseRequest struct {
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// example:
	//
	// FILE
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// testdb
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListDataCenterDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCenterDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ListDataCenterDatabaseRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *ListDataCenterDatabaseRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *ListDataCenterDatabaseRequest) GetImportType() *string {
	return s.ImportType
}

func (s *ListDataCenterDatabaseRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListDataCenterDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDataCenterDatabaseRequest) SetCallFrom(v string) *ListDataCenterDatabaseRequest {
	s.CallFrom = &v
	return s
}

func (s *ListDataCenterDatabaseRequest) SetDmsUnit(v string) *ListDataCenterDatabaseRequest {
	s.DmsUnit = &v
	return s
}

func (s *ListDataCenterDatabaseRequest) SetImportType(v string) *ListDataCenterDatabaseRequest {
	s.ImportType = &v
	return s
}

func (s *ListDataCenterDatabaseRequest) SetLanguage(v string) *ListDataCenterDatabaseRequest {
	s.Language = &v
	return s
}

func (s *ListDataCenterDatabaseRequest) SetSearchKey(v string) *ListDataCenterDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDataCenterDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
