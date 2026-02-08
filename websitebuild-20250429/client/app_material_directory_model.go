// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppMaterialDirectory interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AppMaterialDirectory
	GetBizId() *string
	SetChildren(v []*AppMaterialDirectory) *AppMaterialDirectory
	GetChildren() []*AppMaterialDirectory
	SetDirectoryId(v string) *AppMaterialDirectory
	GetDirectoryId() *string
	SetName(v string) *AppMaterialDirectory
	GetName() *string
	SetSortNum(v string) *AppMaterialDirectory
	GetSortNum() *string
	SetType(v string) *AppMaterialDirectory
	GetType() *string
}

type AppMaterialDirectory struct {
	BizId       *string                 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Children    []*AppMaterialDirectory `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	DirectoryId *string                 `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Name        *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	SortNum     *string                 `json:"SortNum,omitempty" xml:"SortNum,omitempty"`
	Type        *string                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppMaterialDirectory) String() string {
	return dara.Prettify(s)
}

func (s AppMaterialDirectory) GoString() string {
	return s.String()
}

func (s *AppMaterialDirectory) GetBizId() *string {
	return s.BizId
}

func (s *AppMaterialDirectory) GetChildren() []*AppMaterialDirectory {
	return s.Children
}

func (s *AppMaterialDirectory) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AppMaterialDirectory) GetName() *string {
	return s.Name
}

func (s *AppMaterialDirectory) GetSortNum() *string {
	return s.SortNum
}

func (s *AppMaterialDirectory) GetType() *string {
	return s.Type
}

func (s *AppMaterialDirectory) SetBizId(v string) *AppMaterialDirectory {
	s.BizId = &v
	return s
}

func (s *AppMaterialDirectory) SetChildren(v []*AppMaterialDirectory) *AppMaterialDirectory {
	s.Children = v
	return s
}

func (s *AppMaterialDirectory) SetDirectoryId(v string) *AppMaterialDirectory {
	s.DirectoryId = &v
	return s
}

func (s *AppMaterialDirectory) SetName(v string) *AppMaterialDirectory {
	s.Name = &v
	return s
}

func (s *AppMaterialDirectory) SetSortNum(v string) *AppMaterialDirectory {
	s.SortNum = &v
	return s
}

func (s *AppMaterialDirectory) SetType(v string) *AppMaterialDirectory {
	s.Type = &v
	return s
}

func (s *AppMaterialDirectory) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
