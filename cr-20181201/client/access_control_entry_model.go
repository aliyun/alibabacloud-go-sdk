// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessControlEntry interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AccessControlEntry
	GetComment() *string
	SetEntry(v string) *AccessControlEntry
	GetEntry() *string
}

type AccessControlEntry struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Entry   *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s AccessControlEntry) String() string {
	return dara.Prettify(s)
}

func (s AccessControlEntry) GoString() string {
	return s.String()
}

func (s *AccessControlEntry) GetComment() *string {
	return s.Comment
}

func (s *AccessControlEntry) GetEntry() *string {
	return s.Entry
}

func (s *AccessControlEntry) SetComment(v string) *AccessControlEntry {
	s.Comment = &v
	return s
}

func (s *AccessControlEntry) SetEntry(v string) *AccessControlEntry {
	s.Entry = &v
	return s
}

func (s *AccessControlEntry) Validate() error {
	return dara.Validate(s)
}
