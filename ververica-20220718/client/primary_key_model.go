// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrimaryKey interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*string) *PrimaryKey
	GetColumns() []*string
	SetConstraintName(v string) *PrimaryKey
	GetConstraintName() *string
	SetConstraintType(v string) *PrimaryKey
	GetConstraintType() *string
	SetEnforced(v bool) *PrimaryKey
	GetEnforced() *bool
}

type PrimaryKey struct {
	// This parameter is required.
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// This parameter is required.
	ConstraintName *string `json:"constraintName,omitempty" xml:"constraintName,omitempty"`
	// This parameter is required.
	ConstraintType *string `json:"constraintType,omitempty" xml:"constraintType,omitempty"`
	// This parameter is required.
	Enforced *bool `json:"enforced,omitempty" xml:"enforced,omitempty"`
}

func (s PrimaryKey) String() string {
	return dara.Prettify(s)
}

func (s PrimaryKey) GoString() string {
	return s.String()
}

func (s *PrimaryKey) GetColumns() []*string {
	return s.Columns
}

func (s *PrimaryKey) GetConstraintName() *string {
	return s.ConstraintName
}

func (s *PrimaryKey) GetConstraintType() *string {
	return s.ConstraintType
}

func (s *PrimaryKey) GetEnforced() *bool {
	return s.Enforced
}

func (s *PrimaryKey) SetColumns(v []*string) *PrimaryKey {
	s.Columns = v
	return s
}

func (s *PrimaryKey) SetConstraintName(v string) *PrimaryKey {
	s.ConstraintName = &v
	return s
}

func (s *PrimaryKey) SetConstraintType(v string) *PrimaryKey {
	s.ConstraintType = &v
	return s
}

func (s *PrimaryKey) SetEnforced(v bool) *PrimaryKey {
	s.Enforced = &v
	return s
}

func (s *PrimaryKey) Validate() error {
	return dara.Validate(s)
}
