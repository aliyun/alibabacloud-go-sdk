// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebofficePermission interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v bool) *WebofficePermission
	GetCopy() *bool
	SetExport(v bool) *WebofficePermission
	GetExport() *bool
	SetHistory(v bool) *WebofficePermission
	GetHistory() *bool
	SetPrint(v bool) *WebofficePermission
	GetPrint() *bool
	SetReadonly(v bool) *WebofficePermission
	GetReadonly() *bool
	SetRename(v bool) *WebofficePermission
	GetRename() *bool
}

type WebofficePermission struct {
	Copy     *bool `json:"Copy,omitempty" xml:"Copy,omitempty"`
	Export   *bool `json:"Export,omitempty" xml:"Export,omitempty"`
	History  *bool `json:"History,omitempty" xml:"History,omitempty"`
	Print    *bool `json:"Print,omitempty" xml:"Print,omitempty"`
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	Rename   *bool `json:"Rename,omitempty" xml:"Rename,omitempty"`
}

func (s WebofficePermission) String() string {
	return dara.Prettify(s)
}

func (s WebofficePermission) GoString() string {
	return s.String()
}

func (s *WebofficePermission) GetCopy() *bool {
	return s.Copy
}

func (s *WebofficePermission) GetExport() *bool {
	return s.Export
}

func (s *WebofficePermission) GetHistory() *bool {
	return s.History
}

func (s *WebofficePermission) GetPrint() *bool {
	return s.Print
}

func (s *WebofficePermission) GetReadonly() *bool {
	return s.Readonly
}

func (s *WebofficePermission) GetRename() *bool {
	return s.Rename
}

func (s *WebofficePermission) SetCopy(v bool) *WebofficePermission {
	s.Copy = &v
	return s
}

func (s *WebofficePermission) SetExport(v bool) *WebofficePermission {
	s.Export = &v
	return s
}

func (s *WebofficePermission) SetHistory(v bool) *WebofficePermission {
	s.History = &v
	return s
}

func (s *WebofficePermission) SetPrint(v bool) *WebofficePermission {
	s.Print = &v
	return s
}

func (s *WebofficePermission) SetReadonly(v bool) *WebofficePermission {
	s.Readonly = &v
	return s
}

func (s *WebofficePermission) SetRename(v bool) *WebofficePermission {
	s.Rename = &v
	return s
}

func (s *WebofficePermission) Validate() error {
	return dara.Validate(s)
}
