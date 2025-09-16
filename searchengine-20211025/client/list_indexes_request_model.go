// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListIndexesRequest
	GetCatalog() *string
	SetDatabase(v string) *ListIndexesRequest
	GetDatabase() *string
	SetNewMode(v bool) *ListIndexesRequest
	GetNewMode() *bool
	SetTable(v string) *ListIndexesRequest
	GetTable() *string
}

type ListIndexesRequest struct {
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// Specifies whether the OpenSearch Vector Search Edition instance is of the new version.
	//
	// example:
	//
	// true
	NewMode *bool   `json:"newMode,omitempty" xml:"newMode,omitempty"`
	Table   *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s ListIndexesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListIndexesRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListIndexesRequest) GetNewMode() *bool {
	return s.NewMode
}

func (s *ListIndexesRequest) GetTable() *string {
	return s.Table
}

func (s *ListIndexesRequest) SetCatalog(v string) *ListIndexesRequest {
	s.Catalog = &v
	return s
}

func (s *ListIndexesRequest) SetDatabase(v string) *ListIndexesRequest {
	s.Database = &v
	return s
}

func (s *ListIndexesRequest) SetNewMode(v bool) *ListIndexesRequest {
	s.NewMode = &v
	return s
}

func (s *ListIndexesRequest) SetTable(v string) *ListIndexesRequest {
	s.Table = &v
	return s
}

func (s *ListIndexesRequest) Validate() error {
	return dara.Validate(s)
}
