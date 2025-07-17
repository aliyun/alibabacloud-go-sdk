// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLColumn interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DLColumn
	GetComment() *string
	SetName(v string) *DLColumn
	GetName() *string
	SetType(v string) *DLColumn
	GetType() *string
}

type DLColumn struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DLColumn) String() string {
	return dara.Prettify(s)
}

func (s DLColumn) GoString() string {
	return s.String()
}

func (s *DLColumn) GetComment() *string {
	return s.Comment
}

func (s *DLColumn) GetName() *string {
	return s.Name
}

func (s *DLColumn) GetType() *string {
	return s.Type
}

func (s *DLColumn) SetComment(v string) *DLColumn {
	s.Comment = &v
	return s
}

func (s *DLColumn) SetName(v string) *DLColumn {
	s.Name = &v
	return s
}

func (s *DLColumn) SetType(v string) *DLColumn {
	s.Type = &v
	return s
}

func (s *DLColumn) Validate() error {
	return dara.Validate(s)
}
