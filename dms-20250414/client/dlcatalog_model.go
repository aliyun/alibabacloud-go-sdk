// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DLCatalog
	GetDescription() *string
	SetLocation(v string) *DLCatalog
	GetLocation() *string
	SetName(v string) *DLCatalog
	GetName() *string
}

type DLCatalog struct {
	// example:
	//
	// init default catalog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// oss://xxxx
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// hive
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DLCatalog) String() string {
	return dara.Prettify(s)
}

func (s DLCatalog) GoString() string {
	return s.String()
}

func (s *DLCatalog) GetDescription() *string {
	return s.Description
}

func (s *DLCatalog) GetLocation() *string {
	return s.Location
}

func (s *DLCatalog) GetName() *string {
	return s.Name
}

func (s *DLCatalog) SetDescription(v string) *DLCatalog {
	s.Description = &v
	return s
}

func (s *DLCatalog) SetLocation(v string) *DLCatalog {
	s.Location = &v
	return s
}

func (s *DLCatalog) SetName(v string) *DLCatalog {
	s.Name = &v
	return s
}

func (s *DLCatalog) Validate() error {
	return dara.Validate(s)
}
