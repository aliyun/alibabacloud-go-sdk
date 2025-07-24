// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeContactGroup interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*string) *MergeContactGroup
	GetContacts() []*string
	SetExtend(v map[string]interface{}) *MergeContactGroup
	GetExtend() map[string]interface{}
	SetGmtCreate(v string) *MergeContactGroup
	GetGmtCreate() *string
	SetGmtModified(v string) *MergeContactGroup
	GetGmtModified() *string
	SetIdentifier(v string) *MergeContactGroup
	GetIdentifier() *string
	SetName(v string) *MergeContactGroup
	GetName() *string
	SetSource(v string) *MergeContactGroup
	GetSource() *string
}

type MergeContactGroup struct {
	Contacts    []*string              `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	Extend      map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	GmtCreate   *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier  *string                `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	Source      *string                `json:"source,omitempty" xml:"source,omitempty"`
}

func (s MergeContactGroup) String() string {
	return dara.Prettify(s)
}

func (s MergeContactGroup) GoString() string {
	return s.String()
}

func (s *MergeContactGroup) GetContacts() []*string {
	return s.Contacts
}

func (s *MergeContactGroup) GetExtend() map[string]interface{} {
	return s.Extend
}

func (s *MergeContactGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MergeContactGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MergeContactGroup) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeContactGroup) GetName() *string {
	return s.Name
}

func (s *MergeContactGroup) GetSource() *string {
	return s.Source
}

func (s *MergeContactGroup) SetContacts(v []*string) *MergeContactGroup {
	s.Contacts = v
	return s
}

func (s *MergeContactGroup) SetExtend(v map[string]interface{}) *MergeContactGroup {
	s.Extend = v
	return s
}

func (s *MergeContactGroup) SetGmtCreate(v string) *MergeContactGroup {
	s.GmtCreate = &v
	return s
}

func (s *MergeContactGroup) SetGmtModified(v string) *MergeContactGroup {
	s.GmtModified = &v
	return s
}

func (s *MergeContactGroup) SetIdentifier(v string) *MergeContactGroup {
	s.Identifier = &v
	return s
}

func (s *MergeContactGroup) SetName(v string) *MergeContactGroup {
	s.Name = &v
	return s
}

func (s *MergeContactGroup) SetSource(v string) *MergeContactGroup {
	s.Source = &v
	return s
}

func (s *MergeContactGroup) Validate() error {
	return dara.Validate(s)
}
