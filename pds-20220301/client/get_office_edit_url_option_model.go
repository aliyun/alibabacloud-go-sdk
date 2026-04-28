// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficeEditUrlOption interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v bool) *GetOfficeEditUrlOption
	GetCopy() *bool
	SetPrint(v bool) *GetOfficeEditUrlOption
	GetPrint() *bool
	SetReadonly(v bool) *GetOfficeEditUrlOption
	GetReadonly() *bool
}

type GetOfficeEditUrlOption struct {
	Copy     *bool `json:"copy,omitempty" xml:"copy,omitempty"`
	Print    *bool `json:"print,omitempty" xml:"print,omitempty"`
	Readonly *bool `json:"readonly,omitempty" xml:"readonly,omitempty"`
}

func (s GetOfficeEditUrlOption) String() string {
	return dara.Prettify(s)
}

func (s GetOfficeEditUrlOption) GoString() string {
	return s.String()
}

func (s *GetOfficeEditUrlOption) GetCopy() *bool {
	return s.Copy
}

func (s *GetOfficeEditUrlOption) GetPrint() *bool {
	return s.Print
}

func (s *GetOfficeEditUrlOption) GetReadonly() *bool {
	return s.Readonly
}

func (s *GetOfficeEditUrlOption) SetCopy(v bool) *GetOfficeEditUrlOption {
	s.Copy = &v
	return s
}

func (s *GetOfficeEditUrlOption) SetPrint(v bool) *GetOfficeEditUrlOption {
	s.Print = &v
	return s
}

func (s *GetOfficeEditUrlOption) SetReadonly(v bool) *GetOfficeEditUrlOption {
	s.Readonly = &v
	return s
}

func (s *GetOfficeEditUrlOption) Validate() error {
	return dara.Validate(s)
}
