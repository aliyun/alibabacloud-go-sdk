// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DLDatabase
	GetCatalogName() *string
	SetDbId(v int64) *DLDatabase
	GetDbId() *int64
	SetDescription(v string) *DLDatabase
	GetDescription() *string
	SetLocation(v string) *DLDatabase
	GetLocation() *string
	SetName(v string) *DLDatabase
	GetName() *string
	SetParameters(v map[string]interface{}) *DLDatabase
	GetParameters() map[string]interface{}
}

type DLDatabase struct {
	CatalogName *string                `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	DbId        *int64                 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Location    *string                `json:"Location,omitempty" xml:"Location,omitempty"`
	Name        *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters  map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s DLDatabase) String() string {
	return dara.Prettify(s)
}

func (s DLDatabase) GoString() string {
	return s.String()
}

func (s *DLDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DLDatabase) GetDbId() *int64 {
	return s.DbId
}

func (s *DLDatabase) GetDescription() *string {
	return s.Description
}

func (s *DLDatabase) GetLocation() *string {
	return s.Location
}

func (s *DLDatabase) GetName() *string {
	return s.Name
}

func (s *DLDatabase) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DLDatabase) SetCatalogName(v string) *DLDatabase {
	s.CatalogName = &v
	return s
}

func (s *DLDatabase) SetDbId(v int64) *DLDatabase {
	s.DbId = &v
	return s
}

func (s *DLDatabase) SetDescription(v string) *DLDatabase {
	s.Description = &v
	return s
}

func (s *DLDatabase) SetLocation(v string) *DLDatabase {
	s.Location = &v
	return s
}

func (s *DLDatabase) SetName(v string) *DLDatabase {
	s.Name = &v
	return s
}

func (s *DLDatabase) SetParameters(v map[string]interface{}) *DLDatabase {
	s.Parameters = v
	return s
}

func (s *DLDatabase) Validate() error {
	return dara.Validate(s)
}
