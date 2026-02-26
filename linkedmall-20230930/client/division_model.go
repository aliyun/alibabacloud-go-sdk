// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDivision interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v int64) *Division
	GetDivisionCode() *int64
	SetDivisionLevel(v int64) *Division
	GetDivisionLevel() *int64
	SetDivisionName(v string) *Division
	GetDivisionName() *string
	SetParentId(v int64) *Division
	GetParentId() *int64
	SetPinyin(v string) *Division
	GetPinyin() *string
}

type Division struct {
	// example:
	//
	// 310000
	DivisionCode *int64 `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// example:
	//
	// 2
	DivisionLevel *int64  `json:"divisionLevel,omitempty" xml:"divisionLevel,omitempty"`
	DivisionName  *string `json:"divisionName,omitempty" xml:"divisionName,omitempty"`
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// shang hai
	Pinyin *string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
}

func (s Division) String() string {
	return dara.Prettify(s)
}

func (s Division) GoString() string {
	return s.String()
}

func (s *Division) GetDivisionCode() *int64 {
	return s.DivisionCode
}

func (s *Division) GetDivisionLevel() *int64 {
	return s.DivisionLevel
}

func (s *Division) GetDivisionName() *string {
	return s.DivisionName
}

func (s *Division) GetParentId() *int64 {
	return s.ParentId
}

func (s *Division) GetPinyin() *string {
	return s.Pinyin
}

func (s *Division) SetDivisionCode(v int64) *Division {
	s.DivisionCode = &v
	return s
}

func (s *Division) SetDivisionLevel(v int64) *Division {
	s.DivisionLevel = &v
	return s
}

func (s *Division) SetDivisionName(v string) *Division {
	s.DivisionName = &v
	return s
}

func (s *Division) SetParentId(v int64) *Division {
	s.ParentId = &v
	return s
}

func (s *Division) SetPinyin(v string) *Division {
	s.Pinyin = &v
	return s
}

func (s *Division) Validate() error {
	return dara.Validate(s)
}
