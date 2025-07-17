// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageEntity interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *LineageEntity
	GetAttributes() map[string]*string
	SetId(v string) *LineageEntity
	GetId() *string
	SetName(v string) *LineageEntity
	GetName() *string
}

type LineageEntity struct {
	// example:
	//
	// {"key1":"value1"}
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// maxcompute-table:123456::test_project::test_tbl
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test_tbl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s LineageEntity) String() string {
	return dara.Prettify(s)
}

func (s LineageEntity) GoString() string {
	return s.String()
}

func (s *LineageEntity) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *LineageEntity) GetId() *string {
	return s.Id
}

func (s *LineageEntity) GetName() *string {
	return s.Name
}

func (s *LineageEntity) SetAttributes(v map[string]*string) *LineageEntity {
	s.Attributes = v
	return s
}

func (s *LineageEntity) SetId(v string) *LineageEntity {
	s.Id = &v
	return s
}

func (s *LineageEntity) SetName(v string) *LineageEntity {
	s.Name = &v
	return s
}

func (s *LineageEntity) Validate() error {
	return dara.Validate(s)
}
