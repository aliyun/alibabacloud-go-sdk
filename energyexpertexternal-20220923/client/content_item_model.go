// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentItem interface {
	dara.Model
	String() string
	GoString() string
	SetExtInfo(v []*ContentItemExtInfo) *ContentItem
	GetExtInfo() []*ContentItemExtInfo
	SetScore(v float64) *ContentItem
	GetScore() *float64
	SetText(v string) *ContentItem
	GetText() *string
	SetType(v string) *ContentItem
	GetType() *string
}

type ContentItem struct {
	// The coordinates of the text are in list format.
	ExtInfo []*ContentItemExtInfo `json:"extInfo,omitempty" xml:"extInfo,omitempty" type:"Repeated"`
	// Recall confidence
	//
	// example:
	//
	// 0.45
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// Recall text.
	//
	// example:
	//
	// content
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// Text sources: img, table, para.
	//
	// example:
	//
	// img
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ContentItem) String() string {
	return dara.Prettify(s)
}

func (s ContentItem) GoString() string {
	return s.String()
}

func (s *ContentItem) GetExtInfo() []*ContentItemExtInfo {
	return s.ExtInfo
}

func (s *ContentItem) GetScore() *float64 {
	return s.Score
}

func (s *ContentItem) GetText() *string {
	return s.Text
}

func (s *ContentItem) GetType() *string {
	return s.Type
}

func (s *ContentItem) SetExtInfo(v []*ContentItemExtInfo) *ContentItem {
	s.ExtInfo = v
	return s
}

func (s *ContentItem) SetScore(v float64) *ContentItem {
	s.Score = &v
	return s
}

func (s *ContentItem) SetText(v string) *ContentItem {
	s.Text = &v
	return s
}

func (s *ContentItem) SetType(v string) *ContentItem {
	s.Type = &v
	return s
}

func (s *ContentItem) Validate() error {
	if s.ExtInfo != nil {
		for _, item := range s.ExtInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ContentItemExtInfo struct {
	// Enumeration values for page alignment options: center, left, and right.
	//
	// example:
	//
	// center
	Alignment *string `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// The index of the layout in the text.
	//
	// example:
	//
	// 8
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
	// Hierarchy of the layout.
	//
	// example:
	//
	// 2
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// Page numbers of the layout, which may include multiple page numbers.
	PageNum []*int64 `json:"pageNum,omitempty" xml:"pageNum,omitempty" type:"Repeated"`
	// Position information of the layout, in list format.
	Pos []*ContentItemExtInfoPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// Enumeration values for the layout subtypes: picture, para, none
	//
	// example:
	//
	// picture
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// text
	//
	// example:
	//
	// content
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// Enumeration values for the layout types: table, figure, text, and none.
	//
	// example:
	//
	// table
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Unique layout ID.
	//
	// example:
	//
	// 88c712db271443dd4e3697cb9b5dab3a
	UniqueId *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
}

func (s ContentItemExtInfo) String() string {
	return dara.Prettify(s)
}

func (s ContentItemExtInfo) GoString() string {
	return s.String()
}

func (s *ContentItemExtInfo) GetAlignment() *string {
	return s.Alignment
}

func (s *ContentItemExtInfo) GetIndex() *int64 {
	return s.Index
}

func (s *ContentItemExtInfo) GetLevel() *int64 {
	return s.Level
}

func (s *ContentItemExtInfo) GetPageNum() []*int64 {
	return s.PageNum
}

func (s *ContentItemExtInfo) GetPos() []*ContentItemExtInfoPos {
	return s.Pos
}

func (s *ContentItemExtInfo) GetSubType() *string {
	return s.SubType
}

func (s *ContentItemExtInfo) GetText() *string {
	return s.Text
}

func (s *ContentItemExtInfo) GetType() *string {
	return s.Type
}

func (s *ContentItemExtInfo) GetUniqueId() *string {
	return s.UniqueId
}

func (s *ContentItemExtInfo) SetAlignment(v string) *ContentItemExtInfo {
	s.Alignment = &v
	return s
}

func (s *ContentItemExtInfo) SetIndex(v int64) *ContentItemExtInfo {
	s.Index = &v
	return s
}

func (s *ContentItemExtInfo) SetLevel(v int64) *ContentItemExtInfo {
	s.Level = &v
	return s
}

func (s *ContentItemExtInfo) SetPageNum(v []*int64) *ContentItemExtInfo {
	s.PageNum = v
	return s
}

func (s *ContentItemExtInfo) SetPos(v []*ContentItemExtInfoPos) *ContentItemExtInfo {
	s.Pos = v
	return s
}

func (s *ContentItemExtInfo) SetSubType(v string) *ContentItemExtInfo {
	s.SubType = &v
	return s
}

func (s *ContentItemExtInfo) SetText(v string) *ContentItemExtInfo {
	s.Text = &v
	return s
}

func (s *ContentItemExtInfo) SetType(v string) *ContentItemExtInfo {
	s.Type = &v
	return s
}

func (s *ContentItemExtInfo) SetUniqueId(v string) *ContentItemExtInfo {
	s.UniqueId = &v
	return s
}

func (s *ContentItemExtInfo) Validate() error {
	if s.Pos != nil {
		for _, item := range s.Pos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ContentItemExtInfoPos struct {
	// X coordinate of the layout.
	//
	// example:
	//
	// 1
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// Y coordinate of the layout.
	//
	// example:
	//
	// 2
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ContentItemExtInfoPos) String() string {
	return dara.Prettify(s)
}

func (s ContentItemExtInfoPos) GoString() string {
	return s.String()
}

func (s *ContentItemExtInfoPos) GetX() *int64 {
	return s.X
}

func (s *ContentItemExtInfoPos) GetY() *int64 {
	return s.Y
}

func (s *ContentItemExtInfoPos) SetX(v int64) *ContentItemExtInfoPos {
	s.X = &v
	return s
}

func (s *ContentItemExtInfoPos) SetY(v int64) *ContentItemExtInfoPos {
	s.Y = &v
	return s
}

func (s *ContentItemExtInfoPos) Validate() error {
	return dara.Validate(s)
}
